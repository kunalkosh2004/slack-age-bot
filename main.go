package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	os.Setenv("SLACK_BOT_TOKEN","xoxb-8711473533798-8711522503654-v5vBzrBJt6zAdxeXcMgI1Cet")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A08MSNXUT1N-8729349512481-d925822317f0996ab0467718404d680dda3a0cef2089d611f57a35d13a362d77")

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