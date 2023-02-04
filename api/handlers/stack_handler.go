package handlers

import (
        // "context"
        // "encoding/json"
        "fmt"
        "github.com/gofiber/fiber/v2"
        "net/http"
        "time"
				"strings"
				"strconv"
)

// smoke test Hello endpoint
/* func GetStack(c *fiber.Ctx)  error {
        return c.SendString("Hello, World 👋!")
} */

const URL string = "https://api.stackexchange.com/2.2/search?order=desc&sort=activity&intitle=fiber&site=stackoverflow"


// custom unmarshaller for time format
type EpochConversion string
func(ec *EpochConversion) UnmarshalJSON(data []byte) error {
	t := strings.Trim(string(data), `"`)
	sec, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return err
	}
	epochTime := time.Unix(sec, 0).Format(time.RFC3339)
	// epochTime := time.Unix(sec, 0).Format(time.ANSIC)
	*ec = EpochConversion(epochTime)
	return nil
}


// buid custom structs --they should be in the order they appear in the main struct
type Owner struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

type Owner0 struct {
	UserType    string `json:"user_type"`
	DisplayName string `json:"display_name"`
}

type Owner1 struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	AcceptRate   int    `json:"accept_rate"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

type Owner2 struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	AcceptRate   int    `json:"accept_rate"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}
type Owner3 struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	AcceptRate   int    `json:"accept_rate"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}


type Items struct {
	Tags             []string  					`json:"tags"`
	Owner            Owner     					`json:"owner,omitempty"`
	IsAnswered       bool      					`json:"is_answered"`
	ViewCount        int       					`json:"view_count"`
	AnswerCount      int       					`json:"answer_count"`
	Score            int       					`json:"score"`
	LastActivityDate EpochConversion   `json:"last_activity_date"`
	CreationDate     int 		   					`json:"creation_date"`
	QuestionID       int       					`json:"question_id"`
	ContentLicense   string    					`json:"content_license"`
	Link             string    					`json:"link"`
	Title            string    					`json:"title"`
	LastEditDate     int        				`json:"last_edit_date,omitempty"`
	Owner0           Owner0    					`json:"owner,omitempty"`
	Owner1           Owner1    					`json:"owner,omitempty"`
	Owner2           Owner2    					`json:"owner,omitempty"`
	Owner3           Owner3    					`json:"owner,omitempty"`
	AcceptedAnswerID int       					`json:"accepted_answer_id,omitempty"`
}

type Stack struct {
	Items          []Items `json:"items"`
	HasMore        bool    `json:"has_more"`
	QuotaMax       int     `json:"quota_max"`
	QuotaRemaining int     `json:"quota_remaining"`
}




// new implementation of GetFullStack
func (c *Client) getStack(f *fiber.Ctx) (*Stack, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?order=desc&sort=activity&intitle=fiber&site=stackoverflow", c.baseURL), nil)
	if err != nil {
		return nil, err
	}
  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  res := Stack{}
	if err := c.sendRequest(req, &res); err !=nil {
		return nil, err
	}
   // swap example code return with endpoint code
	return &res, nil
}

func GetStack(c *fiber.Ctx)  error {
        return c.SendString("Hello, World 👋!")
}
