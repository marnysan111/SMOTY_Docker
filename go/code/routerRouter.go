package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// "GET" の処理

func router(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_name") == nil {
		panic("ログインしてない")
	}
	router := routerGetAll()
	sort.Slice(router, func(i, j int) bool {
		return router[i].ID < router[j].ID
	})
	c.HTML(200, "router.html", gin.H{"user_name": session.Get("user_name"), "router": router})
}

func root_Router(c *gin.Context) {
	router := routerGetAll()
	c.HTML(http.StatusOK, "rootRouter.html", gin.H{"router": router})
}

func root_RouterDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	router := routerGetOne(id)
	c.HTML(200, "rootRouterDetail.html", gin.H{"router": router})
}

func root_RouterDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	router := routerGetOne(id)
	c.HTML(200, "rootRouterDelete.html", gin.H{"router": router})
}

// "POST" の処理

func router_Check(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("user_name")
	a := c.PostForm("anser")
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	router, anser := check_router(id, a)
	c.HTML(http.StatusOK, "routerCheck.html", gin.H{"user_name": name, "router": router, "anser": anser, "a": a})
}

func root_RouterNew(c *gin.Context) {
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	routerInsert(question, anser, hint)
	c.Redirect(302, "/root/router")
}

func root_RouterDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	routerDelete(id)
	c.Redirect(302, "/root/router")
}

func root_RouterUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	routerUpdate(id, question, hint, anser)
	c.Redirect(302, "/root/router")
}
