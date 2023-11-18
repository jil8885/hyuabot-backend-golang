package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/golang-module/carbon/v2"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var shuttleRouteFactory = factory.NewFactory(
	&models.ShuttleRoute{
		RouteTag:                null.StringFrom(randomdata.Noun()),
		RouteDescriptionKorean:  null.StringFrom(randomdata.Alphanumeric(100)),
		RouteDescriptionEnglish: null.StringFrom(randomdata.Alphanumeric(100)),
		StartStop:               null.StringFrom(randomdata.Noun()),
		EndStop:                 null.StringFrom(randomdata.Noun()),
	},
)

var shuttleStopFactory = factory.NewFactory(
	&models.ShuttleStop{
		Latitude:  null.Float64From(randomdata.Decimal(-180.0, 180.0)),
		Longitude: null.Float64From(randomdata.Decimal(-90.0, 90.0)),
	},
)

var shuttlePeriodType = []string{"semester", "vacation", "vacation_session"}
var shuttlePeriodFactory = factory.NewFactory(
	&models.ShuttlePeriod{
		PeriodType:  shuttlePeriodType[randomdata.Number(0, 2)],
		PeriodStart: carbon.Parse(randomdata.FullDate()).ToStdTime(),
		PeriodEnd:   carbon.Parse(randomdata.FullDate()).ToStdTime(),
	},
)

var shuttleHolidayType = []string{"weekends", "halt"}
var shuttleHolidayFactory = factory.NewFactory(
	&models.ShuttleHoliday{
		CalendarType: "solar",
		HolidayType:  shuttleHolidayType[randomdata.Number(0, 1)],
		HolidayDate:  carbon.Parse(randomdata.FullDate()).ToStdTime(),
	},
)

var shuttleTimetableFactory = factory.NewFactory(
	&models.ShuttleTimetable{
		PeriodType:    shuttlePeriodType[randomdata.Number(0, 2)],
		Weekday:       randomdata.Boolean(),
		RouteName:     randomdata.Noun(),
		DepartureTime: time.Date(0, 0, 0, randomdata.Number(0, 23), randomdata.Number(0, 59), 0, 0, time.Local),
	},
)

func shuttleDataCreate() (
	[]*models.ShuttleStop,
	[]*models.ShuttleRoute,
	[]*models.ShuttleTimetable,
	[]*models.ShuttlePeriod,
	[]*models.ShuttleHoliday,
) {
	stopList := make([]*models.ShuttleStop, 0)
	routeList := make([]*models.ShuttleRoute, 0)
	timetableList := make([]*models.ShuttleTimetable, 0)
	periodList := make([]*models.ShuttlePeriod, 0)
	holidayList := make([]*models.ShuttleHoliday, 0)

	for i := 0; i < 3; i++ {
		period := models.ShuttlePeriodType{PeriodType: shuttlePeriodType[i]}
		err := period.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < 10; i++ {
		stop, _ := shuttleStopFactory.MustCreateWithOption(map[string]interface{}{
			"StopName": randomdata.Alphanumeric(6),
		}).(*models.ShuttleStop)
		err := stop.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		stopList = append(stopList, stop)
	}
	for i := 0; i < 10; i++ {
		route, _ := shuttleRouteFactory.MustCreateWithOption(map[string]interface{}{
			"RouteName": randomdata.Alphanumeric(6),
			"StartStop": null.StringFrom(stopList[randomdata.Number(0, 8)].StopName),
			"EndStop":   null.StringFrom(stopList[randomdata.Number(0, 8)].StopName),
		}).(*models.ShuttleRoute)
		err := route.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		routeList = append(routeList, route)
	}
	for i := 0; i < 9; i++ {
		routeStop := models.ShuttleRouteStop{
			RouteName:      routeList[0].RouteName,
			StopName:       stopList[i].StopName,
			StopOrder:      null.IntFrom(i),
			CumulativeTime: fmt.Sprintf("00:%02d:00", i*5),
		}
		err := routeStop.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
	}
	for i := 0; i < 100; i++ {
		timetable := shuttleTimetableFactory.MustCreateWithOption(map[string]interface{}{
			"PeriodType":    shuttlePeriodType[randomdata.Number(0, 2)],
			"Weekday":       randomdata.Boolean(),
			"RouteName":     routeList[randomdata.Number(0, 9)].RouteName,
			"DepartureTime": time.Date(0, 0, 0, randomdata.Number(0, 23), randomdata.Number(0, 59), 0, 0, time.Local),
		}).(*models.ShuttleTimetable)
		err := timetable.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		timetableList = append(timetableList, timetable)
	}
	for i := 0; i < 9; i++ {
		start := carbon.Parse(fmt.Sprintf(
			"%d-%d-%d",
			randomdata.Number(2010, 2020),
			randomdata.Number(1, 12),
			randomdata.Number(1, 28),
		)).SetTimezone(carbon.Seoul).SetTime(0, 0, 0)
		end := start.AddMonths(randomdata.Number(1, 3)).SetTime(23, 59, 59)
		period := shuttlePeriodFactory.MustCreateWithOption(map[string]interface{}{
			"PeriodStart": start.SetTimezone(carbon.UTC).ToStdTime(),
			"PeriodEnd":   end.SetTimezone(carbon.UTC).ToStdTime(),
		}).(*models.ShuttlePeriod)
		err := period.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		periodList = append(periodList, period)
	}
	for i := 0; i < 9; i++ {
		holiday := shuttleHolidayFactory.MustCreateWithOption(map[string]interface{}{
			"HolidayDate": carbon.Parse(fmt.Sprintf(
				"%d-%d-%d",
				randomdata.Number(2010, 2020),
				randomdata.Number(1, 12),
				randomdata.Number(1, 28),
			)).ToStdTime(),
		}).(*models.ShuttleHoliday)
		err := holiday.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		holidayList = append(holidayList, holiday)
	}
	return stopList, routeList, timetableList, periodList, holidayList
}

func TestGetShuttleRouteList(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request all shuttle routes
	req := httptest.NewRequest("GET", "/api/v1/shuttle/route", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleRouteListResponse responses.ShuttleRouteListResponse
	_ = json.NewDecoder(response.Body).Decode(&shuttleRouteListResponse)
	test.NotEmpty(shuttleRouteListResponse)
	for _, shuttleRoute := range shuttleRouteListResponse.Data {
		test.NotEmpty(shuttleRoute.Tag)
		test.NotEmpty(shuttleRoute.Name)
		test.NotEmpty(shuttleRoute.Description.English)
		test.NotEmpty(shuttleRoute.Description.Korean)
		test.NotEmpty(shuttleRoute.Start)
		test.NotEmpty(shuttleRoute.End)
	}
	// Request shuttle routes with name query
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/route?name=%s", routeList[0].RouteName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	_ = json.NewDecoder(response.Body).Decode(&shuttleRouteListResponse)
	test.NotEmpty(shuttleRouteListResponse)
	for _, shuttleRoute := range shuttleRouteListResponse.Data {
		test.Contains(shuttleRoute.Name, shuttleRouteListResponse.Data[0].Name)
		test.NotEmpty(shuttleRoute.Tag)
		test.NotEmpty(shuttleRoute.Name)
		test.NotEmpty(shuttleRoute.Description.English)
		test.NotEmpty(shuttleRoute.Description.Korean)
		test.NotEmpty(shuttleRoute.Start)
		test.NotEmpty(shuttleRoute.End)
	}
	tearDownDatabase()
}

func TestGetShuttleRouteItem(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle route
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/route/%s", routeList[0].RouteName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleRouteItemResponse responses.ShuttleRouteDetailItem
	_ = json.NewDecoder(response.Body).Decode(&shuttleRouteItemResponse)
	test.NotEmpty(shuttleRouteItemResponse.Name)
	test.NotEmpty(shuttleRouteItemResponse.Tag)
	test.NotEmpty(shuttleRouteItemResponse.Description.English)
	test.NotEmpty(shuttleRouteItemResponse.Description.Korean)
	test.NotEmpty(shuttleRouteItemResponse.Start.Name)
	test.NotEmpty(shuttleRouteItemResponse.Start.Latitude)
	test.NotEmpty(shuttleRouteItemResponse.Start.Longitude)
	test.NotEmpty(shuttleRouteItemResponse.End.Name)
	test.NotEmpty(shuttleRouteItemResponse.End.Latitude)
	test.NotEmpty(shuttleRouteItemResponse.End.Longitude)
	for _, stop := range shuttleRouteItemResponse.Stops {
		test.NotEmpty(stop.Name)
		test.NotEmpty(stop.CumulativeTime)
		test.GreaterOrEqual(stop.Seq, 0)
	}
	// Request shuttle route with invalid tag
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/route/%s", "invalid_tag"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	tearDownDatabase()
}

func TestCreateShuttleRoute(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		Tag                string `json:"tag"`
		Name               string `json:"name"`
		DescriptionKorean  string `json:"descriptionKorean"`
		DescriptionEnglish string `json:"descriptionEnglish"`
		Start              string `json:"start"`
		End                string `json:"end"`
	}{
		{
			Tag:                "TEST",
			Name:               routeList[0].RouteName,
			DescriptionKorean:  routeList[0].RouteDescriptionKorean.String,
			DescriptionEnglish: routeList[0].RouteDescriptionEnglish.String,
			Start:              routeList[0].StartStop.String,
			End:                routeList[0].EndStop.String,
		},
		{
			Tag:                "TEST",
			Name:               randomdata.Alphanumeric(6),
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              randomdata.Noun(),
			End:                stopList[randomdata.Number(0, 9)].StopName,
		},
		{
			Tag:                "TEST",
			Name:               randomdata.Alphanumeric(6),
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              stopList[randomdata.Number(0, 9)].StopName,
			End:                randomdata.Noun(),
		},
		{
			Tag:                "TEST",
			Name:               randomdata.Alphanumeric(6),
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              stopList[randomdata.Number(0, 9)].StopName,
			End:                stopList[randomdata.Number(0, 9)].StopName,
		},
	}
	expectedStatusCodes := []int{409, 400, 400, 201}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/shuttle/route", bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		if response.StatusCode != 201 {
			var responseError responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&responseError)
			test.NotEmpty(responseError.Message)
			if index == 0 {
				test.Equal("ROUTE_ALREADY_EXISTS", responseError.Message)
			} else if index == 1 {
				test.Equal("START_STOP_NOT_FOUND", responseError.Message)
			} else if index == 2 {
				test.Equal("END_STOP_NOT_FOUND", responseError.Message)
			}
		} else {
			var responseShuttleRoute responses.ShuttleRouteDetailItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttleRoute)
			test.NotEmpty(responseShuttleRoute.Name)
			test.NotEmpty(responseShuttleRoute.Tag)
			test.NotEmpty(responseShuttleRoute.Description.English)
			test.NotEmpty(responseShuttleRoute.Description.Korean)
			test.NotEmpty(responseShuttleRoute.Start.Name)
			test.NotEmpty(responseShuttleRoute.Start.Latitude)
			test.NotEmpty(responseShuttleRoute.Start.Longitude)
			test.NotEmpty(responseShuttleRoute.End.Name)
			test.NotEmpty(responseShuttleRoute.End.Latitude)
			test.NotEmpty(responseShuttleRoute.End.Longitude)
		}
	}
	tearDownDatabase()
}

func TestUpdateShuttleRoute(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		Tag                string `json:"tag"`
		Name               string `json:"name"`
		DescriptionKorean  string `json:"descriptionKorean"`
		DescriptionEnglish string `json:"descriptionEnglish"`
		Start              string `json:"start"`
		End                string `json:"end"`
	}{
		{
			Tag:                "TEST",
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              randomdata.Noun(),
			End:                stopList[randomdata.Number(0, 9)].StopName,
		},
		{
			Tag:                "TEST",
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              stopList[randomdata.Number(0, 9)].StopName,
			End:                randomdata.Noun(),
		},
		{
			Tag:                "TEST",
			DescriptionKorean:  randomdata.Alphanumeric(100),
			DescriptionEnglish: randomdata.Alphanumeric(100),
			Start:              stopList[randomdata.Number(0, 9)].StopName,
			End:                stopList[randomdata.Number(0, 9)].StopName,
		},
	}
	expectedStatusCodes := []int{400, 400, 200}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/shuttle/route/%s", routeList[0].RouteName), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)
		if response.StatusCode != 200 {
			var responseError responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&responseError)
			test.NotEmpty(responseError.Message)
			if index == 0 {
				test.Equal("START_STOP_NOT_FOUND", responseError.Message)
			} else if index == 1 {
				test.Equal("END_STOP_NOT_FOUND", responseError.Message)
			}
		} else {
			var responseShuttleRoute responses.ShuttleRouteDetailItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttleRoute)
			test.NotEmpty(responseShuttleRoute.Name)
			test.NotEmpty(responseShuttleRoute.Tag)
			test.NotEmpty(responseShuttleRoute.Description.English)
			test.NotEmpty(responseShuttleRoute.Description.Korean)
			test.NotEmpty(responseShuttleRoute.Start.Name)
			test.NotEmpty(responseShuttleRoute.Start.Latitude)
			test.NotEmpty(responseShuttleRoute.Start.Longitude)
			test.NotEmpty(responseShuttleRoute.End.Name)
			test.NotEmpty(responseShuttleRoute.End.Latitude)
			test.NotEmpty(responseShuttleRoute.End.Longitude)
		}

		req = httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/shuttle/route/%s", "invalid_tag"), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ = app.Test(req, 5000)
		test.Equal(404, response.StatusCode)
		test.Equal("application/json", response.Header.Get("Content-Type"))
		var responseError responses.ErrorResponse
		_ = json.NewDecoder(response.Body).Decode(&responseError)
		test.NotEmpty(responseError.Message)
		test.Equal("ROUTE_NOT_FOUND", responseError.Message)
	}
	tearDownDatabase()
}

func TestDeleteShuttleRoute(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request delete shuttle route
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/route/%s", routeList[0].RouteName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(204, response.StatusCode)
	// Request delete shuttle route with invalid tag
	req = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/route/%s", "invalid_tag"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("ROUTE_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestGetShuttleStopList(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request all shuttle stops
	req := httptest.NewRequest("GET", "/api/v1/shuttle/stop", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleStopListResponse responses.ShuttleStopListResponse
	_ = json.NewDecoder(response.Body).Decode(&shuttleStopListResponse)
	test.NotEmpty(shuttleStopListResponse.Data)
	for _, shuttleRoute := range shuttleStopListResponse.Data {
		test.NotEmpty(shuttleRoute.Name)
		test.NotEmpty(shuttleRoute.Latitude)
		test.NotEmpty(shuttleRoute.Longitude)
	}
	tearDownDatabase()
}

func TestGetShuttleStopItem(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, _, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle stop
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/stop/%s", stopList[0].StopName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleStopItemResponse responses.ShuttleStopItem
	_ = json.NewDecoder(response.Body).Decode(&shuttleStopItemResponse)
	test.NotEmpty(shuttleStopItemResponse.Name)
	test.NotEmpty(shuttleStopItemResponse.Latitude)
	test.NotEmpty(shuttleStopItemResponse.Longitude)
	// Request shuttle stop with invalid name
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/stop/%s", "invalid_name"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	tearDownDatabase()
}

func TestCreateShuttleStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, _, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		{
			Name:      stopList[0].StopName,
			Latitude:  stopList[0].Latitude.Float64,
			Longitude: stopList[0].Longitude.Float64,
		},
		{
			Name:      randomdata.Alphanumeric(6),
			Latitude:  randomdata.Decimal(-90.0, 90.0),
			Longitude: randomdata.Decimal(-180.0, 180.0),
		},
	}
	expectedStatusCodes := []int{409, 201}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/shuttle/stop", bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		if response.StatusCode == 201 {
			var responseShuttleStop responses.ShuttleStopItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttleStop)
			test.NotEmpty(responseShuttleStop.Name)
			test.NotEmpty(responseShuttleStop.Latitude)
			test.NotEmpty(responseShuttleStop.Longitude)
		}
	}
	tearDownDatabase()
}

func TestUpdateShuttleStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, _, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		{
			Latitude:  randomdata.Decimal(-90.0, 90.0),
			Longitude: randomdata.Decimal(-180.0, 180.0),
		},
	}
	expectedStatusCodes := []int{200}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/shuttle/stop/%s", stopList[0].StopName), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		var responseShuttleStop responses.ShuttleStopItem
		_ = json.NewDecoder(response.Body).Decode(&responseShuttleStop)
		test.NotEmpty(responseShuttleStop.Name)
		test.NotEmpty(responseShuttleStop.Latitude)
		test.NotEmpty(responseShuttleStop.Longitude)

		req = httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/shuttle/stop/%s", "invalid_tag"), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ = app.Test(req, 5000)
		test.Equal(404, response.StatusCode)
		test.Equal("application/json", response.Header.Get("Content-Type"))
		var responseError responses.ErrorResponse
		_ = json.NewDecoder(response.Body).Decode(&responseError)
		test.NotEmpty(responseError.Message)
		test.Equal("STOP_NOT_FOUND", responseError.Message)
	}
	tearDownDatabase()
}

func TestDeleteShuttleStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, _, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request delete shuttle route
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/stop/%s", stopList[9].StopName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(204, response.StatusCode)
	// Request delete shuttle route with invalid tag
	req = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/stop/%s", "invalid_tag"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("STOP_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestGetShuttleRouteStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle route
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/route/%s/stop", routeList[0].RouteName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleRouteStopListResponse responses.ShuttleRouteStopListResponse
	_ = json.NewDecoder(response.Body).Decode(&shuttleRouteStopListResponse)
	test.NotEmpty(shuttleRouteStopListResponse.Data)
	for _, shuttleRouteStop := range shuttleRouteStopListResponse.Data {
		test.NotEmpty(shuttleRouteStop.Name)
		test.GreaterOrEqual(shuttleRouteStop.Seq, 0)
		test.NotEmpty(shuttleRouteStop.CumulativeTime)
	}
	// Request shuttle route with invalid tag
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/shuttle/route/%s/stop", "invalid_tag"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	tearDownDatabase()
}

func TestCreateShuttleRouteStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		StopName       string `json:"name"`
		StopOrder      int    `json:"seq"`
		CumulativeTime string `json:"cumulativeTime"`
	}{
		{
			StopName:       stopList[0].StopName,
			StopOrder:      0,
			CumulativeTime: "00:00:00",
		},
		{
			StopName:       stopList[9].StopName,
			StopOrder:      10,
			CumulativeTime: "00:05:00",
		},
	}
	expectedStatusCodes := []int{409, 201}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/shuttle/route/%s/stop", routeList[0].RouteName), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		if response.StatusCode == 201 {
			var responseShuttleRouteStop responses.ShuttleRouteStopItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttleRouteStop)
			test.NotEmpty(responseShuttleRouteStop.Name)
			test.NotEmpty(responseShuttleRouteStop.Seq)
			test.NotEmpty(responseShuttleRouteStop.CumulativeTime)
		} else if response.StatusCode == 409 {
			var responseError responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&responseError)
			test.NotEmpty(responseError.Message)
			test.Equal("STOP_ALREADY_EXISTS", responseError.Message)
		}
	}
	body, _ := json.Marshal(testCases[0])
	req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/shuttle/route/%s/stop", "invalid_tag"), bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("ROUTE_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestUpdateShuttleRouteStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle route
	testCases := []struct {
		StopOrder      int    `json:"seq"`
		CumulativeTime string `json:"cumulativeTime"`
	}{
		{
			StopOrder:      0,
			CumulativeTime: "00:13:00",
		},
	}
	expectedStatusCodes := []int{200}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/api/v1/shuttle/route/%s/stop/%s", routeList[0].RouteName, stopList[0].StopName), bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		var responseShuttleRouteStop responses.ShuttleRouteStopItem
		_ = json.NewDecoder(response.Body).Decode(&responseShuttleRouteStop)
		test.NotEmpty(responseShuttleRouteStop.Name)
		test.GreaterOrEqual(responseShuttleRouteStop.Seq, 0)
		test.NotEmpty(responseShuttleRouteStop.CumulativeTime)
	}
	tearDownDatabase()
}

func TestDeleteShuttleRouteStop(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	stopList, routeList, _, _, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request delete shuttle route
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/route/%s/stop/%s", routeList[0].RouteName, stopList[0].StopName), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(204, response.StatusCode)
	// Request delete shuttle route with invalid tag
	req = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/shuttle/route/%s/stop/%s", routeList[0].RouteName, "invalid_tag"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("STOP_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestGetShuttlePeriodList(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request all shuttle periods
	req := httptest.NewRequest("GET", "/api/v1/shuttle/period", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttlePeriodListResponse responses.ShuttlePeriodListResponse
	_ = json.NewDecoder(response.Body).Decode(&shuttlePeriodListResponse)
	test.NotEmpty(shuttlePeriodListResponse.Data)
	for _, shuttlePeriod := range shuttlePeriodListResponse.Data {
		test.NotEmpty(shuttlePeriod.PeriodType)
		test.NotEmpty(shuttlePeriod.StartDate)
		test.NotEmpty(shuttlePeriod.EndDate)
	}
	tearDownDatabase()
}

func TestGetShuttlePeriodItem(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, periodList, _ := shuttleDataCreate()
	fmt.Println(periodList[0])
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle period
	req := httptest.NewRequest("GET", fmt.Sprintf(
		"/api/v1/shuttle/period/%s/%s",
		periodList[0].PeriodStart.AddDate(0, 0, 1).Format("2006-01-02"),
		periodList[0].PeriodEnd.Format("2006-01-02"),
	), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)
	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttlePeriodItemResponse responses.ShuttlePeriodItem
	_ = json.NewDecoder(response.Body).Decode(&shuttlePeriodItemResponse)
	test.NotEmpty(shuttlePeriodItemResponse.PeriodType)
	test.NotEmpty(shuttlePeriodItemResponse.StartDate)
	test.NotEmpty(shuttlePeriodItemResponse.EndDate)

	// Request shuttle period with invalid id
	req = httptest.NewRequest("GET", fmt.Sprintf(
		"/api/v1/shuttle/period/%s/%s",
		"0000-00-00",
		"0000-00-00",
	), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)
	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("PERIOD_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestCreateShuttlePeriod(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, periodList, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle period
	testCases := []struct {
		PeriodType  string `json:"type"`
		PeriodStart string `json:"start"`
		PeriodEnd   string `json:"end"`
	}{
		{
			PeriodType:  periodList[0].PeriodType,
			PeriodStart: periodList[0].PeriodStart.AddDate(0, 0, 1).Format("2006-01-02"),
			PeriodEnd:   periodList[0].PeriodEnd.Format("2006-01-02"),
		},
		{
			PeriodType:  shuttlePeriodType[randomdata.Number(0, 2)],
			PeriodStart: "2020-01-01",
			PeriodEnd:   "2020-01-03",
		},
	}
	expectedStatusCodes := []int{409, 201}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/shuttle/period", bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)
		test.Equal(expectedStatusCodes[index], response.StatusCode)

		if response.StatusCode == 201 {
			var responseShuttlePeriod responses.ShuttlePeriodItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttlePeriod)
			test.NotEmpty(responseShuttlePeriod.PeriodType)
			test.NotEmpty(responseShuttlePeriod.StartDate)
			test.NotEmpty(responseShuttlePeriod.EndDate)
		} else if response.StatusCode == 409 {
			var responseError responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&responseError)
			test.NotEmpty(responseError.Message)
			test.Equal("PERIOD_ALREADY_EXISTS", responseError.Message)
		}
	}
	tearDownDatabase()
}

func TestDeleteShuttlePeriod(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, periodList, _ := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request delete shuttle period
	req := httptest.NewRequest("DELETE", fmt.Sprintf(
		"/api/v1/shuttle/period/%s/%s",
		periodList[0].PeriodStart.AddDate(0, 0, 1).Format("2006-01-02"),
		periodList[0].PeriodEnd.Format("2006-01-02"),
	), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)

	test.Equal(204, response.StatusCode)
	// Request delete shuttle period with invalid id
	req = httptest.NewRequest("DELETE", fmt.Sprintf(
		"/api/v1/shuttle/period/%s/%s",
		"0000-00-00",
		"0000-00-00",
	), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)

	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("PERIOD_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestGetShuttleHolidayList(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request all shuttle holidays
	req := httptest.NewRequest("GET", "/api/v1/shuttle/holiday", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)

	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleHolidayListResponse responses.ShuttleHolidayListResponse
	_ = json.NewDecoder(response.Body).Decode(&shuttleHolidayListResponse)
	test.NotEmpty(shuttleHolidayListResponse.Data)
	for _, shuttleHoliday := range shuttleHolidayListResponse.Data {
		test.NotEmpty(shuttleHoliday.HolidayType)
		test.NotEmpty(shuttleHoliday.HolidayDate)
	}
	tearDownDatabase()
}

func TestGetShuttleHolidayItem(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, _, holidayList := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle holiday
	req := httptest.NewRequest("GET", fmt.Sprintf(
		"/api/v1/shuttle/holiday/%s/%s",
		holidayList[0].CalendarType,
		holidayList[0].HolidayDate.Format("2006-01-02")), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)

	test.Equal(200, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var shuttleHolidayItemResponse responses.ShuttleHolidayItem
	_ = json.NewDecoder(response.Body).Decode(&shuttleHolidayItemResponse)
	test.NotEmpty(shuttleHolidayItemResponse.HolidayType)
	test.NotEmpty(shuttleHolidayItemResponse.HolidayDate)
	// Request shuttle holiday with invalid id
	req = httptest.NewRequest("GET", fmt.Sprintf(
		"/api/v1/shuttle/holiday/%s/%s",
		"solar",
		"2000-01-01",
	), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)

	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("HOLIDAY_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}

func TestCreateShuttleHoliday(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, _, holidayList := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request create shuttle holiday
	testCases := []struct {
		CalendarType string `json:"calendar"`
		HolidayType  string `json:"holiday"`
		HolidayDate  string `json:"date"`
	}{
		{
			CalendarType: holidayList[0].CalendarType,
			HolidayType:  holidayList[0].HolidayType,
			HolidayDate:  holidayList[0].HolidayDate.Format("2006-01-02"),
		},
		{
			CalendarType: "solar",
			HolidayType:  shuttleHolidayType[randomdata.Number(0, 2)],
			HolidayDate:  "2020-01-01",
		},
	}
	expectedStatusCodes := []int{409, 201}
	for index, testCase := range testCases {
		body, _ := json.Marshal(testCase)
		req := httptest.NewRequest("POST", "/api/v1/shuttle/holiday", bytes.NewBuffer(body))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Content-Type", "application/json")
		response, _ := app.Test(req, 5000)

		test.Equal(expectedStatusCodes[index], response.StatusCode)

		if response.StatusCode == 201 {
			var responseShuttleHoliday responses.ShuttleHolidayItem
			_ = json.NewDecoder(response.Body).Decode(&responseShuttleHoliday)
			test.NotEmpty(responseShuttleHoliday.HolidayType)
			test.NotEmpty(responseShuttleHoliday.HolidayDate)
		} else if response.StatusCode == 409 {
			var responseError responses.ErrorResponse
			_ = json.NewDecoder(response.Body).Decode(&responseError)
			test.NotEmpty(responseError.Message)
			test.Equal("HOLIDAY_ALREADY_EXISTS", responseError.Message)
		}
	}
	tearDownDatabase()
}

func TestDeleteShuttleHoliday(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, _, _, _, holidayList := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request delete shuttle holiday
	req := httptest.NewRequest("DELETE", fmt.Sprintf(
		"/api/v1/shuttle/holiday/%s/%s",
		holidayList[0].CalendarType,
		holidayList[0].HolidayDate.Format("2006-01-02")), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(req, 5000)

	test.Equal(204, response.StatusCode)
	// Request delete shuttle holiday with invalid id
	req = httptest.NewRequest("DELETE", fmt.Sprintf(
		"/api/v1/shuttle/holiday/%s/%s",
		"solar",
		"2000-01-01"), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")
	response, _ = app.Test(req, 5000)

	test.Equal(404, response.StatusCode)
	test.Equal("application/json", response.Header.Get("Content-Type"))
	var responseError responses.ErrorResponse
	_ = json.NewDecoder(response.Body).Decode(&responseError)
	test.NotEmpty(responseError.Message)
	test.Equal("HOLIDAY_NOT_FOUND", responseError.Message)
	tearDownDatabase()
}
