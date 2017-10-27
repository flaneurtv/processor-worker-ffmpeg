package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	durationRegexp = regexp.MustCompile(`Duration: (\d\d:\d\d:\d\d.\d\d), start`)
	outTimeRegexp  = regexp.MustCompile(`out_time=(\d\d:\d\d:\d\d.\d\d)`)
)

var (
	ErrParseDuration    = errors.New("error parsing duration")
	ErrDurationSubmatch = errors.New("error submatching duration")
	ErrOutTimeSubmatch  = errors.New("error submatching out time")
)

// Parsing string "out_time=00:02:01.04" to get 121 (sec)
func ParseOutTime(chunk string) (int, error) {
	c := outTimeRegexp.FindStringSubmatch(chunk)
	if len(c) < 2 {
		return 0, ErrOutTimeSubmatch
	}
	raw := c[1]

	return parseSeconds(raw)
}

// Parsing string "Duration: 00:02:01.04, start" to get 121 (sec)
func ParseDuration(line string) (duration int, err error) {
	c := durationRegexp.FindStringSubmatch(line)
	if len(c) < 2 {
		err = ErrDurationSubmatch
		return
	}
	raw := c[1]

	return parseSeconds(raw)
}

// Parsing string "00:02:01.04" to return 121 (sec)
func parseSeconds(raw string) (seconds int, err error) {
	arr := strings.Split(raw, ":")
	sec := strings.Split(arr[2], ".")[0]
	str := fmt.Sprintf("%sh%sm%ss", arr[0], arr[1], sec)

	var d time.Duration
	if d, err = time.ParseDuration(str); err != nil {
		return
	}
	seconds = int(d.Seconds())
	return
}
