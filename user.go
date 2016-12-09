package twiligo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type User struct {
	SID          string     `json:"sid,omitempty"`
	AccountSID   string     `json:"account_sid,omitempty"`
	ServiceSID   string     `json:"service_sid,omitempty"`
	RoleSID      string     `json:"role_sid,omitempty"`
	Identity     string     `json:"identity,omitempty"`
	FriendlyName string     `json:"friendly_name,omitempty"`
	Attributes   string     `json:"attributes,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Online       bool       `json:"is_online,omitempty"`
	Notifiable   bool       `json:"is_notifiable,omitempty"`
	URL          string     `json:"url,omitempty"`
}

// NewUser creates a new instances of a User with the required fields.
func NewUser(identity, friendlyName, attributes, roleSID string) User {
	return User{
		Identity:     identity,
		FriendlyName: friendlyName,
		Attributes:   attributes,
		RoleSID:      roleSID,
	}
}

// CreateUser creates a new User in Twilio.
func (c *Client) CreateUser(user User) (User, error) {
	var createdUser User

	form := url.Values{}
	form.Add("Identity", user.Identity)
	form.Add("FriendlyName", user.FriendlyName)
	form.Add("Attributes", user.Attributes)
	form.Add("RoleSid", user.RoleSID)
	payload := []byte(form.Encode())

	data, err := c.postResource("Users", payload, getFormHeader())

	if err != nil {
		return createdUser, err
	}

	if err := json.Unmarshal(data, &createdUser); err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

// User retrieves a specific User from Twilio.
func (c *Client) User(identity string) (User, error) {
	var user User

	data, err := c.getResource(fmt.Sprintf("Users/%s", identity), nil)

	if err != nil {
		return user, err
	}

	if err = json.Unmarshal(data, &user); err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser updates an existing User in Twilio.
func (c *Client) UpdateUser(user User) (User, error) {
	var updatedUser User

	form := url.Values{}
	form.Add("FriendlyName", user.FriendlyName)
	form.Add("Attributes", user.Attributes)
	form.Add("RoleSid", user.RoleSID)
	payload := []byte(form.Encode())

	data, err := c.postResource(fmt.Sprintf("Users/%s", user.SID), payload, getFormHeader())

	if err != nil {
		return user, err
	}

	if err = json.Unmarshal(data, &updatedUser); err != nil {
		return user, err
	}

	return user, nil
}

// DeleteUser deletes an existing User from Twilio.
func (c *Client) DeleteUser(sid string) error {
	data, err := c.deleteResource(fmt.Sprintf("Users/%s", sid))

	if err != nil {
		return err
	}

	if len(data) > 0 {
		return fmt.Errorf("Received data in body of DELETE: %s", string(data))
	}

	return nil
}
