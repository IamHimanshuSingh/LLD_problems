package models

import "time"

type Test struct {
	ID           int
	Title        string
	Subject      string
	TimeLimit    int
	Instructions string
	Questions    []Question
	CreatedAt    time.Time
}
