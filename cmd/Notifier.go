package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const DISCORD_API_URL = "https://discord.com/api/v10"

func main() {
	flPort := flag.Int("p", 12177, "port to listen on")

	flag.Parse()

	DISCORD_BOT_TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_BOT_CHANNEL := os.Getenv("DISCORD_BOT_CHANNEL")
	port := *flPort

	switch "" {
	case DISCORD_BOT_TOKEN:
		fmt.Println("DISCORD_BOT_TOKEN envvar required")
		return
	case DISCORD_BOT_CHANNEL:
		fmt.Println("DISCORD_BOT_CHANNEL envvar required")
		return
	}

	MESSAGE_SEND_URL := fmt.Sprintf("%s/channels/%s/messages", DISCORD_API_URL, DISCORD_BOT_CHANNEL)

	sendMessage := func(msg string) {
		body, _ := json.Marshal(map[string]string{
			"content": msg,
		})
		req, _ := http.NewRequest("POST", MESSAGE_SEND_URL, bytes.NewBuffer(body))
		req.Header.Add("Authorization", fmt.Sprintf("Bot %s", DISCORD_BOT_TOKEN))
		req.Header.Add("Content-Type", "application/json")

		resp, _ := http.DefaultClient.Do(req)

		if resp.StatusCode != 200 {
			fmt.Println(resp)
		}
	}

	r := gin.Default()

	r.GET("/notify", func(c *gin.Context) {
		var msg string

		if msg = c.Query("msg"); msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "msg parameter missing"})
			return
		}

		sendMessage(msg)

		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	sendMessage("init")
	r.Run(fmt.Sprintf("127.0.0.1:%d", port))
}
