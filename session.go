// Package authcord is a tiny Discord OAuth2 library over x/oauth2
package authcord

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"

	"golang.org/x/oauth2"
)

// Session is the Discord session used to initialize the OAuth2 flow. All methods called on Session will internally
// refresh the access token if it is invalid
type Session struct {
	State      string // a random 16 bit hex encoded string, to verify callback redirects
	config     *oauth2.Config
	httpClient *http.Client
}

// New initializes a new session with the provided clientID, secret, redirect and scopes
func New(clientID, clientSecret, redirectURL string, scopes []string) *Session {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     Endpoint,
	}
	// random 16 bit hex encoded state
	b := make([]byte, 16)
	rand.Read(b)

	return &Session{
		State:  hex.EncodeToString(b),
		config: conf,
	}
}

// AuthURL returns the authentication URL the end user can visit
func (s *Session) AuthURL() string {
	return s.config.AuthCodeURL(s.State, oauth2.AccessTypeOffline)
}

// Callback fetches an access token with the provided code and initializes an HTTP Client that can be used to do
// authenticated requests
func (s *Session) Callback(code string) error {
	ctx := context.Background()
	tok, err := s.config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	s.httpClient = s.config.Client(ctx, tok)
	return nil
}

// User fetches the currently logged in user
func (s *Session) User() (u *User, err error) {
	err = s.doJSON(http.MethodGet, "/users/@me", &u)
	return
}

// Guilds fetches a slice of partial guilds the user is a member of
func (s *Session) Guilds() (g []*Guild, err error) {
	err = s.doJSON(http.MethodGet, "/users/@me/guilds", &g)
	return
}

// wrapper over Client.Do to automatically encode JSON
func (s *Session) doJSON(method, path string, respBody interface{}) error {
	if s.httpClient == nil {
		return errors.New("http client not initialized, call Callback() with an appropriate code first")
	}
	req := newRequest(method, path, nil)
	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&respBody)
}
