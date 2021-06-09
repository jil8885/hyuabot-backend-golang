package kakao

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jil8885/hyuabot-backend-golang/shuttle"
	"strings"
	"time"
)

// 카카오 i 용 url handler
func Middleware(c *fiber.Ctx) error {
	c.Accepts("application/json") // json 형식으로만 요청 가능
	return c.Next()
}

// 카카오 i 셔틀 도착 정보 제공
func Shuttle(c *fiber.Ctx) error {
	message := parseAnswer(c)
	// 사용자 메세지에서 셔틀버스 정보 추출
	busStop := ""
	temp := ""
	otherStops := [5]string{"🏘️ 기숙사", "🏫 셔틀콕", "🚆 한대앞역", "🚍 예술인A", "🏫 셔틀콕 건너편"}

	temp = strings.TrimSpace(message[strings.Index(message, " "):])

	switch temp {
	case "기숙사":
		busStop = "Residence"
	case "셔틀콕":
		busStop = "Shuttlecock_O"
	case "한대앞역":
		busStop = "Subway"
	case "예술인A":
		busStop = "Terminal"
	case "셔틀콕 건너편":
		busStop = "Shuttlecock_I"
	}

	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	busForStation, busForTerminal := shuttle.GetShuttle(busStop, now)
	message = ""
	switch busStop {
	case "Residence":
		message += "기숙사→한대앞\n"
		if len(busForStation) > 0{
			for _, item := range busForStation{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}

		message += "기숙사→예술인\n"
		if len(busForTerminal) > 0{
			for _, item := range busForTerminal{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}
		message += "기숙사 출발 버스는 셔틀콕을 경유합니다.\n"
	case "Shuttlecock_O":
		message += "셔틀콕→한대앞\n"
		if len(busForStation) > 0{
			for _, item := range busForStation{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}

		message += "셔틀콕→예술인\n"
		if len(busForTerminal) > 0{
			for _, item := range busForTerminal{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}
		message += "한대앞 방면은 순환, 직행 중 앞에 오는 것이 빠릅니다.\n"
	case "Subway":
		message += "한대앞→셔틀콕,기숙사\n"
		if len(busForStation) > 0{
			for _, item := range busForStation{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}

		message += "한대앞→예술인\n"
		if len(busForTerminal) > 0{
			for _, item := range busForTerminal{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}
		
		message += "캠퍼스 방면은 순환, 직행 중 앞에 오는 것이 빠릅니다.\n"
	case "Terminal":
		message += "예술인→셔틀콕,기숙사\n"
		if len(busForTerminal) > 0{
			for _, item := range busForTerminal{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}
	case "Shuttlecock_I":
		message += "셔틀콕 건너편→기숙사\n"
		if len(busForTerminal) > 0{
			for _, item := range busForTerminal{
				message += strings.Replace(item.Time, ":", "시 ", 1) + "분 출발(" + strings.Replace(strings.Replace(item.Heading, "C", "순환", 1), "DH", "직행", 1) + ")\n"
			}
			message += "\n"
		} else {
			message += "운행 종료\n\n"
		}
		message += "일부 차량 기숙사 종착\n"
	}

	// 바로가기 버튼
	var replies []QuickReply
	replies = append(replies, QuickReply{"block", "앱 설치", "앱 설치 안내입니다.", "6077ca2de2039a2ba38c755f"})
	replies = append(replies, QuickReply{"block", "🔍 정류장", temp + " 정류장 정보입니다.", "5ebf702e7a9c4b000105fb25"})
	replies = append(replies, QuickReply{"block", "🚫 오류제보", "셔틀 오류 제보하기", "5cc3fced384c5508fceec5bb"})

	for _, stop := range otherStops{
		replies = append(replies, QuickReply{"block", stop, stop, "5cc3dc8ee82127558b7e6eba"})
	}

	response := setResponse(setTemplate([]Components{setSimpleText(strings.TrimSpace(message))}, replies))
	return c.JSON(response)
}

// ShuttleStop 카카오 i 셔틀 정류장 정보 제공
func ShuttleStop(c *fiber.Ctx) error {
	message := parseAnswer(c)
	temp := strings.TrimSpace(strings.Split(message, " 정류장 정보입니다.")[0])
	var busStop string
	var roadViewLink string

	switch temp {
	case "기숙사":
		busStop = "Residence"
		roadViewLink = "http://kko.to/R-l1jU3DT"
	case "셔틀콕":
		busStop = "Shuttlecock_O"
		roadViewLink = "http://kko.to/TyWyjU3Yp"
	case "한대앞역":
		busStop = "Subway"
		roadViewLink = "http://kko.to/c93C0UFYj"
	case "예술인A":
		busStop = "Terminal"
		roadViewLink = "http://kko.to/7mzoYUFY0"
	case "셔틀콕 건너편":
		busStop = "Shuttlecock_I"
		roadViewLink = "http://kko.to/TyWyjU3Yp"
	}

	message = ""

	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	busForStationFirst, busForStationLast, busForTerminalFirst, busForTerminalLast := shuttle.GetFirstLastShuttle(busStop, now)
	switch busStop {
	case "Residence", "Shuttlecock_O":
		message += "한대앞 : "
		message += busForStationFirst.Time + "(첫차)/"
		message += busForStationLast.Time + "(막차)\n"
		message += "예술인 : "
		message += busForTerminalFirst.Time + "(첫차)/"
		message += busForTerminalLast.Time + "(막차)\n"
	case "Subway":
		message += "셔틀콕,기숙사 : "
		message += busForStationFirst.Time + "(첫차)/"
		message += busForStationLast.Time + "(막차)\n"
		message += "예술인 : "
		message += busForTerminalFirst.Time + "(첫차)/"
		message += busForTerminalLast.Time + "(막차)\n"
	case "Terminal":
		message += "셔틀콕,기숙사 : "
		message += busForTerminalFirst.Time + "(첫차)/"
		message += busForTerminalLast.Time + "(막차)\n"
	case "Shuttlecock_I":
		message += "기숙사 : "
		message += busForTerminalFirst.Time + "(첫차)/"
		message += busForTerminalLast.Time + "(막차)\n"
	}

	var buttons []CardButton
	buttons = append(buttons, CardButton{Action: "webLink", Label: "👀 로드뷰로 보기", Link: roadViewLink})

	replies := make([]QuickReply, 0)

	response := setResponse(setTemplate([]Components{BasicCardResponse{Card: setBasicCard(temp, message, buttons)}}, replies))
	return c.JSON(response)
}

// Subway 카카오 i 전철 도착 정보 제공
func Subway(c *fiber.Ctx) error {
	return c.SendString("카카오 i 전철 도착 정보")
}

// 카카오 i 버스 도착 정보 제공
func Bus(c *fiber.Ctx) error {
	return c.SendString(parseAnswer(c))
}

// 카카오 i 학식 정보 제공
func Food(c *fiber.Ctx) error {
	return c.SendString("카카오 i 학식 정보")
}

// 카카오 i 열람실 정보 제공
func Library(c *fiber.Ctx) error {
	return c.SendString("카카오 i 열람실 정보")
}

// 카카오톡을 통해 넘어온 데이터 중 사용자의 발화 Parse
func parseAnswer(c *fiber.Ctx) string {
	model := new(UserMessage)
	if err := c.BodyParser(model); err != nil{
		return err.Error()
	}
	return model.Request.Message
}