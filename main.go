package main

import (
	"log/slog"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func say(input string) string {
	say, err := cowsay.Say(
		input,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		slog.Error("Error!", "error", err)
		panic(err)
	}
	return say
}

func callCowsay() {
	resp, err := http.Get("http://localhost:8080/cowsay")
	if err != nil {
		slog.Error("Error!", "error", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	router := gin.Default()
	router.GET("/cowsay", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cowsay": say("Howdy 🤠"),
		})
	})

	cowsayCron := cron.New()
	cowsayCron.AddFunc("@every 3s", callCowsay)
	cowsayCron.Start()

	router.Run()
}
