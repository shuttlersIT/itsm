package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/intel/database"
	"github.com/shuttlersIT/intel/structs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type ticket structs.Ticket
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

// Share databases
func ShareDb(d *sql.DB) *sql.DB {
	db := d
	database.TableExists(db, "tickets")
	return d
}

func getID(e string) int, string{
	var id int
	email := e
	err := db.QueryRow("SELECT id FROM staff WHERE email = ?", email).
		Scan(&id)
	if err != nil {
		session := sessions.Default(c)
		email := session.Get("user-email")
		t := handler.CreateUser()
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff email not found in our records"})
		return
	}
	return id, email
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

	// Get user ID
	session.Set("user-email", u.Email)
	session.Set("user-name", u.Name)
	session.Set("user-firstName", u.GivenName)
	session.Set("user-lastName", u.FamilyName)
	session.Set("user-sub", u.Sub)
	err = session.Save()

	//
	userId := getID(u.Email)
	session.Set("id", userId)
	session.Set("user-email", u.Email)
	session.Set("user-name", u.Name)
	session.Set("user-firstName", u.GivenName)
	session.Set("user-lastName", u.FamilyName)
	session.Set("user-sub", u.Sub)
	err = session.Save()

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	//seen := false

	userEmail := session.Get("user-Email")
	fmt.Println(userEmail)
	fmt.Println(session)
	//uName := session.Get("user-name")

	//COMPARE GOOGLE CREDENTIALS TO DB
	itsmUser := handlers.GetStaff()

	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userEmail})
	//c.HTML(http.StatusOK, "home.html", gin.H{"name": uNam, "Username": userID, "seen": seen})

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
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "datarequest.html", gin.H{"Username": userEmail})
}

// ITSM Home
func ItsmHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userEmail})
}

// ITSM Desk
func ItDeskPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userEmail})
}
func ItDeskAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "itdeskadmin.html", gin.H{"Username": userEmail})
}
func ItDeskHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userEmail})
}

// Assets
func AssetsPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "assetsportal.html", gin.H{"Username": userEmail})
}

func AssetsAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "assetsadmin.html", gin.H{"Username": userEmail})
}
func AssetsHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "assetsx.html", gin.H{"Username": userEmail})
}

// Procurement
func ProcurementPortalHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "procurementportal.html", gin.H{"Username": userEmail})
}

func ProcurementAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "procurementadmin.html", gin.H{"Username": userEmail})
}
func ProcurementHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userEmail})
}
