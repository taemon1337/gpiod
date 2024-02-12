package orangepi

import (
	"errors"
	"strconv"
	"strings"
)

// ErrInvalid indicates the pin name does not match a known pin.
var ErrInvalid = errors.New("invalid pin number")

func rangeCheck(p int) (int, error) {
	if p < 2 || p >= 27 {
		return 0, ErrInvalid
	}
	return p, nil
}

// Pin maps a pin string name to a pin number.
//
// Pin names are case insensitive and may be of the form GPIOX, or X.
func Pin(s string) (int, error) {
	s = strings.ToLower(s)
	switch {
	case strings.HasPrefix(s, "gpio"):
		v, err := strconv.ParseInt(s[4:], 10, 8)
		if err != nil {
			return 0, err
		}
		return rangeCheck(int(v))
	default:
		v, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return 0, err
		}
		return rangeCheck(int(v))
	}
}

// MustPin converts the string to the corresponding pin number or panics if that
// is not possible.
func MustPin(s string) int {
	v, err := Pin(s)
	if err != nil {
		panic(err)
	}
	return v
}
