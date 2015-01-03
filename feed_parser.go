package avail

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Feed struct {
	BaseURL string
	Routes  map[int]Route
	Stops   map[int]Stop
}

// Create a new feed
func NewFeed(url string) *Feed {
	a := &Feed{
		BaseURL: url + "/InfoPoint/rest/",
		Routes:  make(map[int]Route),
		Stops:   make(map[int]Stop),
	}

	// Eager load all route and stop information and cache it
	a.loadRoutes()
	a.loadStops()

	return a
}

func (a *Feed) NextDeparturesByStopName(name string) QueryResult {
	result := make(QueryResult)
	stops := a.stopListByName(name)

	for _, stop := range stops {
		depsByRoute := make(DeparturesByRoute)
		for _, routeDir := range a.routeDirectionsByStopId(stop.StopId) {
			depsByRoute[a.Routes[routeDir.RouteId]] = nextDepartureForRouteDirections(routeDir)
		}
		result[stop] = depsByRoute
	}

	return result
}

func (a *Feed) stopListByName(name string) []Stop {
	stops := make([]Stop, 0, 10)
	for _, stop := range a.Stops {
		if strings.Contains(stop.Name, name) {
			stops = append(stops, stop)
		}
	}
	return stops
}

// Make a new GET request to the requested API endpoint
func (a *Feed) NewAvailRequest(action string) (*http.Request, error) {
	url := a.BaseURL + action

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	return req, nil
}

func (a *Feed) routeDirectionsByStopId(stopID int) []RouteDirection {
	req, err := a.NewAvailRequest("stopdepartures/get/" + strconv.Itoa(stopID))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing body:", err)
	}

	var container ArrayOfStopDeparture

	err = xml.Unmarshal(body, &container)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
	}

	return container.StopDeparture.RouteDirections.RouteDirection
}

func nextDepartureForRouteDirections(routeDirections ...RouteDirection) []Departure {
	isds := make([]string, 0, len(routeDirections)*2)
	departures := make([]Departure, 0, len(routeDirections)*2)
	for _, routeDirection := range routeDirections {
		for _, departure := range routeDirection.Departures.Departure {
			if uniqueISD(isds, departure.Trip.InternetServiceDesc) &&
				departure.EDT.After(time.Now()) {
				isds = append(isds, departure.Trip.InternetServiceDesc)
				departures = append(departures, departure)
			}
		}
	}
	return departures
}

func uniqueISD(isds []string, testISD string) bool {
	for _, isd := range isds {
		if isd == testISD {
			return false
		}
	}
	return true
}

func (a *Feed) loadRoutes() {
	req, err := a.NewAvailRequest("routes/getvisibleroutes")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing body:", err)
	}

	var container ArrayOfRoute

	err = xml.Unmarshal(body, &container)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
	}

	for _, route := range container.Route {
		a.Routes[route.RouteId] = route
	}
}

func (a *Feed) loadStops() {
	req, err := a.NewAvailRequest("stops/getallstops")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing body:", err)
	}

	var container ArrayOfStop

	err = xml.Unmarshal(body, &container)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
	}

	for _, stop := range container.Stop {
		a.Stops[stop.StopId] = stop
	}
}
