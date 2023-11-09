package v1

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/requests"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
)

func GetShuttleTimetableView(c *fiber.Ctx) error {
	stopQuery := c.Query("stop")
	weekdayQuery := c.Query("weekday")
	periodQuery := c.Query("period")
	startTimeQuery := c.Query("start")
	endTimeQuery := c.Query("end")

	queries := make([]qm.QueryMod, 0)
	if stopQuery != "" {
		queries = append(
			queries,
			models.ShuttleTimetableViewWhere.StopName.EQ(null.StringFrom(stopQuery)),
		)
	}
	if weekdayQuery != "" {
		queries = append(
			queries,
			models.ShuttleTimetableViewWhere.Weekday.EQ(null.BoolFrom(weekdayQuery == "true")),
		)
	}
	if periodQuery != "" {
		queries = append(
			queries,
			models.ShuttleTimetableViewWhere.PeriodType.EQ(null.StringFrom(periodQuery)),
		)
	}
	if startTimeQuery != "" {
		startTime, err := time.Parse("15:04", startTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_TIME"})
		}
		queries = append(
			queries,
			models.ShuttleTimetableViewWhere.DepartureTime.GTE(
				null.TimeFrom(startTime)),
		)
	}
	if endTimeQuery != "" {
		endTime, err := time.Parse("15:04", endTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_TIME"})
		}
		queries = append(
			queries,
			models.ShuttleTimetableViewWhere.DepartureTime.LTE(
				null.TimeFrom(endTime)),
		)
	}

	shuttleTimetableItems, err := models.ShuttleTimetableViews(
		queries...,
	).All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	items := make([]responses.ShuttleTimetableViewItem, 0)
	for _, shuttleTimetableItem := range shuttleTimetableItems {
		items = append(items, responses.ShuttleTimetableViewItem{
			Seq:        shuttleTimetableItem.Seq.Int,
			PeriodType: shuttleTimetableItem.PeriodType.String,
			Weekday:    shuttleTimetableItem.Weekday.Bool,
			Route: responses.ShuttleTimetableRouteItem{
				Name: shuttleTimetableItem.RouteName.String,
				Tag:  shuttleTimetableItem.RouteTag.String,
			},
			StopName:      shuttleTimetableItem.StopName.String,
			DepartureTime: shuttleTimetableItem.DepartureTime.Time.Format("15:04"),
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableViewResponse{Data: items})
}

func GetShuttleTimetableList(c *fiber.Ctx) error {
	weekdayQuery := c.Query("weekday")
	periodQuery := c.Query("period")
	startTimeQuery := c.Query("start")
	endTimeQuery := c.Query("end")

	queries := make([]qm.QueryMod, 0)
	if weekdayQuery != "" {
		queries = append(
			queries,
			models.ShuttleTimetableWhere.Weekday.EQ(weekdayQuery == "true"),
		)
	}
	if periodQuery != "" {
		queries = append(
			queries,
			models.ShuttleTimetableWhere.PeriodType.EQ(periodQuery),
		)
	}
	if startTimeQuery != "" {
		startTime, err := time.Parse("15:04", startTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_TIME"})
		}
		queries = append(
			queries,
			models.ShuttleTimetableWhere.DepartureTime.GTE(startTime),
		)
	}
	if endTimeQuery != "" {
		endTime, err := time.Parse("15:04", endTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_TIME"})
		}
		queries = append(
			queries,
			models.ShuttleTimetableWhere.DepartureTime.LTE(endTime),
		)
	}

	shuttleTimetableItems, err := models.ShuttleTimetables(
		queries...,
	).All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	items := make([]responses.ShuttleTimetableItem, 0)
	for _, shuttleTimetableItem := range shuttleTimetableItems {
		items = append(items, responses.ShuttleTimetableItem{
			Seq:           shuttleTimetableItem.Seq,
			PeriodType:    shuttleTimetableItem.PeriodType,
			Weekday:       shuttleTimetableItem.Weekday,
			Route:         shuttleTimetableItem.RouteName,
			DepartureTime: shuttleTimetableItem.DepartureTime.Format("15:04"),
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableResponse{Data: items})
}

func GetShuttleTimetable(c *fiber.Ctx) error {
	seq, err := c.ParamsInt("seq")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "SEQ_NOT_PROVIDED"})
	}
	exists, err := models.ShuttleTimetableExists(c.Context(), database.DB, seq)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "TIMETABLE_NOT_FOUND"})
	}
	item, err := models.FindShuttleTimetable(c.Context(), database.DB, seq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableItem{
		Seq:           item.Seq,
		PeriodType:    item.PeriodType,
		Weekday:       item.Weekday,
		Route:         item.RouteName,
		DepartureTime: item.DepartureTime.Format("15:04"),
	})
}

func CreateShuttleTimetable(c *fiber.Ctx) error {
	var request requests.ShuttleTimetableRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	departureTime, err := time.Parse("15:04", request.DepartureTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DEPARTURE_TIME"})
	}
	item, err := models.ShuttleTimetables(
		models.ShuttleTimetableWhere.PeriodType.EQ(request.PeriodType),
		models.ShuttleTimetableWhere.Weekday.EQ(request.Weekday),
		models.ShuttleTimetableWhere.RouteName.EQ(request.Route),
		models.ShuttleTimetableWhere.DepartureTime.EQ(departureTime),
	).One(c.Context(), database.DB)
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableItem{
		Seq:           item.Seq,
		PeriodType:    item.PeriodType,
		Weekday:       item.Weekday,
		Route:         item.RouteName,
		DepartureTime: item.DepartureTime.Format("15:04"),
	})
}

func UpdateShuttleTimetable(c *fiber.Ctx) error {
	var request requests.ShuttleTimetableUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	seq, err := c.ParamsInt("seq")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "SEQ_NOT_PROVIDED"})
	}
	exists, err := models.ShuttleTimetableExists(c.Context(), database.DB, seq)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "TIMETABLE_NOT_FOUND"})
	}
	item, err := models.FindShuttleTimetable(c.Context(), database.DB, seq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	if request.PeriodType.Valid {
		item.PeriodType = request.PeriodType.String
	}
	if request.Weekday.Valid {
		item.Weekday = request.Weekday.Bool
	}
	if request.Route.Valid {
		item.RouteName = request.Route.String
	}
	if request.DepartureTime.Valid {
		departureTime, err := time.Parse("15:04", request.DepartureTime.String)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DEPARTURE_TIME"})
		}
		item.DepartureTime = departureTime
	}
	_, err = item.Update(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableItem{
		Seq:           item.Seq,
		PeriodType:    item.PeriodType,
		Weekday:       item.Weekday,
		Route:         item.RouteName,
		DepartureTime: item.DepartureTime.Format("15:04"),
	})
}

func DeleteShuttleTimetable(c *fiber.Ctx) error {
	seq, err := c.ParamsInt("seq")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "SEQ_NOT_PROVIDED"})
	}
	exists, err := models.ShuttleTimetableExists(c.Context(), database.DB, seq)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "TIMETABLE_NOT_FOUND"})
	}
	item, err := models.FindShuttleTimetable(c.Context(), database.DB, seq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	_, err = item.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}

func GetShuttleRouteList(c *fiber.Ctx) error {
	nameQuery := c.Query("name")
	tagQuery := c.Query("tag")
	queries := make([]qm.QueryMod, 0)
	if nameQuery != "" {
		queries = append(
			queries,
			models.ShuttleRouteWhere.RouteName.LIKE(fmt.Sprintf("%%%s%%", nameQuery)),
		)
	}
	if tagQuery != "" {
		queries = append(
			queries,
			models.ShuttleRouteWhere.RouteTag.EQ(null.StringFrom(tagQuery)),
		)
	}
	// Join query
	shuttleRouteItems, err := models.ShuttleRoutes(
		queries...,
	).All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	items := make([]responses.ShuttleRouteItem, 0)
	for _, shuttleRouteItem := range shuttleRouteItems {
		items = append(items, responses.ShuttleRouteItem{
			Name: shuttleRouteItem.RouteName,
			Tag:  shuttleRouteItem.RouteTag.String,
			Description: responses.ShuttleRouteDescription{
				Korean:  shuttleRouteItem.RouteDescriptionKorean.String,
				English: shuttleRouteItem.RouteDescriptionEnglish.String,
			},
			Start: responses.ShuttleStopItem{
				Name:      shuttleRouteItem.R.StartStopShuttleStop.StopName,
				Latitude:  shuttleRouteItem.R.StartStopShuttleStop.Latitude,
				Longitude: shuttleRouteItem.R.StartStopShuttleStop.Longitude,
			},
			End: responses.ShuttleStopItem{
				Name:      shuttleRouteItem.R.EndStopShuttleStop.StopName,
				Latitude:  shuttleRouteItem.R.EndStopShuttleStop.Latitude,
				Longitude: shuttleRouteItem.R.EndStopShuttleStop.Longitude,
			},
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteListResponse{Data: items})
}

func GetShuttleRoute(c *fiber.Ctx) error {
	nameParam := c.Params("route")
	if nameParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_NAME"})
	}
	exists, err := models.ShuttleRouteExists(c.Context(), database.DB, nameParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}
	item, err := models.FindShuttleRoute(c.Context(), database.DB, nameParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	stops := make([]responses.ShuttleRouteStopItem, 0)
	for _, shuttleRouteStop := range item.R.RouteNameShuttleRouteStops {
		timeDelta, err := time.ParseDuration(shuttleRouteStop.CumulativeTime)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		}
		stops = append(stops, responses.ShuttleRouteStopItem{
			Name:           shuttleRouteStop.StopName,
			Seq:            shuttleRouteStop.StopOrder.Int,
			CumulativeTime: timeDelta,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteDetailItem{
		Name: item.RouteName,
		Tag:  item.RouteTag.String,
		Description: responses.ShuttleRouteDescription{
			Korean:  item.RouteDescriptionKorean.String,
			English: item.RouteDescriptionEnglish.String,
		},
		Start: responses.ShuttleStopItem{
			Name:      item.R.StartStopShuttleStop.StopName,
			Latitude:  item.R.StartStopShuttleStop.Latitude,
			Longitude: item.R.StartStopShuttleStop.Longitude,
		},
		End: responses.ShuttleStopItem{
			Name:      item.R.EndStopShuttleStop.StopName,
			Latitude:  item.R.EndStopShuttleStop.Latitude,
			Longitude: item.R.EndStopShuttleStop.Longitude,
		},
		Stops: stops,
	})
}

func CreateShuttleRoute(c *fiber.Ctx) error {
	var request requests.ShuttleRouteRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	exists, err := models.ShuttleRouteExists(c.Context(), database.DB, request.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if exists {
		return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{Message: "ROUTE_ALREADY_EXISTS"})
	}

	startStop, err := models.FindShuttleStop(c.Context(), database.DB, request.Start)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if startStop == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "START_STOP_NOT_FOUND"})
	}

	endStop, err := models.FindShuttleStop(c.Context(), database.DB, request.End)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if endStop == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "END_STOP_NOT_FOUND"})
	}

	route := models.ShuttleRoute{
		RouteName:               request.Name,
		RouteTag:                null.StringFrom(request.Tag),
		RouteDescriptionKorean:  null.StringFrom(request.DescriptionKorean),
		RouteDescriptionEnglish: null.StringFrom(request.DescriptionEnglish),
		StartStop:               null.StringFrom(request.Start),
		EndStop:                 null.StringFrom(request.End),
	}
	err = route.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusCreated).JSON(responses.ShuttleRouteItem{
		Name: request.Name,
		Tag:  request.Tag,
		Description: responses.ShuttleRouteDescription{
			Korean:  request.DescriptionKorean,
			English: request.DescriptionEnglish,
		},
		Start: responses.ShuttleStopItem{
			Name:      request.Start,
			Latitude:  startStop.Latitude,
			Longitude: startStop.Longitude,
		},
		End: responses.ShuttleStopItem{
			Name:      request.End,
			Latitude:  endStop.Latitude,
			Longitude: endStop.Longitude,
		},
	})
}

func UpdateShuttleRoute(c *fiber.Ctx) error {
	var request requests.ShuttleRouteUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	nameParam := c.Params("route")
	if nameParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_NAME"})
	}
	exists, err := models.ShuttleRouteExists(c.Context(), database.DB, nameParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}
	item, err := models.FindShuttleRoute(c.Context(), database.DB, nameParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	if request.Tag.Valid {
		item.RouteTag = request.Tag
	}
	if request.DescriptionKorean.Valid {
		item.RouteDescriptionKorean = request.DescriptionKorean
	}
	if request.DescriptionEnglish.Valid {
		item.RouteDescriptionEnglish = request.DescriptionEnglish
	}
	if request.Start.Valid {
		startStop, err := models.FindShuttleStop(c.Context(), database.DB, request.Start.String)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		} else if startStop == nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "START_STOP_NOT_FOUND"})
		}
	}
	if request.End.Valid {
		endStop, err := models.FindShuttleStop(c.Context(), database.DB, request.End.String)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		} else if endStop == nil {
			return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "END_STOP_NOT_FOUND"})
		}
	}
	_, err = item.Update(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteItem{
		Name: item.RouteName,
		Tag:  item.RouteTag.String,
		Description: responses.ShuttleRouteDescription{
			Korean:  item.RouteDescriptionKorean.String,
			English: item.RouteDescriptionEnglish.String,
		},
		Start: responses.ShuttleStopItem{
			Name:      item.R.StartStopShuttleStop.StopName,
			Latitude:  item.R.StartStopShuttleStop.Latitude,
			Longitude: item.R.StartStopShuttleStop.Longitude,
		},
		End: responses.ShuttleStopItem{
			Name:      item.R.EndStopShuttleStop.StopName,
			Latitude:  item.R.EndStopShuttleStop.Latitude,
			Longitude: item.R.EndStopShuttleStop.Longitude,
		},
	})
}

func DeleteShuttleRoute(c *fiber.Ctx) error {
	var request requests.ShuttleRouteUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	nameParam := c.Params("route")
	if nameParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_NAME"})
	}
	exists, err := models.ShuttleRouteExists(c.Context(), database.DB, nameParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}
	item, err := models.FindShuttleRoute(c.Context(), database.DB, nameParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	_, err = models.ShuttleRouteStops(models.ShuttleRouteStopWhere.RouteName.EQ(item.RouteName)).DeleteAll(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	_, err = models.ShuttleTimetables(models.ShuttleTimetableWhere.RouteName.EQ(item.RouteName)).DeleteAll(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	_, err = item.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}

func GetShuttleStopList(c *fiber.Ctx) error {
	stops, err := models.ShuttleStops().All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	items := make([]responses.ShuttleStopItem, 0)
	for _, stop := range stops {
		items = append(items, responses.ShuttleStopItem{
			Name:      stop.StopName,
			Latitude:  stop.Latitude,
			Longitude: stop.Longitude,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleStopListResponse{Data: items})
}

func GetShuttleStop(c *fiber.Ctx) error {
	nameParam := c.Params("stop")
	if nameParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_STOP_NAME"})
	}
	exists, err := models.ShuttleStopExists(c.Context(), database.DB, nameParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}
	item, err := models.FindShuttleStop(c.Context(), database.DB, nameParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleStopItem{
		Name:      item.StopName,
		Latitude:  item.Latitude,
		Longitude: item.Longitude,
	})

}

func CreateShuttleStop(c *fiber.Ctx) error {
	var request requests.ShuttleStopRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	exists, err := models.ShuttleStopExists(c.Context(), database.DB, request.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if exists {
		return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{Message: "STOP_ALREADY_EXISTS"})
	}
	stop := models.ShuttleStop{
		StopName:  request.Name,
		Latitude:  null.Float64From(request.Latitude),
		Longitude: null.Float64From(request.Longitude),
	}
	err = stop.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusCreated).JSON(responses.ShuttleStopItem{
		Name:      stop.StopName,
		Latitude:  stop.Latitude,
		Longitude: stop.Longitude,
	})
}

func UpdateShuttleStop(c *fiber.Ctx) error {
	stopParam := c.Params("stop")
	if stopParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_STOP_NAME"})
	}
	exists, err := models.ShuttleStopExists(c.Context(), database.DB, stopParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}
	item, err := models.FindShuttleStop(c.Context(), database.DB, stopParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	var request requests.ShuttleStopUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	if request.Latitude.Valid {
		item.Latitude = request.Latitude
	}
	if request.Longitude.Valid {
		item.Longitude = request.Longitude
	}
	_, err = item.Update(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleStopItem{
		Name:      item.StopName,
		Latitude:  item.Latitude,
		Longitude: item.Longitude,
	})
}

func DeleteShuttleStop(c *fiber.Ctx) error {
	stopParam := c.Params("stop")
	if stopParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_STOP_NAME"})
	}
	exists, err := models.ShuttleStopExists(c.Context(), database.DB, stopParam)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}
	item, err := models.FindShuttleStop(c.Context(), database.DB, stopParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	for _, shuttleRouteStop := range item.R.StopNameShuttleRouteStops {
		_, err = shuttleRouteStop.Delete(c.Context(), database.DB)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		}
	}
	_, err = item.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}

func GetShuttleRouteStopList(c *fiber.Ctx) error {
	routeParam := c.Params("route")
	if routeParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_NAME"})
	}
	route, err := models.FindShuttleRoute(c.Context(), database.DB, routeParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if route == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}

	items := make([]responses.ShuttleRouteStopItem, 0)
	for _, shuttleRouteStop := range route.R.RouteNameShuttleRouteStops {
		timeDelta, err := time.ParseDuration(shuttleRouteStop.CumulativeTime)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
		}
		items = append(items, responses.ShuttleRouteStopItem{
			Name:           shuttleRouteStop.StopName,
			Seq:            shuttleRouteStop.StopOrder.Int,
			CumulativeTime: timeDelta,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteStopListResponse{Data: items})
}

func GetShuttleRouteStop(c *fiber.Ctx) error {
	routeParam := c.Params("route")
	stopParam := c.Params("stop")
	if routeParam == "" || stopParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_OR_STOP"})
	}
	route, err := models.FindShuttleRoute(c.Context(), database.DB, routeParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if route == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}

	stop, err := models.FindShuttleRouteStop(c.Context(), database.DB, routeParam, stopParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if stop == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}
	duration, err := time.ParseDuration(stop.CumulativeTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteStopItem{
		Name:           stop.StopName,
		Seq:            stop.StopOrder.Int,
		CumulativeTime: duration,
	})
}

func CreateShuttleRouteStop(c *fiber.Ctx) error {
	routeParam := c.Params("route")
	if routeParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_NAME"})
	}
	route, err := models.FindShuttleRoute(c.Context(), database.DB, routeParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if route == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}

	var request requests.ShuttleRouteStopRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	item, err := models.ShuttleRouteStops(
		models.ShuttleRouteStopWhere.RouteName.EQ(routeParam),
		models.ShuttleRouteStopWhere.StopName.EQ(request.Name),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item != nil {
		return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{Message: "STOP_ALREADY_EXISTS"})
	}

	stop := models.ShuttleRouteStop{
		RouteName:      routeParam,
		StopName:       request.Name,
		StopOrder:      null.IntFrom(request.Seq),
		CumulativeTime: request.CumulativeTime,
	}
	err = stop.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	duration, err := time.ParseDuration(stop.CumulativeTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusCreated).JSON(responses.ShuttleRouteStopItem{
		Name:           stop.StopName,
		Seq:            stop.StopOrder.Int,
		CumulativeTime: duration,
	})
}

func UpdateShuttleRouteStop(c *fiber.Ctx) error {
	routeParam := c.Params("route")
	stopParam := c.Params("stop")
	if routeParam == "" || stopParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_OR_STOP"})
	}
	route, err := models.FindShuttleRoute(c.Context(), database.DB, routeParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if route == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}

	var request requests.ShuttleRouteStopUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	item, err := models.FindShuttleRouteStop(c.Context(), database.DB, routeParam, stopParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if err == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}

	if request.Seq.Valid {
		item.StopOrder = request.Seq
	}
	if request.CumulativeTime.Valid {
		_, err := time.ParseDuration(request.CumulativeTime.String)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_CUMULATIVE_TIME"})
		}
		item.CumulativeTime = request.CumulativeTime.String
	}
	_, err = item.Update(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	duration, err := time.ParseDuration(item.CumulativeTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleRouteStopItem{
		Name:           item.StopName,
		Seq:            item.StopOrder.Int,
		CumulativeTime: duration,
	})
}

func DeleteShuttleRouteStop(c *fiber.Ctx) error {
	routeParam := c.Params("route")
	stopParam := c.Params("stop")
	if routeParam == "" || stopParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_ROUTE_OR_STOP"})
	}
	route, err := models.FindShuttleRoute(c.Context(), database.DB, routeParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if route == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "ROUTE_NOT_FOUND"})
	}

	stop, err := models.FindShuttleRouteStop(c.Context(), database.DB, routeParam, stopParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if stop == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "STOP_NOT_FOUND"})
	}
	_, err = stop.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}

func GetShuttlePeriodList(c *fiber.Ctx) error {
	items, err := models.ShuttlePeriods().All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	periods := make([]responses.ShuttlePeriodItem, 0)
	for _, item := range items {
		periods = append(periods, responses.ShuttlePeriodItem{
			StartDate:  item.PeriodStart.Format("2006-01-02"),
			EndDate:    item.PeriodEnd.Format("2006-01-02"),
			PeriodType: item.PeriodType,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttlePeriodListResponse{Data: periods})
}

func GetShuttlePeriod(c *fiber.Ctx) error {
	startParam := c.Params("start")
	endParam := c.Params("end")
	if startParam == "" || endParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_PERIOD"})
	}
	startDate, err := time.Parse("2006-01-02", startParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_DATE"})
	}
	endDate, err := time.Parse("2006-01-02", endParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_DATE"})
	}
	item, err := models.ShuttlePeriods(
		models.ShuttlePeriodWhere.PeriodStart.EQ(startDate),
		models.ShuttlePeriodWhere.PeriodEnd.EQ(endDate),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "PERIOD_NOT_FOUND"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttlePeriodItem{
		StartDate:  item.PeriodStart.Format("2006-01-02"),
		EndDate:    item.PeriodEnd.Format("2006-01-02"),
		PeriodType: item.PeriodType,
	})
}

func CreateShuttlePeriod(c *fiber.Ctx) error {
	var request requests.ShuttlePeriodRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}
	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_DATE"})
	}
	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_DATE"})
	}
	item, err := models.ShuttlePeriods(
		models.ShuttlePeriodWhere.PeriodStart.EQ(startDate),
		models.ShuttlePeriodWhere.PeriodEnd.EQ(endDate),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item != nil {
		return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{Message: "PERIOD_ALREADY_EXISTS"})
	}
	period := models.ShuttlePeriod{
		PeriodStart: startDate,
		PeriodEnd:   endDate,
		PeriodType:  request.PeriodType,
	}
	err = period.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusCreated).JSON(responses.ShuttlePeriodItem{
		StartDate:  period.PeriodStart.Format("2006-01-02"),
		EndDate:    period.PeriodEnd.Format("2006-01-02"),
		PeriodType: period.PeriodType,
	})
}

func DeleteShuttlePeriod(c *fiber.Ctx) error {
	startDateQuery := c.Query("start")
	endDateQuery := c.Query("end")
	start, err := time.Parse("2006-01-02", startDateQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_DATE"})
	}
	end, err := time.Parse("2006-01-02", endDateQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_DATE"})
	}
	_, err = models.ShuttlePeriods(
		models.ShuttlePeriodWhere.PeriodStart.EQ(start),
		models.ShuttlePeriodWhere.PeriodEnd.EQ(end),
	).DeleteAll(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}

func GetShuttleHolidayList(c *fiber.Ctx) error {
	startDateQuery := c.Query("start")
	endDateQuery := c.Query("end")
	targetDateQuery := c.Query("date")
	calendarQuery := c.Query("calendar")
	typeQuery := c.Query("type")

	queries := make([]qm.QueryMod, 0)
	if (startDateQuery != "" || endDateQuery != "") && targetDateQuery != "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "CONFLICTING_DATE_PARAMS"})
	}
	if startDateQuery != "" {
		startDate, err := time.Parse("2006-01-02", startDateQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_DATE"})
		}
		queries = append(
			queries,
			models.ShuttleHolidayWhere.HolidayDate.GTE(startDate),
		)
	}
	if endDateQuery != "" {
		endDate, err := time.Parse("2006-01-02", endDateQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_DATE"})
		}
		queries = append(
			queries,
			models.ShuttleHolidayWhere.HolidayDate.LTE(endDate),
		)
	}
	if targetDateQuery != "" {
		targetDate, err := time.Parse("2006-01-02", targetDateQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DATE"})
		}
		queries = append(
			queries,
			models.ShuttleHolidayWhere.HolidayDate.EQ(targetDate),
		)
	}
	if calendarQuery != "" {
		queries = append(
			queries,
			models.ShuttleHolidayWhere.CalendarType.EQ(calendarQuery),
		)
	}
	if typeQuery != "" {
		queries = append(
			queries,
			models.ShuttleHolidayWhere.HolidayType.EQ(typeQuery),
		)
	}

	shuttleHolidayItems, err := models.ShuttleHolidays(
		queries...,
	).All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}

	items := make([]responses.ShuttleHolidayItem, 0)
	for _, shuttleHolidayItem := range shuttleHolidayItems {
		items = append(items, responses.ShuttleHolidayItem{
			CalendarType: shuttleHolidayItem.CalendarType,
			HolidayType:  shuttleHolidayItem.HolidayType,
			HolidayDate:  shuttleHolidayItem.HolidayDate.Format("2006-01-02"),
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleHolidayListResponse{Data: items})
}

func GetShuttleHoliday(c *fiber.Ctx) error {
	dateQuery := c.Params("date")
	date, err := time.Parse("2006-01-02", dateQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DATE"})
	}
	item, err := models.ShuttleHolidays(
		models.ShuttleHolidayWhere.HolidayDate.EQ(date),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "NOT_FOUND"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleHolidayItem{
		CalendarType: item.CalendarType,
		HolidayType:  item.HolidayType,
		HolidayDate:  item.HolidayDate.Format("2006-01-02"),
	})
}

func CreateShuttleHoliday(c *fiber.Ctx) error {
	var request requests.ShuttleHolidayRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_JSON_PROVIDED"})
	}

	date, err := time.Parse("2006-01-02", request.HolidayDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DATE"})
	}

	item, err := models.ShuttleHolidays(
		models.ShuttleHolidayWhere.HolidayDate.EQ(date),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item != nil {
		return c.Status(fiber.StatusConflict).JSON(responses.ErrorResponse{Message: "HOLIDAY_ALREADY_EXISTS"})
	}

	newItem := models.ShuttleHoliday{
		CalendarType: request.CalendarType,
		HolidayType:  request.HolidayType,
		HolidayDate:  date,
	}
	err = newItem.Insert(c.Context(), database.DB, boil.Infer())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusCreated).JSON(responses.ShuttleHolidayItem{
		CalendarType: newItem.CalendarType,
		HolidayType:  newItem.HolidayType,
		HolidayDate:  newItem.HolidayDate.Format("2006-01-02"),
	})
}

func DeleteShuttleHoliday(c *fiber.Ctx) error {
	dateParam := c.Params("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_DATE"})
	}

	item, err := models.ShuttleHolidays(
		models.ShuttleHolidayWhere.HolidayDate.EQ(date),
	).One(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	} else if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.ErrorResponse{Message: "NOT_FOUND"})
	}

	_, err = item.Delete(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{Message: "INTERNAL_SERVER_ERROR"})
	}
	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "DELETED"})
}
