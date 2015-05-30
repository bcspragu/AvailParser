package avail

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Feed struct {
	BaseURL string
}

// Create a new feed
func NewFeed(url string) *Feed {
	a := &Feed{url + "/InfoPoint/rest/"}

	return a
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

func (a *Feed) VisibleRoutes() []Route {
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

	return container.Route
}

func (a *Feed) Stops() []Stop {
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

	return container.Stop
}

func (a *Feed) CurrentMessages() []PublicMessage {
	req, err := a.NewAvailRequest("PublicMessages/GetCurrentMessages")

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

	var container ArrayOfPublicMessage

	err = xml.Unmarshal(body, &container)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
	}

	return container.PublicMessage
}

func (a *Feed) Route(id int) Route {
	req, err := a.NewAvailRequest("routedetails/get/" + strconv.Itoa(id))

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

	var container Route

	err = xml.Unmarshal(body, &container)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
	}

	return container
}

func (a *Feed) StopDeparture(id int) StopDeparture {
	req, err := a.NewAvailRequest("stopdepartures/get/" + strconv.Itoa(id))

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

	return container.StopDeparture
}
