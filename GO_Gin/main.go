package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func getApi(c *gin.Context) {
	//r.GET("/user/:name/:action", getApi)
	name := c.Param("name")
	action := c.Param("action")
	c.String(200, fmt.Sprintf(name+"is %s", action))

}

func getUrl(c *gin.Context) {
	name := c.Query("name")
	action := c.Query("action")
	c.String(200, fmt.Sprintf("name=%s, action=%s", name, action))
}

func postApi(c *gin.Context) {
	name := c.PostForm("name")
	action := c.PostForm("action")
	c.String(200, fmt.Sprintf("name=%s, action=%s", name, action))
}

func uploadFile(c *gin.Context) {
	//r.MaxMultipartMemory = 8 << 20
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, file.Filename)
	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func uploadfiles(c *gin.Context) {
	//r.MaxMultipartMemory = 8 << 20
	form, err := c.MultipartForm()
	if err != nil {
		c.String(400, "err")
	}
	files := form.File["file"]

	for _, file := range files {
		c.SaveUploadedFile(file, file.Filename)
	}
	c.String(200, fmt.Sprintf("%d files uploaded!", len(files)))
}

type Person struct {
	Name string `form:"name" json:"name" uri:"name" xml:"name" `
}

func bindings(c *gin.Context) {
	var p Person
	c.ShouldBindJSON(&p)
	c.Bind(&p)
	c.ShouldBindUri(&p)
}

func cookies() {
	r := gin.New()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("MyCookie", "114", 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"msg": "login yes",
		})
	})

	r.GET("/home", AuthCheck(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "home",
		})
	})
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("MyCookie")
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "err get cookie",
			})
			c.Abort()
			return
		}
		if cookie != "114" {
			c.JSON(200, gin.H{
				"msg": "wrong cookie",
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"msg": "login yes"})
		c.Next()
	}
}

func test_sessions() {
	var secretKey = []byte("GenshinImpact")
	var store = cookie.NewStore(secretKey)

	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/get", func(c *gin.Context) {
		session0 := sessions.Default(c)
		if session0.Get("甘雨") != "椰羊" {
			session0.Set("甘雨", "椰羊")
			session0.Delete("大司马")
			session0.Save()
		}
		c.JSON(200, gin.H{
			"甘雨":  session0.Get("甘雨"),
			"大司马": session0.Get("大司马"),
		})
	})
	r.Run()
}

var store = cookie.NewStore([]byte("ILoveYuanshen"))

func test_sessions_proto() {
	http.HandleFunc("/get", GetSession)
	http.HandleFunc("/save", SaveSession)
	http.ListenAndServe(":8080", nil)
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	session1, _ := store.Get(r, "session-name")
	session1.Values["原神"] = "启动"
	session1.Save(r, w)

}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session1, _ := store.Get(r, "session-name")
	if session1.Values["原神"] != "启动" {
		session1.Values["原神"] = "启动"
		session1.Save(r, w)
	}
	fmt.Fprintln(w, "原神", session1.Values["原神"])
}

func gin_log() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongpongponghahahaaaaaasdasdagsdhgsdfhdfjdfgjdf",
		})
	})
	r.GET("/this_is_first_site", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")

	})
	r.Run()
}

// func main() {
// 	gin_log()
// }
