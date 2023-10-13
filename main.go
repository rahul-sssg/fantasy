package main

import (
	"fantasy/base"
	"fantasy/routes"
	"fmt"
	"math/rand"
	"os"
	"time"

	// usq_google "self/Downloads/sssg/fantasy/model/google"

	"github.com/gin-gonic/gin"
)

// main function
func main() {
	base.Init()
	if os.Getenv("DEBUG") == "false" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	name := GenerateRandomString(10)
	fmt.Println("rrr ", name)
	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie(name)
		if err != nil {
			cookie = "NotSet"
			c.SetCookie(name, "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	routes.InitRoutes(router)
	router.Run(fmt.Sprintf(":%v", "8080"))
}

func GenerateRandomString(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	timestamp := fmt.Sprintf("%d", (time.Now().Unix())%1000)
	b := make([]byte, length-len(timestamp))
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b) + timestamp
}

// example
// Cookie:
// G_ENABLED_IDPS=google; gin_cookie=test

//actual
// Cookie:
// _ga=GA1.2.473074045.1695280679; _gid=GA1.2.1547942679.1697005250; _ga_Y31MW8GGCK=GS1.2.1697005250.2.0.1697005250.0.0.0
