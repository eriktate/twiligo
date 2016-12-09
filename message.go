package twiligo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Message is the structured representation of a Message in a Twilio Channel.
type Message struct {
	SID         string     `json:"sid,omitempty"`
	AccountSID  string     `json:"account_sid,omitempty"`
	ServiceSID  string     `json:"service_sid,omitempty"`
	To          string     `json:"to,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	WasEdited   bool       `json:"was_edited,omitempty"`
	From        string     `json:"from,omitempty"`
	Body        string     `json:"body,omitempty"`
	Attributes  string     `json:"attributes,omitempty"`
	Index       int        `json:"index,omitempty"`
	URL         string     `json:"url,omitempty"`
}

// MessagesResponse represents the structure of a response from Twilio for multiple Messages.
type MessagesResponse struct {
	Messages []Message `json:"messages,omitempty"`
	//TODO: Add Meta property
}

// NewMessage creates a new Message with the required fields.
func NewMessage(body, attributes, from string) Message {
	return Message{
		Body:       body,
		Attributes: attributes,
		From:       from,
	}
}

// Message retrieves an individual message from a Channel in Twilio.
func (c *Client) Message(channelSID, messageSID string) (Message, error) {
	var message Message
	data, err := c.get(fmt.Sprintf("Channels/%s/Messages/%s", channelSID, messageSID), nil)

	if err != nil {
		return message, err
	}

	if err := json.Unmarshal(data, &message); err != nil {
		return message, err
	}

	return message, nil
}

// Messages retrieves ALL Messages from a Channel in Twilio.
func (c *Client) Messages(channelSID string) ([]Message, error) {
	var messageRes MessagesResponse
	data, err := c.get(fmt.Sprintf("Channels/%s/Messages", channelSID), nil)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &messageRes); err != nil {
		return nil, err
	}

	return messageRes.Messages, nil
}

// SendMessage sends a Message to a Channel in Twilio.
func (c *Client) SendMessage(channelSID string, message Message) (Message, error) {
	var sentMessage Message

	form := url.Values{}
	form.Add("Body", message.Body)
	form.Add("Attributes", message.Attributes)
	form.Add("From", message.From)
	payload := []byte(form.Encode())
	data, err := c.post(fmt.Sprintf("Channels/%s/Messages", channelSID), payload, getFormHeader())

	if err != nil {
		return sentMessage, err
	}

	if err := json.Unmarshal(data, &sentMessage); err != nil {
		return sentMessage, err
	}

	return sentMessage, nil
}

// UpdateMessage updates a specific Message within a Channel in Twilio.
func (c *Client) UpdateMessage(channelSID string, message Message) (Message, error) {
	var updatedMessage Message

	form := url.Values{}
	form.Add("Body", message.Body)
	form.Add("Attributes", message.Attributes)
	payload := []byte(form.Encode())

	data, err := c.post(fmt.Sprintf("Channels/%s/Messages/%s", channelSID, message.SID), payload, getFormHeader())

	if err != nil {
		return updatedMessage, err
	}

	if err := json.Unmarshal(data, &updatedMessage); err != nil {
		return updatedMessage, err
	}

	return updatedMessage, nil
}

// DeleteMessage deletes a specific Message within a Channel in Twilio.
func (c *Client) DeleteMessage(channelSID, messageSID string) error {
	data, err := c.delete(fmt.Sprintf("Channels/%s/Messages/%s", channelSID, messageSID))

	if err != nil {
		return err
	}

	if len(data) > 0 {
		return fmt.Errorf("Received data in body of DELETE: %s", string(data))
	}

	return nil
}
