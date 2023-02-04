package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

// testing with const
const epoch = 1234567890
const layout =  "2006-01-02 15:04:05"

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




// read file ./basicstruct.json and open in Go
// this will be replaced by getting the json in a stream from the API

// 1. first create the struct from the data source
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
	Tags             []string  						`json:"tags"`
	Owner            Owner     						`json:"owner,omitempty"`
	IsAnswered       bool      						`json:"is_answered"`
	ViewCount        int       						`json:"view_count"`
	AnswerCount      int       						`json:"answer_count"`
	Score            int       						`json:"score"`
	LastActivityDate EpochConversion     	`json:"last_activity_date"`
	CreationDate     int 		   						`json:"creation_date"`
	QuestionID       int       						`json:"question_id"`
	ContentLicense   string    						`json:"content_license"`
	Link             string    						`json:"link"`
	Title            string    						`json:"title"`
	LastEditDate     int        					`json:"last_edit_date,omitempty"`
	Owner0           Owner0    						`json:"owner,omitempty"`
	Owner1           Owner1    						`json:"owner,omitempty"`
	Owner2           Owner2    						`json:"owner,omitempty"`
	Owner3           Owner3    						`json:"owner,omitempty"`
	AcceptedAnswerID int       						`json:"accepted_answer_id,omitempty"`
}

type Stack struct {
	Items          []Items `json:"items"`
	HasMore        bool    `json:"has_more"`
	QuotaMax       int     `json:"quota_max"`
	QuotaRemaining int     `json:"quota_remaining"`
}




func main() {
	stackFile, err := ioutil.ReadFile("./full-stack.json")
	if err != nil {
		log.Fatal("error when opening file:", err)
	}
	// println(content)
	var s Stack
	// err := json.Unmarshal([]byte(content), &o)
	err = json.Unmarshal(stackFile, &s)
	if err != nil {
		log.Fatal("error during umarshal(): ", err)
	}
  fmt.Println(s.QuotaMax)
	for _, i := range s.Items {
		fmt.Println(i.LastActivityDate)
	}

	// print the unmarshelled data
	// log.Printf("Item ID: %s\n", o.Id)
	// log.Printf("Date Ordered: %s\n", o.DateOrdered)
	// log.Printf("Customer ID: %s\n", o.CustomerId)
	// log.Printf("Items: %s\n", o.Items)

	fmt.Printf("Now printing unix time: \n")
	ut := time.Now().Unix()
	fmt.Println(ut)
	fmt.Printf("Now lets convert unix time to human readable\n")
	fmt.Println("The time was", time.Unix(epoch, 0).Format(time.RFC822Z))
}