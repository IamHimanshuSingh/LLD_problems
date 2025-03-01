package models

type Question struct {
	ID            int
	Text          string
	Type          string
	Options       []string
	CorrectAnswer string
	MaxScore      int
}
