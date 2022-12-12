package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"webgoapp.go/db"
	"webgoapp.go/service"
)

const port = 8000

func main() {
	// initialize DB connection
	dsn := db.DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	if err := db.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	// prepare session
	store := cookie.NewStore([]byte("my-secret"))
	engine.Use(sessions.Sessions("user-session", store))

	// routing
	engine.Static("/public", "./public")
	engine.GET("/", service.Home)
	//engine.GET("/list", service.LoginCheck, service.TaskList)
	engine.GET("/logout", service.Logout)
	engine.GET("/openchat", service.OpenchatForm)
	engine.POST("/openchat", service.NotImplemented)
	engine.POST("/privatechat", service.NotImplemented)

	//use group for middleware
	/*taskGroup := engine.Group("/thread")
	  taskGroup.Use(service.LoginCheck)
	  {
	      taskGroup.GET("/:id", service.ShowThread)
	      taskGroup.POST("/:id", service.PostMessage)
	  }*/

	//register user
	engine.GET("/user/new", service.NewUserForm)
	engine.POST("/user/new", service.RegisterUser)

	//login user
	engine.GET("/login", service.UserLoginForm)
	engine.POST("/login", service.Login)

	//delete user
	/*engine.GET("/withdraw", service.LoginCheck, service.Withdraw)

	//update user information
	engine.GET("/user/edit", service.LoginCheck, service.EditUserForm)
	engine.POST("/user/edit", service.UpdateUser)*/

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}
