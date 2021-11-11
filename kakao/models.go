package kakao

// 카카오 서버에서 넘어온 json 형식
type UserMessage struct {
	Request UserRequest `json:"userRequest"`
}

// 수신된 json 형식 중 사용자 발화 부분 분리
type UserRequest struct {
	Message string `json:"utterance"`
}


// 송신 json 형식
type ServerResponse struct {
	Version string `json:"version"`
	Template SkillTemplate `json:"template"`
}

// 송신 Response Template
type SkillTemplate struct {
	Outputs []Components `json:"outputs"`
	QuickReplies []QuickReply `json:"quickReplies"`
}

// Output Template
type Components interface {}

type SimpleTextResponse struct {
	SimpleText TextContent `json:"simpleText"`
}

type TextContent struct {
	Text string `json:"text"`
}

type BasicCardResponse struct {
	Card TextCard `json:"basicCard"`
}

type TextCard struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Buttons []CardButton `json:"buttons"`
}

type CarouselResponse struct {
	Carousel Carousel `json:"carousel"`
}

type Carousel struct {
	Type  string     `json:"type"`
	Items []TextCard `json:"items"`
}

// Quick Reply Template
type QuickReply struct {
	Action string `json:"action"`
	Label string `json:"label"`
	MessageText string `json:"messageText"`
	BlockID string `json:"blockId"`
}

// Card Button
type CardButton struct {
	Action string `json:"action"`
	Label string `json:"label"`
	Link string `json:"webLinkUrl"`
}