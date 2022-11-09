package types

type Event struct {
	Font        int    `json:"font"`
	Message     string `json:"message"`
	MessageId   int    `json:"message_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	RawMessage  string `json:"raw_message"`
	SelfId      int    `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Role     string `json:"role"`
		Title    string `json:"title"`
		Level    string `json:"level"`
		Sex      string `json:"sex"`
		UserId   int    `json:"user_id"`
	} `json:"sender"`
	SubType    string `json:"sub_type"`
	TargetId   int    `json:"target_id"`
	Time       int    `json:"time"`
	UserId     int    `json:"user_id"`
	Anonymous  string `json:"anonymous"`
	GroupId    int    `json:"group_id"`
	MessageSeq int    `json:"message_seq"`
	Comment    string `json:"comment"`
	Flag       string `json:"flag"`
}
