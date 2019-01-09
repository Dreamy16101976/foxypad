/*  foxypad - simple online scratchpad with authorization
    Copyright (C) 2018 - Alexey "FoxyLab" Voronin
    Email:    support@foxylab.com
    This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program; if not, write to the Free Software
    Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307 USA
*/
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const filename string = "foxypad.txt"
const port string = ":8888"
const rows string = "20"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user": "password",
	}))

	authorized.GET("/", func(c *gin.Context) {
		f, err := os.Open(filename)
		check(err)
		_, err = f.Seek(0, 0)
		check(err)
		b, err := ioutil.ReadAll(f)
		check(err)
		f.Close()
		text := string(b)
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"text": text,
				"rows": rows,
			},
		)
	})

	authorized.POST("/send", func(c *gin.Context) {
		text := c.PostForm("text")
		f, err := os.Create(filename)
		check(err)
		defer f.Close()
		_, err = f.WriteString(text)
		f.Sync()
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"text": text,
				"rows": rows,
			},
		)
	})

	r.Run(port)
}
