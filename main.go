package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"time"
)

// Entry is mongo structure
type Entry struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Time time.Time
}

func main() {
	lab := os.Getenv("MONGOLAB_URI")

	session, err := mgo.Dial(lab)
	col := session.DB("go-test").C("names")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/public", "public")

	r.GET("/", func(c *gin.Context) {
		var results []Entry
		col.Find(nil).All(&results)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": results,
		})
	})

	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		err = col.Insert(&Entry{Name: name, Time: time.Now()})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"this": "works"})
	})

	r.GET("/cleardb", func(c *gin.Context) {
		// col.RemoveAll(nil)
		col.DropCollection()
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":" + port)
}
