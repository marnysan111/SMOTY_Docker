package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// "GET" の処理

func top(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func smoty(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_name") == nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("ログインしてない"))
		return
	}
	c.HTML(200, "smoty.html", gin.H{"user_name": session.Get("user_name")})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_name") == nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("ログインしてない"))
		return
	}
	session.Clear()
	session.Save()
	c.HTML(200, "logout.html", gin.H{})
}

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "root.html", gin.H{})
}

// "POST" の処理

func signup(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	dbSignup(name, password)
	c.HTML(http.StatusOK, "signup.html", gin.H{"name": name, "password": password})
}

func login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	session := sessions.Default(c)
	dblogin(name, password)
	session.Set("user_name", name)
	session.Save()
	c.Redirect(302, "/smoty")
}
