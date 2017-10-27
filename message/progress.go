package message

func NewProgress() *Progress {
	pr := Progress{}
	return &pr
}

type Progress struct {
	Progress float32 `json:"progress"`
}
