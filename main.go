
package main

import (
	"fmt"
	"log"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shuttlersIT/intel/handlers"
	"github.com/shuttlersIT/intel/middleware"
)

func main() {
	//RedisHost := "127.0.0.1"
	//RedisPort := "6379"

	//Dashboard APP
	router := gin.Default()
	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		fmt.Println(storeError)
		log.Fatal("Unable to create work with redis session ", storeError)
	}
	//store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("portalsession", store))

	// We check for errors.

	//sessionStore, _ := sessions.NewRedisStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte(token))
	//router.Use(sessions.Sessions("portalsession", sessionStore))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/css", "templates/css")
	router.Static("/img", "templates/img")
	router.Static("/js", "templates/js")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handlers.IndexHandler)
	router.GET("/index", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/logout", handlers.LogoutHandler)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/portal", handlers.PortalHandler)
		authorized.GET("/cx", handlers.CxHandler)
		authorized.GET("/sales", handlers.SalesHandler)
		authorized.GET("/home", handlers.PerformanceHandler)
		authorized.GET("/marketing", handlers.MarketingHandler)
		authorized.GET("/operations", handlers.OperationsHandler)
		authorized.GET("/driverscorecard", handlers.DriverHandler)
		authorized.GET("/feedbacktracker", handlers.FeedbackHandler)
		authorized.GET("/marshaldashboard", handlers.MarshalHandler)
		authorized.GET("/peopleandculture", handlers.PeopleHandler)
		authorized.GET("/seatoccupancy", handlers.SeatHandler)
		authorized.GET("/shuttlersqa", handlers.QaHandler)
		authorized.GET("/datarequest", handlers.RequestHandler)
		authorized.GET("/testing", homeTest)
	}
	//router.Use(static.Serve("/", static.LocalFile("./templates", true)))

	if err := router.Run(":9193"); err != nil {
		log.Fatal(err)
	}

}

func homeTest(c *gin.Context) {
	s := sessions.Default(c)
	var count int
	v := s.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	s.Set("count", count)
	s.Save()

	c.JSON(200, gin.H{"count": count})
}
