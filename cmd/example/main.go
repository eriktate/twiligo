package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/eriktate/twiligo"
)

func main() {
	accountSID := os.Getenv("TWILIO_ACCOUNTSID")
	serviceSID := os.Getenv("TWILIO_SERVICESID")
	token := os.Getenv("TWILIO_AUTHTOKEN")

	// Get a client ready.
	client := twiligo.NewClient("https://chat.twilio.com/v1", accountSID, serviceSID, token)

	// Creating channels
	log.Println("CREATING CHANNELS")
	_, err := client.CreateChannel(twiligo.NewChannel("Test Channel", "test-channel", "", "service"))
	if err != nil {
		log.Printf("Failed to create channel: %s", err)
	}
	_, err = client.CreateChannel(twiligo.NewChannel("New Channel", "new-channel", "", "service"))
	if err != nil {
		log.Printf("Failed to create channel: %s", err)
	}

	// Retrieving channels
	log.Println("RETRIEVING CHANNELS")
	channels, err := client.Channels()
	if err != nil {
		log.Printf("Failed to get channels: %s", err)
	}

	data, _ := json.Marshal(&channels)

	log.Printf("Channels from Twilio:/n%s", string(data))

	// Updating a channel
	log.Println("UPDATING CHANNELS")
	channels[0].FriendlyName = "Updated"
	_, err = client.UpdateChannel(channels[0])

	if err != nil {
		log.Printf("Failed to update channel: %s", err)
	}

	// Retrieving a specific channel
	log.Println("RETRIEVING SPECIFIC CHANNEL")
	channel, err := client.Channel(channels[0].SID)
	if err != nil {
		log.Printf("Failed to get channel: %s", err)
	}

	data, _ = json.Marshal(&channel)

	log.Printf("Channel from Twilio:/n%s", string(data))

	//BEGINNING MESSAGE EXAMPLE
	channelSID := channels[0].SID

	// Send Message
	log.Println("SEND MESSAGE")
	message := twiligo.NewMessage("This is a test", "", "system")
	sentMessage, err := client.SendMessage(channelSID, message)
	if err != nil {
		log.Printf("Failed to send message: %s", err)
	}
	message.Body = "This is also a test"
	_, err = client.SendMessage(channelSID, message)

	// Get Message
	getMessage, err := client.Message(channelSID, sentMessage.SID)
	if err != nil {
		log.Printf("Failed to get message: %s", err)
	}

	// Update Message
	getMessage.Body = "Actually this wasn't a test"
	_, err = client.UpdateMessage(channelSID, getMessage)
	if err != nil {
		log.Printf("Failed to update message: %s", err)
	}

	// Get Messages
	messages, err := client.Messages(channelSID)
	if err != nil {
		log.Printf("Failed to get messages: %s", err)
	}

	// Deleting Messages
	for _, m := range messages {
		err = client.DeleteMessage(channelSID, m.SID)

		if err != nil {
			log.Printf("Failed to delete message %s: %s", m.SID, err)
		}
	}

	// Deleting channels
	log.Println("DELETING CHANNELS")
	for _, ch := range channels {
		err = client.DeleteChannel(ch.SID)

		if err != nil {
			log.Printf("Failed to delete channel %s: %s", ch.SID, err)
		}
	}

	// BEGINNING USER EXAMPLE
	// Create a User
	user1, err := client.CreateUser(twiligo.NewUser("test@redventures.com", "Tester", "", ""))
	if err != nil {
		log.Printf("Failed to create user: %s", err)
	}

	// Update a User
	user1.FriendlyName = "SUPER Tester"
	_, err = client.UpdateUser(user1)
	if err != nil {
		log.Printf("Failed to update user: %s", err)
	}

	user1, err = client.User("test@redventures.com")
	if err != nil {
		log.Printf("Failed to get user: %s", err)
	}

	if err = client.DeleteUser(user1.SID); err != nil {
		log.Printf("Failed to delete user %s: %s", user1.SID, err)
	}

	// BEGINNING SERVICE EXAMPLE
	log.Println("DOING SERVICE THINGS")
	// Create a Service
	_, err = client.CreateService(twiligo.NewService("Test Service"))
	if err != nil {
		log.Printf("Failed to create service: %s", err)
	}
	_, err = client.CreateService(twiligo.NewService("Also a test"))
	if err != nil {
		log.Printf("Failed to create service: %s", err)
	}

	// Get Services
	services, err := client.Services()
	if err != nil {
		log.Printf("Failed to get services: %s", err)
	}

	service, err := client.Service(services[0].SID)
	if err != nil {
		log.Printf("Failed to get service: %s", err)
	}

	// Update service
	service.FriendlyName = "Not a test"
	_, err = client.UpdateService(service)
	if err != nil {
		log.Printf("Failed to update service: %s", err)
	}

	// Delete Services
	for _, s := range services {
		if s.FriendlyName != "twilio-dev" {
			if err = client.DeleteService(s.SID); err != nil {
				log.Printf("Failed to delete service %s: %s", s.SID, err)
			}
		}
	}
}
