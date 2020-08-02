package main

import (
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
		panic("ログインしてない")
	}
	linux := linuxGetAll()
	sort.Slice(linux, func(i, j int) bool {
		return linux[i].ID < linux[j].ID
	})
	c.HTML(200, "linux.html", gin.H{"user_name": session.Get("user_name"), "linux": linux})
}

func root_Linux(c *gin.Context) {
	linux := linuxGetAll()
	c.HTML(http.StatusOK, "rootLinux.html", gin.H{"linux": linux})
}

func root_LinuxDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	linux := linuxGetOne(id)
	c.HTML(200, "rootLinuxDetail.html", gin.H{"linux": linux})
}

func root_LinuxDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	linux := linuxGetOne(id)
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
		panic(err)
	}
	linux, anser := check_linux(id, a)
	c.HTML(http.StatusOK, "linuxCheck.html", gin.H{"user_name": name, "linux": linux, "anser": anser, "a": a})
}

func root_LinuxNew(c *gin.Context) {
	question := c.PostForm("question")
	anser := c.PostForm("anser")
	hint := c.PostForm("hint")
	linuxInsert(question, anser, hint)
	c.Redirect(302, "/root/linux")
}

func root_LinuxDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	linuxDelete(id)
	c.Redirect(302, "/root/linux")
}

func root_LinuxUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	question := c.PostForm("question")
	hint := c.PostForm("hint")
	anser := c.PostForm("anser")
	linuxUpdate(id, question, hint, anser)
	c.Redirect(302, "/root/linux")
}
