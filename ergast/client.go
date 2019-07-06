package ergast

import "net/http"

// BaseURL is the base url for the ergast api
const BaseURL = "https://ergast.com/api/f1"

// Client is be the http client used for ergast http requests
var Client http.Client
