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
	if strings.Contains(message, "의 셔틀버스 도착 정보"){
		temp = strings.Split(message, "의 셔틀버스 도착 정보입니다")[0]

	} else {
		temp = strings.TrimSpace(strings.Split(message, " ")[1])
	}
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
		for index, item := range busForStation{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "기숙사→예술인\n"
		for index, item := range busForTerminal{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "기숙사 출발 버스는 셔틀콕을 경유합니다.\n"
	case "Shuttlecock_O":
		message += "셔틀콕→한대앞\n"
		for index, item := range busForStation{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "셔틀콕→예술인\n"
		for index, item := range busForTerminal{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "한대앞 방면은 순환, 직행 중 앞에 오는 것이 빠릅니다.\n"
	case "Subway":
		message += "한대앞→ERICA\n"
		for index, item := range busForStation{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "캠퍼스 방면은 순환, 직행 중 앞에 오는 것이 빠릅니다.\n"
	case "Terminal":
		message += "예술인→ERICA\n"
		for index, item := range busForTerminal{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "예술인 출발 버스는 셔틀콕, 기숙사 방면으로 운행합니다.\n"
	case "Shuttlecock_I":
		message += "셔틀콕 건너편→기숙사\n"
		for index, item := range busForTerminal{
			message += strings.Replace(item.Time, ":", "시", 1) + "분 출발 예정\n"
			if index > 1{
				break
			}
		}
		message += "일부 차량은 기숙사로 가지 않을 수 있습니다.\n"
	}
	message += "제공되는 출발 시간표는 시간표 기반으로, 미리 정류장에서 기다리는 것을 추천드립니다."

	// 바로가기 버튼
	var replies []QuickReply
	replies = append(replies, QuickReply{"blocks", "앱 설치", "앱 설치 안내입니다.", "6077ca2de2039a2ba38c755f"})
	replies = append(replies, QuickReply{"blocks", "🔍 정류장", busStop + " 정류장 정보입니다.", "5ebf702e7a9c4b000105fb25"})
	replies = append(replies, QuickReply{"blocks", "🚫 오류제보", "셔틀 오류 제보하기", "5cc3fced384c5508fceec5bb"})

	response := setResponse(setTemplate([]Components{setSimpleText(message)}, replies))
	return c.JSON(response)
}

// 카카오 i 셔틀 정류장 정보 제공
func ShuttleStop(c *fiber.Ctx) error {
	return c.SendString("카카오 i 셔틀 정류장 정보")
}

// 카카오 i 전철 도착 정보 제공
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