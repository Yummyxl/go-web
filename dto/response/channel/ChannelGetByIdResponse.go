package channel

type ChannelGetByIdResponse struct {
	Id          int64  `json:"id"`
	DisplayName string `json:"displayName"`
	NoteName    string `json:"noteName"`
	IsHandpick  int    `json:"isHandpick"`
}
