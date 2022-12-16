package main

import (
	"discord-as-cdn/src/media"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Return struct {
	Url string `json:"url"`
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/sendFile", func(ctx *fiber.Ctx) error {
		dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

		if err != nil {
			fmt.Println("error creating Discord session,", err)
			return fiber.NewError(fiber.StatusBadRequest, "error creating Discord session")
		}

		err = dg.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return fiber.NewError(fiber.StatusBadRequest, "error opening connection,")

		}

		file, err := ctx.FormFile("file")

		if err != nil {
			fmt.Println("Error getting file", err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Error getting file")
		}
		// do the checks and returns a file to be sent to discord
		osFile, err := media.ToOSFile(ctx, file)
		if err != nil {
			fmt.Println("Error saving the file", err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Error saving the file")
		}

		data := &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Name:        file.Filename,
					ContentType: file.Header.Get("Content-Type"),
					Reader:      osFile,
				}},
		}

		msg, err := dg.ChannelMessageSendComplex(os.Getenv("CHANNEL_ID"), data)

		if err != nil {
			fmt.Println("Error sending the message ", err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Error sending the message")
		}

		r := new(Return)
		r.Url = msg.Attachments[0].URL

		// clean up and return
		if err := os.Remove(media.Path); err != nil {
			fmt.Println("Was not able to remove file", err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Was not able to remove file")
		}

		defer func(dg *discordgo.Session) {
			err := dg.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(dg)
		return ctx.JSON(r)
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
