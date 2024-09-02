package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	AccountSid = "YOUR_ACCOUNTSID"
	AuthToken  = "YOUR_AUTH_TOKEN"
	TwilioNum  = "YOUR_TWILIO_NUM"
	msg        = "this is the msg"
)

func (s *Service) ScheduleTasks() {
	s.scheduler.GetScheduler().Every(1).Day().At("09:00").Do(s.SendReminder)
	s.scheduler.GetScheduler().StartAsync()
}

func (s *Service) SendReminder() {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: AccountSid,
		Password: AuthToken,
	})

	phoneNo, err := s.store.UserPhoneNumber()
	if err != nil {
		log.Printf("error retrieving phone numbers: %s", err.Error())
		return
	}

	for _, v := range phoneNo {
		params := &twilioApi.CreateMessageParams{}
		params.SetTo(fmt.Sprint(v))
		params.SetFrom(TwilioNum) // Your Twilio phone number
		params.SetBody(msg)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			log.Printf("error sending message to %s: ", err.Error())
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}
	}
}
