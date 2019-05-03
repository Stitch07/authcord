// A simple script to demo OAuth2 flow using authcord. This assumes you have a valid OAuth2 application with
// the redirect URL set to localhost:3030.
package main

import (
	"fmt"
	"github.com/Soumil07/authcord"
	"log"
	"os"
)

func main() {
	clientID, clientSecret := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
	// Initialize the session with the identify, email and guild scopes.
	session := authcord.New(clientID, clientSecret, "http://localhost:3030", []string{"identify", "email", "guilds"})

	fmt.Printf("visit the url: %s\n", session.AuthURL())
	fmt.Print("enter the code provided by the redirect: ")
	var code string
	// scan the code sent in the querystring params
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	// callback fetches an OAuth2 access token and initializes an HTTP client which is used for further requests
	// before doing this, in a regular application you should be verifying the state querystring parameter with
	// session.State
	if err := session.Callback(code); err != nil {
		log.Fatal(err)
	}
	// get the logged in user
	user, err := session.User()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the logged in user is: %s (%s)\n", user.Username, user.Email)
	// get the user's guilds
	guilds, err := session.Guilds()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the logged in user is in %d guilds\n", len(guilds))
}
