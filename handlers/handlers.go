package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mcnijman/go-emailaddress"
	"github.com/shuttlersIT/intel/database"
	"github.com/shuttlersIT/intel/structs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var c string
var conf *oauth2.Config

// RandToken generates a random @l length token.
func RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {
	cid := "946670882701-dcidm9tcfdpcikpbjj8rfsb6uci22o4s.apps.googleusercontent.com"
	cs := "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"

	conf = &oauth2.Config{
		ClientID:     cid,
		ClientSecret: cs,
		RedirectURL:  "https://intelligence.shuttlers.africa/rest/oauth2-credential/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// IndexHandler handles the location /.
func IndexHandler(c *gin.Context) {
	state, err := RandToken(32)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "Error while generating random data."})
		return
	}
	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "login.html", gin.H{"link": link})
}

// IndexHandler handles the login
func IndexHandler2(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func AuthHandler(c *gin.Context) {
	//Declare shuttlers domain
	shuttlersDomain := "shuttlers.ng"

	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
		return
	}
	code := c.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := structs.User{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}

	usersEmail, err := emailaddress.Parse(u.Email)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Looks like your email address is invalid, try again."})
	}

	if strings.Compare(usersEmail.Domain, shuttlersDomain) == 0 {
		session.Set("user-id", u.Email)
		session.Set("user-name", u.Name)
		err = session.Save()
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusBadRequest, "index.html", gin.H{"message": "Error while saving session. Please try again."})
			return
		}
		seen := false
		db := database.MongoDBConnection{}
		if _, mongoErr := db.LoadUser(u.Email); mongoErr == nil {
			seen = true
		} else {
			err = db.SaveUser(&u)
			if err != nil {
				log.Println(err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Something went wrong... it's not you, its us. Please try again."})
				return
			}
		}
		c.HTML(http.StatusOK, "home.html", gin.H{"email": u.Email, "Username": u.Name, "seen": seen})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Looks like do not have a shuttlers email address, Please signin with your shuttlers email account."})

	}

}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	state, err := RandToken(32)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "Error while generating random data."})
		return
	}
	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "login.html", gin.H{"link": link})
}

// CxHandler is a rudementary handler for logged in users.
func PortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "portal.html", gin.H{"Username": userID})
}

// CxHandler is a rudementary handler for logged in users.
func CxHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "cx.html", gin.H{"Username": userID})
}

// SalesHandler is a rudementary handler for logged in users.
func SalesHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "sales.html", gin.H{"Username": userID})
}

// MarketingHandler is a rudementary handler for logged in users.
func MarketingHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "marketing.html", gin.H{"Username": userID})
}

// PeopleHandler is a rudementary handler for logged in users.
func PeopleHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "peopleandculture.html", gin.H{"Username": userID})
}

// PerformanceHandler is a rudementary handler for logged in users.
func PerformanceHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "home.html", gin.H{"Username": userID})
}

// RequestHandler is a rudementary handler for logged in users.
func RequestHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "datarequest.html", gin.H{"Username": userID})
}

// DriverHandler is a rudementary handler for logged in users.
func DriverHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "driverscorecard.html", gin.H{"Username": userID})
}

// MarshalHandler is a rudementary handler for logged in users.
func MarshalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "marshaldashboard.html", gin.H{"Username": userID})
}

// SeatHandler is a rudementary handler for logged in users.
func SeatHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "seatoccupancy.html", gin.H{"Username": userID})
}

// QaHandler is a rudementary handler for logged in users.
func QaHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "shuttlersqa.html", gin.H{"Username": userID})
}

// FeedbackHandler is a rudementary handler for logged in users.
func FeedbackHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "feedbacktracker.html", gin.H{"Username": userID})
}

// OperationsHandler is a rudementary handler for logged in users.
func OperationsHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "operations.html", gin.H{"Username": userID})
}
