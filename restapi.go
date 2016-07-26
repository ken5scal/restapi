package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"io/ioutil"

)

// usr gorm for ORMapper

type Todo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo

func Index(c *gin.Context) {
	c.String(200, "hello world")
}

func TodoIndex(c *gin.Context) {
	todos := Todos {
		Todo{Name: "Write presentation", Completed:true, Due:time.Now()},
		Todo{Name: "Host meetup", Completed:false, Due:time.Now()},
	}
	c.JSON(200, todos)
}

func TodoPost(c *gin.Context) {
	title := c.PostForm("title")
	message := c.PostForm("message")

	c.JSON(200, gin.H{
		"status": "posted",
		"title" : title,
		"message" :message,
	})
}

func FileOutput(c *gin.Context) {
	contents, err := ioutil.ReadFile("sample.json")
	if err != nil {
		//fmt.Println(contents, err)
		c.JSON(400, gin.H{"status": "bad request"})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(200, string(contents))
}

//func handle(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello world!")
//}

func init() {
	//http.HandleFunc("/", handle)
	//fmt.Printf("Start GO HTTP Server")
	//http.ListenAndServe(":8080", nil)
	r := gin.Default()
	r.GET("/", Index)
	r.GET("/todos", TodoIndex)
	r.POST("/todo", TodoPost)
	r.GET("/file", FileOutput)
	//r.Run(":8080")
	//fmt.Println("server start port 8080")
	http.Handle("/", r)
}
