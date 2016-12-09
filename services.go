package twiligo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// TODO: This is missing some things because the docs are not clear on structure. Need to research.

// Service is the structured representation of a Twilio Service.
type Service struct {
	SID                          string      `json:"sid,omitempty"`
	AccountSID                   string      `json:"account_sid,omitempty"`
	FriendlyName                 string      `json:"friendly_name,omitempty"`
	DateCreated                  *time.Time  `json:"date_created,omitempty"`
	DateUpdated                  *time.Time  `json:"date_updated,omitempty"`
	DefaultServiceRoleSID        string      `json:"default_service_role_sid,omitempty"`
	DefaultChannelRoleSID        string      `json:"default_channel_role_sid,omitempty"`
	DefaultChannelCreatorRoleSID string      `json:"default_channel_creator_rolesid,omitempty"`
	TypingIndicatorTimeout       int         `json:"typing_indicator_timeout,omitempty"`
	ReadStatusEnabled            bool        `json:"read_status_enabled,omitempty"`
	ConsumptionReportInterval    int         `json:"consumption_report_interval,omitempty"`
	ReachabilityEnabled          bool        `json:"reachability_enabled,omitempty"`
	PreWebhookURL                string      `json:"pre_webhook_url,omitempty"`
	PostWebhookURL               string      `json:"post_webhook_url,omitempty"`
	WebhookMethod                string      `json:"webhook_method,omitempty"`
	WebhookFilters               string      `json:"webhook_filters,omitempty"`
	URL                          string      `json:"url,omitempty"`
	Links                        ServiceLink `json:"links,omitempty"`
}

// ServiceLink is the structured representation of a Link within a Service.
type ServiceLink struct {
	Channels string `json:"channels,omitempty"`
	Roles    string `json:"roles,omitempty"`
	Users    string `json:"users,omitempty"`
}

// ServicesResponse is the structured representation of the response Twilio issues when asking for a list of Services.
type ServicesResponse struct {
	Services []Service `json:"services,omitempty"`
}

// NewService returns a Service with the required fields.
func NewService(friendlyName string) Service {
	return Service{
		FriendlyName: friendlyName,
	}
}

// Service retrieves a specific Service from Twilio given its SID.
func (c *Client) Service(sid string) (Service, error) {
	var service Service

	data, err := c.getService(fmt.Sprintf("Services/%s", sid), nil)

	if err != nil {
		return service, err
	}

	if err = json.Unmarshal(data, &service); err != nil {
		return service, err
	}

	return service, nil
}

// Services retrieves a list of all Services from Twilio.
func (c *Client) Services() ([]Service, error) {
	var serviceRes ServicesResponse

	data, err := c.getService("", nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serviceRes); err != nil {
		return nil, err
	}

	return serviceRes.Services, nil
}

// CreateService creates a new Service within Twilio.
func (c *Client) CreateService(service Service) (Service, error) {
	var createdService Service

	forms := url.Values{}
	forms.Add("FriendlyName", service.FriendlyName)
	payload := []byte(forms.Encode())

	data, err := c.postService("", payload, getFormHeader())

	if err != nil {
		return createdService, err
	}

	if err = json.Unmarshal(data, &createdService); err != nil {
		return createdService, err
	}

	return createdService, nil
}

// TODO: See TODO for Service struct. Same problem.

// UpdateService updates an existing Service in Twilio.
func (c *Client) UpdateService(service Service) (Service, error) {
	var updatedService Service

	forms := url.Values{}
	forms.Add("FriendlyName", service.FriendlyName)
	forms.Add("DefaultChannelRoleSid", service.DefaultChannelRoleSID)
	forms.Add("DefaultChannelCreatorRoleSid", service.DefaultChannelCreatorRoleSID)
	forms.Add("ReadStatusEnabled", strconv.FormatBool(service.ReadStatusEnabled))
	forms.Add("ReachabilityEnabled", strconv.FormatBool(service.ReachabilityEnabled))
	forms.Add("ConsumptionReportInterval", strconv.Itoa(service.ConsumptionReportInterval))
	forms.Add("TypingIndicatorTimeout", strconv.Itoa(service.TypingIndicatorTimeout))
	forms.Add("PreWebhookUrl", service.PreWebhookURL)
	forms.Add("PostWebhookUrl", service.PostWebhookURL)
	forms.Add("WebhookMethod", service.WebhookMethod)
	forms.Add("WebhookFilters", service.WebhookFilters)
	payload := []byte(forms.Encode())

	data, err := c.postService(service.SID, payload, getFormHeader())

	if err != nil {
		return updatedService, err
	}

	if err = json.Unmarshal(data, &updatedService); err != nil {
		return updatedService, err
	}

	return updatedService, nil
}

// DeleteService deletes a Service from Twilio given its SID.
func (c *Client) DeleteService(sid string) error {
	data, err := c.deleteService(sid)

	if err != nil {
		return err
	}

	if len(data) > 0 {
		return fmt.Errorf("Received data in body of DELETE: %s", string(data))
	}

	return nil
}
