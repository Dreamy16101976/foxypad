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

const filename string = "foxypad.txt" //data filename
const port string = ":8888"           //port number
const rows string = "20"              //rows number
const login string = "user"           //login
const password string = "password"    //password

//error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		login: password,
	}))

	//404 handler
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	//main page handler
	authorized.GET("/", func(c *gin.Context) {
		f, err := os.Open(filename) //data file open
		check(err)
		_, err = f.Seek(0, 0) //go to begin of file
		check(err)
		b, err := ioutil.ReadAll(f) //data read from opened file
		check(err)
		f.Close()         //data file close
		text := string(b) //conversion binary data to string
		//page display
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"text": text,
				"rows": rows,
			},
		)
	})

	//data send handler
	authorized.POST("/send", func(c *gin.Context) {
		text := c.PostForm("text")    //get data from POST query
		f, err := os.Create(filename) //data file create
		check(err)
		defer f.Close()
		_, err = f.WriteString(text) //write data to file
		f.Sync()                     //file sync
		c.Redirect(302, "/")         //go to main page
	})

	r.Run(port)
}
