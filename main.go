package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main()  {
	// os.Setenv("SLACK_BOT_TOKEN","your-token")
	// os.Setenv("SLACK_APP_TOKEN","your-token")
	godotenv.Load()

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		// Example: "my yob is 2004",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob ,err := strconv.Atoi(year)
			if err!=nil{
				fmt.Println("Error")
			}
			age := time.Now().Year()-yob
			r := fmt.Sprintf("Age is: %d",age)
			response.Reply(r)
		}, 
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err!=nil{
		log.Fatal(err)
	}

}