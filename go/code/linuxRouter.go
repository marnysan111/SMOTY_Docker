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

func linux(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_name") == nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("ログインしてない"))
		return
	}
	linux, err := linuxGetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sort.Slice(linux, func(i, j int) bool {
		return linux[i].ID < linux[j].ID
	})
	c.HTML(200, "linux.html", gin.H{"user_name": session.Get("user_name"), "linux": linux})
}

func root_Linux(c *gin.Context) {
	linux, err := linuxGetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "rootLinux.html", gin.H{"linux": linux})
}

func root_LinuxDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return

	}
	linux, err := linuxGetOne(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(200, "rootLinuxDetail.html", gin.H{"linux": linux})
}

func root_LinuxDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return

	}
	linux, nil := linuxGetOne(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(200, "rootLinuxDelete.html", gin.H{"linux": linux})
}

// "POST" の処理
func linux_Check(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("user_name")
	a := c.PostForm("anser")
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	linux, anser, err := check_linux(id, a)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "linuxCheck.html", gin.H{"user_name": name, "linux": linux, "anser": anser, "a": a})
}

func root_LinuxNew(c *gin.Context) {
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	err := linuxInsert(question, anser, hint)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(302, "/root/linux")
}

func root_LinuxDelete(c *gin.Context) {
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

func root_LinuxUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	question := c.PostForm("question")
	hint := c.PostForm("hint")
	anser := c.PostForm("anser")
	err = linuxUpdate(id, question, hint, anser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(302, "/root/linux")
}
