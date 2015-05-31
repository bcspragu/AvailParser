package avail

import (
	"encoding/xml"
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

func (a *Feed) VisibleRoutes() ([]Route, error) {
	var container ArrayOfRoute

	req, err := a.NewAvailRequest("routes/getvisibleroutes")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return container.Route, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return container.Route, err
	}

	err = xml.Unmarshal(body, &container)
	if err != nil {
		return container.Route, err
	}

	return container.Route, nil
}

func (a *Feed) Stops() ([]Stop, error) {
	var container ArrayOfStop

	req, err := a.NewAvailRequest("stops/getallstops")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return container.Stop, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return container.Stop, err
	}

	err = xml.Unmarshal(body, &container)
	if err != nil {
		return container.Stop, err
	}

	return container.Stop, nil
}

func (a *Feed) CurrentMessages() ([]PublicMessage, error) {
	var container ArrayOfPublicMessage

	req, err := a.NewAvailRequest("PublicMessages/GetCurrentMessages")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return container.PublicMessage, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return container.PublicMessage, err
	}

	err = xml.Unmarshal(body, &container)
	if err != nil {
		return container.PublicMessage, err
	}

	return container.PublicMessage, nil
}

func (a *Feed) Route(id int) (Route, error) {
	var container Route

	req, err := a.NewAvailRequest("routedetails/get/" + strconv.Itoa(id))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return container, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return container, err
	}

	err = xml.Unmarshal(body, &container)
	if err != nil {
		return container, err
	}

	return container, nil
}

func (a *Feed) StopDeparture(id int) (StopDeparture, error) {
	var container ArrayOfStopDeparture

	req, err := a.NewAvailRequest("stopdepartures/get/" + strconv.Itoa(id))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return container.StopDeparture, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return container.StopDeparture, err
	}

	err = xml.Unmarshal(body, &container)
	if err != nil {
		return container.StopDeparture, err
	}

	return container.StopDeparture, nil
}
