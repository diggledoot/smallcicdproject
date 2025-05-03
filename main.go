package main

import (
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/gin-gonic/gin"
)

func say(input string) string {
	say, err := cowsay.Say(
		input,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	return say
}
func main() {
	router := gin.Default()
	router.GET("/cowsay", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cowsay": say("Howdy 🤠"),
		})
	})
	router.Run()
}
