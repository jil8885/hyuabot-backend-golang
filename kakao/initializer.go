package kakao

func setResponse(template SkillTemplate) ServerResponse {
	return ServerResponse{"2.0", template}
}

func setTemplate(outputs []Components, replies []QuickReply) SkillTemplate {
	return SkillTemplate{outputs, replies}
}

func setSimpleText(message string) SimpleTextResponse {
	return SimpleTextResponse{SimpleText: TextContent{message}}
}

func setBasicCard(title string, message string, buttons []CardButton) TextCard {
	return TextCard{Title: title, Description: message, Buttons: buttons}
}