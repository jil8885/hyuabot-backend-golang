package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
)

func GetShuttleTimetableView(c *fiber.Ctx) error {
	stopQuery := c.Query("stop")
	weekdayQuery := c.Query("weekday")
	periodQuery := c.Query("period")
	startTimeQuery := c.Query("start")
	endTimeQuery := c.Query("end")

	whereQueries := make([]qm.QueryMod, 0)
	if stopQuery != "" {
		whereQueries = append(
			whereQueries,
			models.ShuttleTimetableViewWhere.StopName.EQ(null.StringFrom(stopQuery)),
		)
	}
	if weekdayQuery != "" {
		whereQueries = append(
			whereQueries,
			models.ShuttleTimetableViewWhere.Weekday.EQ(null.BoolFrom(weekdayQuery == "true")),
		)
	}
	if periodQuery != "" {
		whereQueries = append(
			whereQueries,
			models.ShuttleTimetableViewWhere.PeriodType.EQ(null.StringFrom(periodQuery)),
		)
	}
	if startTimeQuery != "" {
		startTime, err := time.Parse("15:04", startTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_START_TIME"})
		}
		whereQueries = append(
			whereQueries,
			models.ShuttleTimetableViewWhere.DepartureTime.GTE(
				null.TimeFrom(startTime)),
		)
	}
	if endTimeQuery != "" {
		endTime, err := time.Parse("15:04", endTimeQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{Message: "INVALID_END_TIME"})
		}
		whereQueries = append(
			whereQueries,
			models.ShuttleTimetableViewWhere.DepartureTime.LTE(
				null.TimeFrom(endTime)),
		)
	}

	shuttleTimetableItems, err := models.ShuttleTimetableViews(
		whereQueries...,
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
			Route: responses.ShuttleRouteItem{
				Name: shuttleTimetableItem.RouteName.String,
				Tag:  shuttleTimetableItem.RouteTag.String,
			},
			StopName:      shuttleTimetableItem.StopName.String,
			DepartureTime: shuttleTimetableItem.DepartureTime.Time.Format("15:04"),
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses.ShuttleTimetableViewResponse{Data: items})
}
