package models

type TestAttempt struct {
	ID        int
	StudentID int
	TestID    int
	Answers   []Answer
	Score     int
	Completed bool
}

type Answer struct {
	QuestionID int
	Answer     string
	IsCorrect  bool
}
