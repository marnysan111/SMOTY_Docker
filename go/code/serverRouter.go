package main

import (
	"errors"
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
		c.AbortWithError(http.StatusUnauthorized, errors.New("ログインしてない"))
		return
	}
	server, err := serverGetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sort.Slice(server, func(i, j int) bool {
		return server[i].ID < server[j].ID
	})
	c.HTML(200, "server.html", gin.H{"user_name": session.Get("user_name"), "server": server})
}

func root_Server(c *gin.Context) {
	server, err := serverGetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "rootServer.html", gin.H{"server": server})
}

func root_ServerDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	server, err := serverGetOne(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(200, "rootServerDetail.html", gin.H{"server": server})
}

func root_ServerDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	server, nil := serverGetOne(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
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
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	server, anser, err := check_server(id, a)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "serverCheck.html", gin.H{"user_name": name, "server": server, "anser": anser, "a": a})
}

func root_ServerNew(c *gin.Context) {
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	err := serverInsert(question, anser, hint)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(302, "/root/server")
}

func root_ServerDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = linuxDelete(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(302, "/root/linux")
}

func root_ServerUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	err = serverUpdate(id, question, hint, anser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(302, "/root/server")
}
