package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/hyuabot-developers/hyuabot-backend-golang/database"
	"github.com/hyuabot-developers/hyuabot-backend-golang/dto/responses"
	"github.com/hyuabot-developers/hyuabot-backend-golang/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http/httptest"
	"testing"
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

func shuttleDataCreate() ([]*models.ShuttleStop, []*models.ShuttleRoute) {
	stopList := make([]*models.ShuttleStop, 0)
	routeList := make([]*models.ShuttleRoute, 0)
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
			"StartStop": null.StringFrom(stopList[randomdata.Number(0, 9)].StopName),
			"EndStop":   null.StringFrom(stopList[randomdata.Number(0, 9)].StopName),
		}).(*models.ShuttleRoute)
		err := route.Insert(context.Background(), database.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
		}
		routeList = append(routeList, route)
	}
	return stopList, routeList
}

func TestGetShuttleRouteList(t *testing.T) {
	test := assert.New(t)
	// Create shuttle data
	setupDatabase()
	createAdminUser()
	_, routeList := shuttleDataCreate()
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
	_, routeList := shuttleDataCreate()
	// Get access token
	app := setup()
	accessToken := loginWithAdminUser(app)
	// Request shuttle route
	fmt.Println(routeList[0].RouteName)
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
		test.NotEmpty(stop.Seq)
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
	stopList, routeList := shuttleDataCreate()
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
	stopList, routeList := shuttleDataCreate()
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
	_, routeList := shuttleDataCreate()
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
