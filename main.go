package main
 
import (
    "fmt"
    "os"
    "time"
 
    "github.com/joho/godotenv"
    "github.com/slack-go/slack"
)
 
func main() {
 
    godotenv.Load(".env")
 
    token := os.Getenv("slack_auth_token")
    channelID := os.Getenv("slack_channel_id")
 
    client := slack.New(token, slack.OptionDebug(true))
    attachment := slack.Attachment{
        Pretext: "Super Bot Message",
        Text:    "some text",
        Color: "4af030",
        Fields: []slack.AttachmentField{
            {
                Title: "Date",
                Value: time.Now().String(),
            },
        },
    }
 
    _, timestamp, err := client.PostMessage(
        channelID,
 
        slack.MsgOptionAttachments(attachment),
    )
 
    if err != nil {
        panic(err)
    }
    fmt.Printf("Message sent at %s", timestamp)
}
