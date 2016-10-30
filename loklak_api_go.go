// Package loklak_api_go is a library for interacting with a Loklak server
package loklak_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/hokaccha/go-prettyjson"
)

// The Loklak Object structure.
type Loklak struct {
	baseURL    string
	Name       string
	Followers  string
	Following  string
	Query      string
	Since      string
	Until      string
	Source     string
	Count      string
	Fields     string
	FromUser   string
	Limit      string
	ScreenName string
	Order      string
	OrderBy    string
}

// Initiation of the loklak object
func (l *Loklak) Connect(urlString string) {
	u, err := url.Parse(urlString)
	if (err != nil) {
		fmt.Println(u)
		fatal(err)
	} else {
		l.baseURL = urlString
	}
}

// A generic query URL request and fetch JSON response
// This should be suitable for a majority of the JSON based responses
// Plain text and CSV format responses need another custom control function.
// Function name: getJSON()
// Scope        : globally accessible
// Parameters   : string              , Variable => route
// Return Types : JSON Response       , Variable => string
//              : Error Response      , Variable => error
//
// Makes a request to the given URL and returns the JSON response obtained
// and error if any.

func getJSON(route string) (string, error) {
	r, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	var b interface{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return "", err
	}
	out, err := prettyjson.Marshal(b)
	return string(out), err
}


// The API function for the /api/hello.json api call.
func (l *Loklak) Hello() string {
	apiQuery := l.baseURL + "api/hello.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API function for the /api/peers.json api call.
func (l *Loklak) Peers() string {
	apiQuery := l.baseURL + "api/peers.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API function for the /api/status.json api call.
func (l *Loklak) Status() string {
	apiQuery := l.baseURL + "api/status.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for the /api/apps.json api call.
func (l *Loklak) Apps() string {
	apiQuery := l.baseURL + "api/apps.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for /api/settings.json api call.
// This is only a localhost query
func (l *Loklak) Settings() string {
	apiQuery := "http://localhost:9000/api/settings.json"
	out, err := getJSON(apiQuery)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for /api/search.json api call
// Format in order as 
// Search function is implemented as a function and not as a method
// Package the parameters required in the loklak object and pass accordingly
func Search(l *Loklak) string {
	apiQuery := l.baseURL + "api/search.json"
	req, _ := http.NewRequest("GET", apiQuery, nil)

	q := req.URL.Query()
	
	// Query constructions
	if l.Query != "" {
		constructString := l.Query
		if l.Since != "" {
			constructString += " since:" + l.Since
		}
		if l.Until != "" {
			constructString += " until:" + l.Until
		}
		if l.FromUser != "" {
			constructString += " from:" + l.FromUser
		}
		fmt.Println(constructString)
		q.Add("q", constructString)
	}
	if l.Count != "" {
		q.Add("count", l.Count)
	}
	if l.Source != "" {
		q.Add("source", l.Source)
	}
	req.URL.RawQuery = q.Encode()
	queryURL := req.URL.String()
	out, err := getJSON(queryURL)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for /api/user.json api call
func User(l *Loklak) string {
	apiQuery := l.baseURL + "api/user.json"
	req, _ := http.NewRequest("GET", apiQuery, nil)

	q := req.URL.Query()

	// Query construction
	if l.ScreenName != "" {
		q.Add("screen_name", l.ScreenName)
	}
	if l.Following != "" {
		q.Add("following", l.Following)
	}
	if l.Followers != "" {
		q.Add("followers", l.Followers)
	}
	req.URL.RawQuery = q.Encode()
	queryURL := req.URL.String()
	out, err := getJSON(queryURL)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for the /api/account.json api call
func Account(l *Loklak) string {
	apiQuery := "http://localhost:9000/api/account.json"

	req, _ := http.NewRequest("GET", apiQuery, nil)

	q := req.URL.Query()

	// Query construction
	if l.ScreenName != "" {
		q.Add("screen_name", l.ScreenName)
	}
	req.URL.RawQuery = q.Encode()
	queryURL := req.URL.String()

	out, err := getJSON(queryURL)
	if err != nil {
		fatal(err)
	}
	return out
}

// The API Function for /api/suggest.json api call
func Suggest(l *Loklak) string {
	apiQuery := l.baseURL + "api/suggest.json"

	req, _ := http.NewRequest("GET", apiQuery, nil)

	q := req.URL.Query()

	// Query construction
	if l.Query != "" {
		q.Add("q", l.Query)
	}
	if l.Count != "" {
		q.Add("count", l.Count)
	}
	if l.Source != "" {
		q.Add("source", l.Source)
	}
	if l.Order != "" {
		q.Add("order", l.Order)
	}
	if l.OrderBy != "" {
		q.Add("orderby", l.OrderBy)
	}
	if l.Since != "" {
		q.Add("since", l.Since)
	}
	if l.Until != "" {
		q.Add("until", l.Until)
	}
	req.URL.RawQuery = q.Encode()
	queryURL := req.URL.String()

	out, err := getJSON(queryURL)
	if err != nil {
		fatal(err)
	}
	return out
}


// Helper function to return the error responses to stderr
// Function name: fatal()
// Scope        : globally accessible
// Parameters   : Error               , Variable => err
// Exits the program as soon as a fatal error is obtained.

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
