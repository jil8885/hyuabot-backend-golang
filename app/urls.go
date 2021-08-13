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
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	busStopList := [5]string{"Residence", "Shuttlecock_O", "Subway", "Terminal", "Shuttlecock_I"}
	response := map[string]ShuttleDepartureByStop{}
	for _, item := range busStopList{
		busForStation, busForTerminal := shuttle.GetShuttle(item, now)
		response[item] = ShuttleDepartureByStop{BusForStation: busForStation, BusForTerminal: busForTerminal}
	}
	return c.JSON(response)
}

func GetShuttleDepartureByStop(c *fiber.Ctx) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	busForStation, busForTerminal := shuttle.GetShuttle(parseShuttleStop(c), now)
	response := ShuttleDepartureByStop{BusForStation: busForStation, BusForTerminal: busForTerminal}
	return c.JSON(response)
}

func GetShuttleStopInfoByStop(c *fiber.Ctx) error {
	// 현재 시간 로딩 (KST)
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)

	busStop := parseShuttleStop(c)
	firstBusForStation, lastBusForStation, firstBusForTerminal, lastBusForTerminal := shuttle.GetFirstLastShuttle(busStop, now)
	weekBusForStation, weekBusForTerminal := shuttle.GetShuttleTimetable(busStop, now, "week")
	weekendBusForStation, weekendBusForTerminal := shuttle.GetShuttleTimetable(busStop, now, "weekend")
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

func GetSubwayDeparture(c *fiber.Ctx) error {
	campus := strings.ToLower(parseCampus(c)) == "seoul"

	now := time.Now()
	_, isWeekends := shuttle.GetDate(now)
	isHoliday := shuttle.IsHoliday(now)

	if campus {
		return c.JSON(SubwayDepartureSeoul{
			Line2: subway.GetRealtimeSubway(1, 1002),
		})
	} else{
		timetableWeekdaysLine4, timetableWeekendsLine4 := subway.GetTimetableSubwayAll(1004)
		timetableWeekdaysLineSuin, timetableWeekendsLineSuin := subway.GetTimetableSubwayAll(1071)

		if isWeekends == "weekend" || isHoliday{
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
	response := Bus{Realtime: map[string][]bus.DepartureItem{"10-1": []bus.DepartureItem{}, "3102": []bus.DepartureItem{}, "707-1": []bus.DepartureItem{}}}
	response.Realtime["707-1"] = bus.GetRealtimeBusDeparture("216000719", "216000070")
	for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList{
		if item.PredictTime1 > 0{
			if item.RouteID == 216000061{
				response.Realtime["3102"] = append(response.Realtime["3102"], bus.DepartureItem{
					Location : item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
				})
			} else if item.RouteID == 216000068{
				response.Realtime["10-1"] = append(response.Realtime["10-1"], bus.DepartureItem{
					Location : item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
				})
			}
			if item.PredictTime2 > 0{
				if item.RouteID == 216000061{
					response.Realtime["3102"] = append(response.Realtime["3102"], bus.DepartureItem{
						Location : item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
					})
				} else if item.RouteID == 216000068{
					response.Realtime["10-1"] = append(response.Realtime["10-1"], bus.DepartureItem{
						Location : item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
					})
				}
			}
		}
	}
	response.Timetable = bus.GetTimetable()
	return c.JSON(response)
}

func GetBusDepartureByLine(c *fiber.Ctx) error {
	routeID := parseBusRouteID(c)
	responseTimetable := bus.GetTimetable()
	if routeID == "10-1"{
		responseRealtimeByStop := bus.GetRealtimeStopDeparture("216000379")
		var responseRealtime []bus.DepartureItem
		for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList{
			if item.RouteID == 216000068{
				if item.PredictTime1 > 0{
					responseRealtime = append(responseRealtime, bus.DepartureItem{
						Location : item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
					})
					if item.PredictTime2 > 0{
						responseRealtime = append(responseRealtime, bus.DepartureItem{
							Location : item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
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
		for _, item := range responseRealtimeByStop.MsgBody.BusArrivalList{
			if item.RouteID == 216000061{
				if item.PredictTime1 > 0{
					responseRealtime = append(responseRealtime, bus.DepartureItem{
						Location : item.LocationNo1, RemainedTime: item.PredictTime1, RemainedSeat: item.RemainSeatCnt1,
					})
					if item.PredictTime2 > 0{
						responseRealtime = append(responseRealtime, bus.DepartureItem{
							Location : item.LocationNo2, RemainedTime: item.PredictTime2, RemainedSeat: item.RemainSeatCnt2,
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
	} else if routeID == "707-1"{
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

func GetBusTimetableByRoute(c *fiber.Ctx) error  {
	routeID := parseBusRouteID(c)
	responseTimetable := bus.GetTimetable()
	if routeID == "10-1"{
		return c.JSON(responseTimetable.Line10_1)
	} else if routeID == "3102" {
		return c.JSON(responseTimetable.Line3102)
	} else if routeID == "707-1"{
		return c.JSON(responseTimetable.Line707_1)
	}
	return c.JSON(BusByRoute{
		Realtime:  []bus.DepartureItem{},
		Timetable: bus.BusTimeTableLine{},
	})}

func GetReadingRoomSeatByCampus(c *fiber.Ctx) error{
	campus := strings.ToLower(parseCampus(c)) == "seoul"

	if campus {
		return c.JSON(ReadingRoomByCampus{OpenedRooms: nil})
	} else{
		return c.JSON(ReadingRoomByCampus{OpenedRooms: library.GetLibrary()})
	}
}

func GetFoodMenuByCampus(c *fiber.Ctx) error  {
	return c.JSON(food.GetFoodMenuAll())
}
