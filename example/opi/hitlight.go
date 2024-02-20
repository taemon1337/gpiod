package main

import (
	"fmt"
	"os"
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/taemon1337/gpiod"
	"github.com/taemon1337/gpiod/device/orangepi"
)

var (
	OFF int = 0
	ON int = 1
  BLINK_COUNT int = 5
)

func printEvent(evt gpiod.LineEvent, led *gpiod.Line) {
	t := time.Now()
	edge := "rising"
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
	} else {
		fmt.Println("turning ON")
		led.SetValue(ON)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("turning OFF")
		led.SetValue(OFF)
		time.Sleep(100 * time.Millisecond)
	}

	if evt.Seqno != 0 {
		// only uAPI v2 populates the sequence numbers
		fmt.Printf("event: #%d(%d)%3d %-7s %s (%s)\n",
			evt.Seqno,
			evt.LineSeqno,
			evt.Offset,
			edge,
			t.Format(time.RFC3339Nano),
			evt.Timestamp)
	} else {
		fmt.Printf("event:%3d %-7s %s (%s)\n",
			evt.Offset,
			edge,
			t.Format(time.RFC3339Nano),
			evt.Timestamp)
	}
}

func main() {
	echan := make(chan gpiod.LineEvent, 6)
	ledpin := orangepi.GPIO13
	hitpin := orangepi.GPIO2

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	eh := func(evt gpiod.LineEvent) {
		select {
		case echan <- evt:
		default:
			fmt.Printf("event chan overflow - discarding event")
		}
	}

	led, err := gpiod.RequestLine("gpiochip0", ledpin, gpiod.AsOutput(OFF))
	if err != nil {
		fmt.Printf("RequestLine returned error: %s\n", err)
		if err == syscall.Errno(22) {
			fmt.Println("Note that the WithPullUp optiON requires kernel V5.5 or later - check your kernel versiON.")
		}
		os.Exit(1)
	}

	// period := 10 * time.Millisecond
	hit, err := gpiod.RequestLine("gpiochip0", hitpin, gpiod.WithPullUp, gpiod.WithRisingEdge, gpiod.WithEventHandler(eh))
	if err != nil {
		fmt.Printf("RequestLine returned error: %s\n", err)
		if err == syscall.Errno(22) {
			fmt.Println("Note that the WithPullUp optiON requires kernel V5.5 or later - check your kernel versiON.")
		}
		os.Exit(1)
	}

	// start by blinking n times
  for i := 1; i <= BLINK_COUNT; i++ {
		fmt.Println("turning ON")
		led.SetValue(ON)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("turning OFF")
		led.SetValue(OFF)
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)

	// revert line to input ON the way out
	defer func() {
		led.Reconfigure(gpiod.AsInput)
		hit.Reconfigure(gpiod.AsInput)
		led.Close()
		hit.Close()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	fmt.Println("watching pin %d", hitpin)
	done := false
	for !done {
		select {
		case evt := <-echan:
			printEvent(evt, led)
		case <-ctx.Done():
			fmt.Println("exiting...")
			done = true
		case <-quit:
			fmt.Printf("stopping...")
			done = true
		}
	}
	cancel()
}
