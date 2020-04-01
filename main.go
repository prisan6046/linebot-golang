package main

import (
	"net/http"
	model "go-bot-line/model"
	"log"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/webhook", func(c echo.Context) error {
		Line := new(model.LineMessageUser)
		if err := c.Bind(Line); err != nil {
			log.Println("err")
			return c.String(http.StatusOK, "error")
		}
		var message = Line.Events[0].Message.Text
		var Verifytoken = ""
		var ChannalToken = ""

		bot, err := linebot.New(Verifytoken, ChannalToken)
		if err != nil {
			log.Println("err")
			return c.String(http.StatusOK, "error")
		}

		if _, err := bot.ReplyMessage(Line.Events[0].ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
			log.Println("%% err ReplyMessage")
		}
		log.Println("%% message success")
		return c.String(http.StatusOK, "ok")
		
	})

	e.Logger.Fatal(e.Start(":1323"))
}