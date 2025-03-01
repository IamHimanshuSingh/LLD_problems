package svc

import (
	"testassessment/data"
	"testassessment/scoringstrategy"
)

type TestSvcHandler struct {
	db data.InMemoryDB
}

func NewTestSvcHandler(db data.InMemoryDB) *TestSvcHandler {
	return &TestSvcHandler{
		db: db,
	}
}

func (ts *TestSvcHandler) CalculateScore(testAttemptID int) (float64, error) {
	testAttempt, err := ts.db.FetchTestAttemptData(testAttemptID)
	if err != nil {
		return -1, err
	}
	scorecalculationstrategy := scoringstrategy.NewSimpleScoreCalculationStrategy(ts.db)
	return scorecalculationstrategy.CalculateScore(testAttempt)
}
