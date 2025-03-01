package scoringstrategy

import (
	"testassessment/data"
	"testassessment/models"
)

type SimpleScoreCalculationStrategy struct {
	db data.InMemoryDB
}

func NewSimpleScoreCalculationStrategy(db data.InMemoryDB) ScoringStrategy {
	return &SimpleScoreCalculationStrategy{
		db: db,
	}
}

var questionTypeMarks = map[string]int{
	"MCQ":                5,
	"FILL_IN_THE_BLANKS": 2,
	"TRUE_FALSE":         3,
}

func (s SimpleScoreCalculationStrategy) CalculateScore(attempt *models.TestAttempt) (float64, error) {
	marksScored := 0
	totalScore := 0
	for _, answer := range attempt.Answers {
		// fetch question details and type
		questionDetails, err := s.db.FetchQuestionDetails(answer.QuestionID)
		if err != nil {
			return -1, err
		}
		totalScore += getMarksBasedOnQuestionType(questionDetails.Type)

		// Marks Scored added
		if answer.IsCorrect {
			marksScored += getMarksBasedOnQuestionType(questionDetails.Type)
		}
	}
	return float64(marksScored) / float64(totalScore) * 100, nil
}

func getMarksBasedOnQuestionType(qType string) int {
	return questionTypeMarks[qType]
}
