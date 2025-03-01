package models

import "time"

type TestResult struct {
	ID          int
	StudentID   int
	TestID      int
	Score       int
	CompletedAt time.Time
}
