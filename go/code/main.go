package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const connectString = "root:smoty@tcp(172.26.0.2:3306)/score?charset=utf8&parseTime=True&loc=Local"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../views/*.html")
	r.Static("../assets", "../assets")
	r.Static("../pictures", "../pictures")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("user", store))

	dbInit_users()
	dbInit_linux()
	dbInit_server()
	dbInit_router()

	r.GET("/", top)
	r.GET("/smoty", smoty)
	r.GET("/smoty/linux", linux)
	r.GET("/smoty/server", server)
	r.GET("/smoty/router", router)
	r.GET("/logout", logout)
	r.GET("root", root)
	r.GET("/root/linux", root_Linux)
	r.GET("/root/linux/detail/;id", root_LinuxDetail)
	r.GET("/root/linux/deleteCheck/:id", root_LinuxDeleteCheck)
	r.GET("/root/server", root_Server)
	r.GET("/root/server/detail/:id", root_ServerDetail)
	r.GET("/root/server/deleteCheck/:id", root_ServerDeleteCheck)
	r.GET("/root/router", root_Router)
	r.GET("/root/router/detail/:id", root_RouterDetail)
	r.GET("/root/router/deleteCheck/:id", root_RouterDeleteCheck)

	r.POST("/signup", signup)
	r.POST("/login", login)
	r.POST("/smoty/linux/check/:id", linux_Check)
	r.POST("/root/linux/new", root_LinuxNew)
	r.POST("/root/linux/delete/:id", root_LinuxDelete)
	r.POST("/root/linux/update/:id", root_LinuxUpdate)
	r.POST("/smoty/server/check/:id", server_Check)
	r.POST("/root/server/new", root_ServerNew)
	r.POST("/root/server/delete/:id", root_ServerDelete)
	r.POST("/root/server/update/:id", root_ServerUpdate)
	r.POST("/smoty/router/check/:id", router_Check)
	r.POST("/root/router/new", root_RouterNew)
	r.POST("/root/router/delete/:id", root_RouterDelete)
	r.POST("/root/router/update/:id", root_RouterUpdate)

	r.Run(":8080")
}
