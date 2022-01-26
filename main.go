package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "你好，世界",
		})
	})

	// 如果指定了端口就监听端口
	if port := os.Getenv("PORT"); port != "" {
		log.Println("正在监听端口", port)
		err := router.Run(":" + port)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// 否则监听 Unix socket
	if socket := os.Getenv("SOCKET"); socket == "" {
		log.Fatal("未指定 socket")
	} else {
		log.Println("正在监听 socket ", socket)
		err := router.RunUnix(socket)
		if err != nil {
			log.Fatal(err)
		}
	}

}
