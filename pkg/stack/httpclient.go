package stack
// http://localhost:8000/hello?user_id=300
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	// "github.com/gofiber/fiber/v2"
)

// Client .
type Client struct {
	apiKey     string
	baseURL    string
	HTTPClient *http.Client
}

// NewClient creates new Facest.io client with given API key
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: "https://api.stackexchange.com/2.2",
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Content-type and body should be already added to req
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}