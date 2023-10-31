package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
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
		RedirectURL:  "https://intel.shuttlers.africa/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

/* IndexHandler handles the location /.
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
}*/

// IndexHandler handles the login
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func AuthHandler(c *gin.Context) {
	//Declare shuttlers domain
	//shuttlersDomain := "shuttlers.ng"

	// Handle the exchange code to initiate a transport.
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"message": "Invalid session state."})
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

	session.Set("user-id", u.Email)
	session.Set("user-name", u.Name)
	err = session.Save()

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	//seen := false

	/*
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
	*/
	userID := session.Get("user-id")
	fmt.Println(userID)
	fmt.Println(session)
	//uName := session.Get("user-name")
	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userID})
	//c.HTML(http.StatusOK, "home.html", gin.H{"name": uNam, "Username": userID, "seen": seen})

	/*
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
						c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
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
				c.HTML(http.StatusOK, "portal.html", gin.H{"email": u.Email, "Username": u.Name, "seen": seen})
		} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Looks like do not have a shuttlers email address, Please signin with your shuttlers email account."})

		}
	*/
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	state, err := RandToken(32)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while generating random data."})
		return
	}
	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.html", gin.H{"link": link})
}

// Logout Handler
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Signed out successfully",
	})
}

// RequestHandler is a rudementary handler for logged in users.
func RequestHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "datarequest.html", gin.H{"Username": userID})
}

// ITSM Home
func ItsmHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userID})
}

// ITSM Desk
func ItDeskPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userID})
}
func ItDeskAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "itdeskadmin.html", gin.H{"Username": userID})
}
func ItDeskHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userID})
}

// Assets
func AssetsPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "assetsportal.html", gin.H{"Username": userID})
}

func AssetsAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "assetsadmin.html", gin.H{"Username": userID})
}
func AssetsHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "assetsx.html", gin.H{"Username": userID})
}

// Procurement
func ProcurementPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementportal.html", gin.H{"Username": userID})
}

func ProcurementAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementadmin.html", gin.H{"Username": userID})
}
func ProcurementHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user-id")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userID})
}
