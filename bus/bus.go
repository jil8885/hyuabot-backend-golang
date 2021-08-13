package bus

func GetBusDepartureInfo() ([]DepartureItem, StopResponse, BusTimeTableJson) {
	line707Realtime := GetRealtimeBusDeparture("216000719", "216000070")
	guestHouseRealtime := GetRealtimeStopDeparture("216000379")
	timetable := GetTimetable()

	return line707Realtime, guestHouseRealtime, timetable
}