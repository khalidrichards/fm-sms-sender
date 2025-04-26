package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	client := twilio.NewRestClient()

	twilio_phone_number := os.Getenv("TWILIO_PHONE_NUMBER")
	params := &api.CreateMessageParams{}
	params.SetBody("This is a Flatbush Mixtape SMS Test")
	params.SetFrom(twilio_phone_number)
	params.SetTo("+18888888888") // Replace with the recipient's phone number

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}
}
