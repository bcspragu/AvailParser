# AvailParser
A Go-based query parser for InfoPoint installations

## Example Usage

```Golang
package main

import (
	"avail"
	"html/template"
	"net/http"
)

var feed *avail.Feed

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	feed = avail.NewFeed("http://bustracker.pvta.com")

	http.HandleFunc("/", messageHandler)
	http.ListenAndServe(":9000", nil)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("Body")

	var message string
	if len(query) < 5 {
		message = "Please enter at least 5 letters for your search"
	} else {
		message = avail.HumanReadableMessage(feed.NextDeparturesByStopName(query))
		if message == "" {
			message = "No results found, please check your spelling and try again."
		}
	}

	messageData := struct {
		Message string
	}{message}

	w.Header().Add("Content-Type", "text/xml")

	templates.ExecuteTemplate(w, "index.html", messageData)
}
```
