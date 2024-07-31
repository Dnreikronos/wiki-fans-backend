package models

type Character struct {
	ID          int64  `json:"id"`
	Age         int64  `json:"age"`
	Image       byte   `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Height      string `json:"height"`
	Abilities   string `json:"abilities"`
	Team        string `json:"team"`
	Coach       string `json:"coach"`
	Gender      string `json:"Gender"`
	Weight      string `json:"weight"`
}
