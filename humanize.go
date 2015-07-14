package strings

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode"
)

var sizes = [...]string{"B", "KiB", "MiB", "GiB", "TiB"}

func logn(n, b float64) float64 { return math.Log(n) / math.Log(b) }

// HumanizeSize takes a uint64 and produces a string of it to the highest IEC size
func HumanizeSize(size uint64) string {
	const base float64 = 1024
	if size < 10 {
		return fmt.Sprintf("%d B", size)
	}

	e := math.Floor(logn(float64(size), base))
	val := float64(size) / math.Pow(base, math.Floor(e))

	f := "%0.f"
	if val < 10 {
		f = "%.1f"
	}
	return fmt.Sprintf(f+" %s", val, sizes[int(e)])
}

// HumanizeNumber takes a uint64 and returns a string with a comma in the thousands (and such) place
func HumanizeNumber(input uint64) string {
	return humanize(input, 1000, ",", "%d", "%03d")
}

// HumanizeTime takes a uint64 and returns a string with it formatted as HH:NN:SS
func HumanizeTime(input uint64) string {
	res := humanize(input, 60, ":", "%02d", "%02d")
	if len(res) == 2 {
		res = "00:" + res
	}
	return res
}

func humanize(input uint64, step int, sep, lead, mid string) string {
	temp, scale := 0, 1
	for input >= step {
		temp += scale * (input % step)
		input /= step
		scale *= step
	}

	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, lead, input)

	for scale != 1 {
		scale /= step
		input = temp / scale
		temp %= scale
		fmt.Fprintf(buf, fmt.Sprintf("%s%s", sep, mid), input)
	}
	return buf.String()
}

// Iso8601Humanize takes an ISO8601 timestamp and returns it formated as HH:NN:SS, with minutes set as a minimum truncation
func Iso8601Humanize(s string) string {
	if len(s) <= 2 || s[:2] != "PT" {
		return ""
	}
	s = s[2:]

	var p int
	var prev, hours, minutes, seconds []uint8
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			prev = append(prev, s[i])
			continue
		}
		switch s[i] {
		case 'H':
			hours = append(hours, prev[p:]...)
			p = len(prev)
		case 'M':
			minutes = append(minutes, prev[p:]...)
			p = len(prev)
		case 'S':
			seconds = append(seconds, prev[p:]...)
			p = len(prev)
		}
	}

	fix := func(s []uint8) string {
		if len(s) == 1 {
			return "0" + string(s[0])
		}
		if len(s) == 0 {
			return "00"
		}
		return string(s[0]) + string(s[1])
	}

	var fixed []string
	if h := fix(hours); h != "00" {
		fixed = []string{h, fix(minutes), fix(seconds)}
	} else {
		fixed = []string{fix(minutes), fix(seconds)}
	}
	return strings.Join(fixed, ":")
}
