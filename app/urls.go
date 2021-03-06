package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jil8885/hyuabot-backend-golang/bus"
	"github.com/jil8885/hyuabot-backend-golang/food"
	"github.com/jil8885/hyuabot-backend-golang/library"
	"github.com/jil8885/hyuabot-backend-golang/shuttle"
	"github.com/jil8885/hyuabot-backend-golang/subway"
	"strings"
	"time"
)

// App 용 url handler
func Middleware(c *fiber.Ctx) error {
	c.Accepts("application/json") // json 형식으로만 요청 가능
	return c.Next()
}

func GetShuttleDeparture(c *fiber.Ctx) error {
	request := new(ShuttleStopRequest)

	if err := c.QueryParser(request); err != nil || request.BusStop == "" {
		// 현재 시간 로딩 (KST)
		loc, _ := time.LoadLocation("Asia/Seoul")
		now := time.Now().In(loc)
		busStopList := [5]string{"Residence", "Shuttlecock_O", "Subway", "Terminal", "Shuttlecock_I"}
		response := map[string]ShuttleDepartureByStop{}
		for _, item := range busStopList {
			busForStation, busForTerminal := shuttle.GetShuttle(item, now, loc, 2)
			response[item] = ShuttleDepartureByStop{BusForStation: busForStation, BusForTerminal: busForTerminal}
		}
		return c.JSON(response)
	} else {
		return GetShuttleDepartureByStop(c, request.BusStop)
	}
}

func GetShuttleDepartureByStopBackport(c *fiber.Ctx) error {
	return GetShuttleDepartureByStop(c, parseShuttleStop(c))
}

func GetShuttleDepartureByStop(c *fiber.Ctx, busStop string) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	busForStation, busForTerminal := shuttle.GetShuttle(busStop, now, loc, 2)
	response := ShuttleDepartureByStop{BusForStation: busForStation, BusForTerminal: busForTerminal}
	return c.JSON(response)
}

func GetShuttleStopInfoByStopWithParams(c *fiber.Ctx) error {
	request := new(ShuttleStopRequest)

	err := c.QueryParser(request)
	if err != nil {
		return GetShuttleDepartureByStop(c, "")
	}
	return GetShuttleDepartureByStop(c, request.BusStop)
}

func GetShuttleStopInfoByStopBackport(c *fiber.Ctx) error {
	return GetShuttleStopInfoByStop(c, parseShuttleStop(c))
}

func GetShuttleStopInfoByStop(c *fiber.Ctx, busStop string) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	firstBusForStation, lastBusForStation, firstBusForTerminal, lastBusForTerminal := shuttle.GetFirstLastShuttle(busStop, now, loc)
	weekBusForStation, weekBusForTerminal := shuttle.GetShuttleTimetable(busStop, now, loc, "week", true)
	weekendBusForStation, weekendBusForTerminal := shuttle.GetShuttleTimetable(busStop, now, loc, "weekend", true)
	roadViewMap := map[string]string{"Shuttlecock_I": "http://kko.to/TyWyjU3Yp", "Subway": "http://kko.to/c93C0UFYj", "Residence": "http://kko.to/R-l1jU3DT", "Terminal": "http://kko.to/7mzoYUFY0", "Shuttlecock_O": "http://kko.to/v-3DYI3YM"}

	return c.JSON(ShuttleStop{
		RoadViewLink:        roadViewMap[busStop],
		FirstBusForStation:  firstBusForStation.Time,
		LastBusForStation:   lastBusForStation.Time,
		FirstBusForTerminal: firstBusForTerminal.Time,
		LastBusForTerminal:  lastBusForTerminal.Time,
		Weekdays:            ShuttleDepartureByStop{BusForStation: weekBusForStation, BusForTerminal: weekBusForTerminal},
		Weekends:            ShuttleDepartureByStop{BusForStation: weekendBusForStation, BusForTerminal: weekendBusForTerminal},
	})
}

func GetSubwayDepartureWithParams(c *fiber.Ctx) error {
	request := new(CampusRequest)

	err := c.QueryParser(request)
	if err != nil {
		return GetSubwayDeparture(c, false)
	}
	return GetSubwayDeparture(c, strings.ToLower(request.Campus) == "seoul")
}

func GetSubwayDepartureBackport(c *fiber.Ctx) error {
	return GetSubwayDeparture(c, strings.ToLower(parseCampus(c)) == "seoul")
}

func GetSubwayDeparture(c *fiber.Ctx, campus bool) error {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Seoul")
	_, isWeekends := shuttle.GetDate(now, loc)
	isHoliday := shuttle.IsHoliday(now)

	if campus {
		return c.JSON(SubwayDepartureSeoul{
			Line2: subway.GetRealtimeSubway(1, 1002),
		})
	} else {
		timetableWeekdaysLine4, timetableWeekendsLine4 := subway.GetTimetableSubwayAll(1004)
		timetableWeekdaysLineSuin, timetableWeekendsLineSuin := subway.GetTimetableSubwayAll(1071)

		if isWeekends == "weekend" || isHoliday {
			return c.JSON(SubwayDepartureERICA{
				Line4: SubwayDepartureByLine{
					RealtimeList:  subway.GetRealtimeSubway(0, 1004),
					TimetableList: timetableWeekendsLine4,
				},
				LineSuin: SubwayDepartureByLine{
					RealtimeList:  subway.GetRealtimeSubway(0, 1071),
					TimetableList: timetableWeekendsLineSuin,
				},
			})
		} else {
			return c.JSON(SubwayDepartureERICA{
				Line4: SubwayDepartureByLine{
					RealtimeList:  subway.GetRealtimeSubway(0, 1004),
					TimetableList: timetableWeekdaysLine4,
				},
				LineSuin: SubwayDepartureByLine{
					RealtimeList:  subway.GetRealtimeSubway(0, 1071),
					TimetableList: timetableWeekdaysLineSuin,
				},
			})
		}

	}
}

func GetBusDeparture(c *fiber.Ctx) error {
	responseRealtimeByStop := bus.GetRealtimeStopDeparture("216000379")
	response := Bus{
		LineGreenToStation: BusByRoute{
			Realtime:  []bus.DepartureItem{},
			Timetable: bus.BusTimeTableLine{Weekdays: []bus.BusTimeTableItem{}, Sat: []bus.BusTimeTableItem{}, Sun: []bus.BusTimeTableItem{}},
		},
		LineGreenToCampus: BusByRoute{
			Realtime:  bus.GetRealtimeBusDeparture("216000138", "216000068"),
			Timetable: bus.BusTimeTableLine{Weekdays: []bus.BusTimeTableItem{}, Sat: []bus.BusTimeTableItem{}, Sun: []bus.BusTimeTableItem{}},
		},
		LineBlue: BusByRoute{
			Realtime:  bus.GetRealtimeBusDeparture("216000719", "216000070"),
			Timetable: bus.BusTimeTableLine{Weekdays: []bus.BusTimeTableItem{}, Sat: []bus.BusTimeTableItem{}, Sun: []bus.BusTimeTableItem{}},
		},
		LineRed: BusByRoute{
			Realtime:  []bus.DepartureItem{},
			Timetable: bus.BusTimeTableLine{Weekdays: []bus.BusTimeTableItem{}, Sat: []bus.BusTimeTableItem{}, Sun: []bus.BusTimeTableItem{}},
		},
	}
	for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList {
		if item.PredictTime1 > 0 {
			if item.RouteID == 216000061 {
				response.LineRed.Realtime = append(response.LineRed.Realtime, bus.DepartureItem{
					Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
				})
			} else if item.RouteID == 216000068 {
				response.LineGreenToStation.Realtime = append(response.LineGreenToStation.Realtime, bus.DepartureItem{
					Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
				})
			}
			if item.PredictTime2 > 0 {
				if item.RouteID == 216000061 {
					response.LineRed.Realtime = append(response.LineRed.Realtime, bus.DepartureItem{
						Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
					})
				} else if item.RouteID == 216000068 {
					response.LineGreenToStation.Realtime = append(response.LineGreenToStation.Realtime, bus.DepartureItem{
						Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
					})
				}
			}
		}
	}
	timetable := bus.GetTimetable()
	response.LineGreenToStation.Timetable = timetable.Line10_1
	response.LineGreenToCampus.Timetable = timetable.Line10_1
	response.LineBlue.Timetable = timetable.Line707_1
	response.LineRed.Timetable = timetable.Line3102
	return c.JSON(response)
}

func GetBusDepartureByLineWithParams(c *fiber.Ctx) error {
	request := new(BusRouteRequest)

	err := c.QueryParser(request)
	if err != nil || request.Route == "" {
		return GetBusDeparture(c)
	}
	return GetBusDepartureByLine(c, request.Route)
}

func GetBusDepartureByLineBackport(c *fiber.Ctx) error {
	return GetBusDepartureByLine(c, parseBusRouteID(c))
}

func GetBusDepartureByLine(c *fiber.Ctx, routeID string) error {
	responseTimetable := bus.GetTimetable()
	if routeID == "10-1" {
		responseRealtimeByStop := bus.GetRealtimeStopDeparture("216000379")
		var responseRealtime []bus.DepartureItem
		for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList {
			if item.RouteID == 216000068 {
				if item.PredictTime1 > 0 {
					responseRealtime = append(responseRealtime, bus.DepartureItem{
						Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
					})
					if item.PredictTime2 > 0 {
						responseRealtime = append(responseRealtime, bus.DepartureItem{
							Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
						})
					}
				}
				break
			}
		}
		return c.JSON(BusByRoute{
			Realtime:  responseRealtime,
			Timetable: responseTimetable.Line10_1,
		})
	} else if routeID == "3102" {
		responseRealtimeByStop := bus.GetRealtimeStopDeparture("216000379")
		var responseRealtime []bus.DepartureItem
		for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList {
			if item.RouteID == 216000061 {
				if item.PredictTime1 > 0 {
					responseRealtime = append(responseRealtime, bus.DepartureItem{
						Location: item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
					})
					if item.PredictTime2 > 0 {
						responseRealtime = append(responseRealtime, bus.DepartureItem{
							Location: item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
						})
					}
				}
				break
			}
		}
		return c.JSON(BusByRoute{
			Realtime:  responseRealtime,
			Timetable: responseTimetable.Line3102,
		})
	} else if routeID == "707-1" {
		return c.JSON(BusByRoute{
			Realtime:  bus.GetRealtimeBusDeparture("216000719", "216000070"),
			Timetable: responseTimetable.Line707_1,
		})
	}
	return c.JSON(BusByRoute{
		Realtime:  []bus.DepartureItem{},
		Timetable: bus.BusTimeTableLine{},
	})
}

func GetBusTimetableByRouteWithParams(c *fiber.Ctx) error {
	request := new(BusRouteRequest)

	err := c.QueryParser(request)
	if err != nil {
		return GetBusTimetableByRoute(c, "")
	}
	return GetBusTimetableByRoute(c, request.Route)
}

func GetBusTimetableByRouteBackport(c *fiber.Ctx) error {
	return GetBusTimetableByRoute(c, parseBusRouteID(c))
}

func GetBusTimetableByRoute(c *fiber.Ctx, routeID string) error {
	responseTimetable := bus.GetTimetable()
	if routeID == "10-1" {
		return c.JSON(responseTimetable.Line10_1)
	} else if routeID == "3102" {
		return c.JSON(responseTimetable.Line3102)
	} else if routeID == "707-1" {
		return c.JSON(responseTimetable.Line707_1)
	}
	return c.JSON(BusByRoute{
		Realtime:  []bus.DepartureItem{},
		Timetable: bus.BusTimeTableLine{},
	})
}

func GetReadingRoomSeatByCampusWithParams(c *fiber.Ctx) error {
	request := new(CampusRequest)

	err := c.QueryParser(request)
	if err != nil {
		return GetReadingRoomSeatByCampus(c, false)
	}
	return GetReadingRoomSeatByCampus(c, strings.ToLower(request.Campus) == "seoul")
}

func GetReadingRoomSeatByCampusBackport(c *fiber.Ctx) error {
	return GetReadingRoomSeatByCampus(c, strings.ToLower(parseCampus(c)) == "seoul")
}

func GetReadingRoomSeatByCampus(c *fiber.Ctx, campus bool) error {
	if campus {
		return c.JSON(ReadingRoomByCampus{OpenedRooms: nil})
	} else {
		return c.JSON(ReadingRoomByCampus{OpenedRooms: library.GetLibrary()})
	}
}

func GetFoodMenuByCampus(c *fiber.Ctx) error {
	return c.JSON(food.GetFoodMenuAll())
}
