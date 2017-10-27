package worker

var seconds, duration int

func Progress() (ratio float32) {
	if ratio = float32(seconds) / float32(duration); ratio > 1 {
		return 1
	}
	return
}

func SetSeconds(s int) {
	seconds = s
}

func SetDuration(d int) {
	duration = d
}
