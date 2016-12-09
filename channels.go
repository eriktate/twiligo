package twiligo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Channel is the structured representation of a Twilio channel.
type Channel struct {
	SID          string     `json:"sid"`
	AccountSID   string     `json:"account_sid"`
	ServiceSID   string     `json:"service_sid"`
	UniqueName   string     `json:"unique_name,omitempty"`
	FriendlyName string     `json:"friendly_name,omitempty"`
	Attributes   string     `json:"attributes,omitempty"` // This is weirdly a string even though it's expected to be a JSON object.
	Type         string     `json:"type"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	CreatedBy    string     `json:"created_by,omitempty"`
	URL          string     `json:"url,omitempty"`
	Links        Link       `json:"links,omitempty"`
}

// ChannelsResponse is the structured representation of a response from Twilio for multiple channels.
type ChannelsResponse struct {
	Channels []Channel `json:"channels"`
}

// Link is the structured representation of a Twilio response link.
type Link struct {
	Members  string `json:"members,omitempty"`
	Messages string `json:"messages,omitempty"`
}

// Channel retrieves a specific Channel from Twilio. The id can either be the SID for the channel
// or the Unique Name assigned to the channel.
func (c *Client) Channel(id string) (Channel, error) {
	var channel Channel
	data, err := c.get(fmt.Sprintf("Channels/%s", id), nil)

	if err != nil {
		return channel, err
	}

	if err := json.Unmarshal(data, &channel); err != nil {
		return channel, err
	}

	return channel, nil
}

// Channels returns all channels currently tied to the given
func (c *Client) Channels() ([]Channel, error) {
	data, err := c.get("Channels", nil)

	if err != nil {
		return nil, err
	}

	var channelRes ChannelsResponse
	err = json.Unmarshal(data, &channelRes)

	if err != nil {
		return nil, err
	}

	// TODO: Iterate through pages if there are any.

	return channelRes.Channels, nil
}

// CreateChannel creates a new Channel in Twilio.
func (c *Client) CreateChannel(channel Channel) (Channel, error) {
	var newChannel Channel
	form := url.Values{}
	form.Add("FriendlyName", channel.FriendlyName)
	form.Add("UniqueName", channel.UniqueName)
	form.Add("Attributes", channel.Attributes)
	form.Add("Type", channel.Type)

	payload := []byte(form.Encode())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	data, err := c.post("Channels", payload, headers)

	if err != nil {
		return newChannel, err
	}

	if err := json.Unmarshal(data, &newChannel); err != nil {
		return newChannel, err
	}
	return newChannel, nil
}

// UpdateChannel updats an existing Channel in Twilio.
func (c *Client) UpdateChannel(channel Channel) (Channel, error) {
	var updatedChannel Channel
	form := url.Values{}
	form.Add("FriendlyName", channel.FriendlyName)
	form.Add("UniqueName", channel.UniqueName)
	form.Add("Attributes", channel.Attributes)

	payload := []byte(form.Encode())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	data, err := c.post(fmt.Sprintf("Channels/%s", channel.SID), payload, headers)

	if err != nil {
		return updatedChannel, err
	}

	if err := json.Unmarshal(data, &updatedChannel); err != nil {
		return updatedChannel, err
	}

	return updatedChannel, nil
}

// DeleteChannel deletes a Channel from Twilio.
func (c *Client) DeleteChannel(sid string) error {
	data, err := c.delete(fmt.Sprintf("Channels/%s", sid))

	if err != nil {
		return err
	}
	if len(data) > 0 {
		return fmt.Errorf("Received data in body of DELETE: %s", string(data))
	}

	return nil
}

// NewChannel creates a new Channel with required fields.
func NewChannel(friendlyName, uniqueName, attributes, chanType string) Channel {
	return Channel{
		FriendlyName: friendlyName,
		UniqueName:   uniqueName,
		Attributes:   attributes,
		Type:         chanType,
	}
}
