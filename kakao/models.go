package kakao

// 카카오 서버에서 넘어온 json 형식
type UserMessage struct {
	Request UserRequest `json:"userRequest"`
}

// 수신된 json 형식 중 사용자 발화 부분 분리
type UserRequest struct {
	Message string `json:"utterance"`
}
