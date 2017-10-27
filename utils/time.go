package utils

import "time"

func Now() string {
	return Format(time.Now())
}

func Format(t time.Time) string {
	return t.Format("02.01.06 15:04:05")
}
