package scoringstrategy

import "testassessment/models"

type ScoringStrategy interface {
	CalculateScore(attempt *models.TestAttempt) (float64, error)
}
