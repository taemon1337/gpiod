package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/taemon1337/gpiod"
	"github.com/taemon1337/gpiod/device/orangepi"
)

func printEvent(evt gpiod.LineEvent) {
	t := time.Now()
	edge := "rising"
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
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
	values := map[int]string{0: "inactive", 1: "active"}
	offset := orangepi.GPIO13
	v := 0
	l, err := gpiod.RequestLine("gpiochip0", offset, gpiod.AsOutput(v))
	if err != nil {
		fmt.Printf("RequestLine returned error: %s\n", err)
		if err == syscall.Errno(22) {
			fmt.Println("Note that the WithPullUp option requires kernel V5.5 or later - check your kernel version.")
		}
		os.Exit(1)
	}

	// revert line to input on the way out
	defer func() {
		l.Reconfigure(gpiod.AsInput)
		l.Close()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	for {
		select {
		case <-time.After(2 * time.Second):
			v ^= 1
			l.SetValue(v)
			fmt.Printf("Set pin %d %s\n", offset, values[v])
		case <-quit:
			fmt.Printf("stopping...")
			return
		}
	}
}

