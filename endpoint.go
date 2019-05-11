package authcord

import (
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

// APIUrl is Discord's base API Url
const APIUrl = "http://discordapp.com/api/v6"

// a wrapper over http.NewRequest to append the APIUrl to the path
func newRequest(method, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, APIUrl+path, body)
	return req
}

// Endpoint is Discord's OAuth2 endpoint
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discordapp.com/api/oauth2/authorize",
	TokenURL:  "https://discordapp.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInHeader,
}
