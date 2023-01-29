package entities

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"time"
// )

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
	Tags             []string `json:"tags"`
	Owner            Owner    `json:"owner,omitempty"`
	IsAnswered       bool     `json:"is_answered"`
	ViewCount        int      `json:"view_count"`
	AnswerCount      int      `json:"answer_count"`
	Score            int      `json:"score"`
	LastActivityDate int      `json:"last_activity_date"`
	CreationDate     int      `json:"creation_date"`
	QuestionID       int      `json:"question_id"`
	ContentLicense   string   `json:"content_license"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
	LastEditDate     int      `json:"last_edit_date,omitempty"`
	Owner0           Owner0   `json:"owner,omitempty"`
	Owner1           Owner1   `json:"owner,omitempty"`
	Owner2           Owner2   `json:"owner,omitempty"`
	Owner3           Owner3   `json:"owner,omitempty"`
	AcceptedAnswerID int      `json:"accepted_answer_id,omitempty"`
}

type Stack struct {
	Items          []Items `json:"items"`
	HasMore        bool    `json:"has_more"`
	QuotaMax       int     `json:"quota_max"`
	QuotaRemaining int     `json:"quota_remaining"`
}