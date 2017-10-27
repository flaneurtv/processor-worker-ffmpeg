package message

func NewTickReference() *TickReference {
	return &TickReference{}
}

type TickReference struct {
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}
