package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// "GET" の処理

func server(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_name") == nil {
		panic("ログインしてない")
	}
	server := serverGetAll()
	sort.Slice(server, func(i, j int) bool {
		return server[i].ID < server[j].ID
	})
	c.HTML(200, "server.html", gin.H{"user_name": session.Get("user_name"), "server": server})
}

func root_Server(c *gin.Context) {
	server := serverGetAll()
	c.HTML(http.StatusOK, "rootServer.html", gin.H{"server": server})
}

func root_ServerDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	server := serverGetOne(id)
	c.HTML(200, "rootServerDetail.html", gin.H{"server": server})
}

func root_ServerDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	server := serverGetOne(id)
	c.HTML(200, "rootServerDelete.html", gin.H{"server": server})
}

// "POST" の処理

func server_Check(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("user_name")
	a := c.PostForm("anser")
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	server, anser := check_linux(id, a)
	c.HTML(http.StatusOK, "serverCheck.html", gin.H{"user_name": name, "server": server, "anser": anser, "a": a})
}

func root_ServerNew(c *gin.Context) {
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	serverInsert(question, anser, hint)
	c.Redirect(302, "/root/server")
}

func root_ServerDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	linuxDelete(id)
	c.Redirect(302, "/root/linux")
}

func root_ServerUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	serverUpdate(id, question, hint, anser)
	c.Redirect(302, "/root/server")
}
