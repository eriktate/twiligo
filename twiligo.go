package twiligo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Client houses Twilio account information to be used in AuthN/AuthZ and provides all methods
// for interacting with the Twilio REST API.
type Client struct {
	baseURL    string
	http       *http.Client
	serviceSID string
	sid        string
	token      string

	logger *log.Logger
}

// NewClient creates a new Client and returns its pointer.
func NewClient(baseURL, sid, serviceSID, token string) *Client {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	return &Client{
		baseURL:    baseURL,
		sid:        sid,
		serviceSID: serviceSID,
		token:      token,
		http:       http.DefaultClient,
		logger:     logger,
	}
}

func (c *Client) post(path string, payload []byte, headers map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/Services/%s/%s", c.baseURL, c.serviceSID, path)
	log.Println("Posting to URL: %s", url)
	log.Println("POST body:\n%s", string(payload))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.SetBasicAuth(c.sid, c.token)

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("Post failed: %s", string(data))
	}

	log.Printf("Returned POST data:\n%s", string(data))
	return data, nil
}

func (c *Client) get(path string, headers map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/Services/%s/%s", c.baseURL, c.serviceSID, path)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.SetBasicAuth(c.sid, c.token)

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	log.Printf("Returned GET data:\n%s", string(data))
	return data, nil
}

func (c *Client) delete(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/Services/%s/%s", c.baseURL, c.serviceSID, path)
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.sid, c.token)
	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	log.Printf("Returned DELETE data:\n%s", string(data))
	return data, nil
}
