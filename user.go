package twiligo

import "time"

type User struct {
	SID          string     `json:"sid,omitempty"`
	AccountSID   string     `json:"account_sid,omitempty"`
	ServiceSID   string     `json:"service_sid,omitempty"`
	Identity     string     `json:"identity,omitempty"`
	FriendlyName string     `json:"friendly_name,omitempty"`
	Attributes   string     `json:"attributes,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	Online       bool       `json:"is_online,omitempty"`
	Notifiable   bool       `json:"is_notifiable,omitempty"`
	URL          string     `json:"url,omitempty"`
}
