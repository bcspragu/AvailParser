# AvailParser
A Go-based query parser for InfoPoint installations

## Example Snippets

```Go
func init() {
	feed = avail.NewFeed("http://bustracker.pvta.com")
	avail.SetLocation("America/New_York")
}

func responseFromDeparture(stopDep avail.StopDeparture) StopResponse {
	res := StopResponse{
		StopName: stops[stopDep.StopId].Name,
		Routes:   []RouteInfo{},
	}
	uniqueISDs := make(map[string]bool)

	for _, dir := range stopDep.RouteDirections {
		route := routes[dir.RouteId]
		for _, dep := range dir.Departures {
			// If we haven't seen this one before
			if _, ok := uniqueISDs[dep.Trip.InternetServiceDesc]; !ok {
				if dep.EDT.After(time.Now()) {
					uniqueISDs[dep.Trip.InternetServiceDesc] = true
					routeInfo := RouteInfo{
						Number:        route.ShortName,
						Name:          dep.Trip.InternetServiceDesc,
						Color:         route.Color,
						TextColor:     route.TextColor,
						DepartureTime: dep.EDT,
					}
					res.Routes = append(res.Routes, routeInfo)
				}
			}
		}
	}
	return res
}
```
