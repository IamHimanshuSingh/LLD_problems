package data

import (
	"fmt"
	"testassessment/models"
)

var testByQuestionsMap map[int][]models.Question
var studentsMap map[int]models.Student
var testsMap map[int]models.Test
var questions map[int]models.Question
var testAttemptMaps map[int]models.TestAttempt

type InMemoryDB interface {
	AddStudent(student models.Student) error
	AddQuestion(question models.Question) error
	AddTest(test models.Test) error
	AddTestAttempt(ta models.TestAttempt) error
	AddQuestionsToTest(testID int, questions []models.Question) error
	FetchTestAttemptData(testAttemptID int) (*models.TestAttempt, error)
	FetchQuestionDetails(questionID int) (*models.Question, error)
}

type InMemoryDBHandler struct {
}

func init() {
	testByQuestionsMap = make(map[int][]models.Question)
	studentsMap = make(map[int]models.Student)
	testsMap = make(map[int]models.Test)
	questions = make(map[int]models.Question)
	testAttemptMaps = make(map[int]models.TestAttempt)
}

func (*InMemoryDBHandler) AddStudent(student models.Student) error {
	studentsMap[student.ID] = student
	return nil
}

func (*InMemoryDBHandler) AddQuestion(question models.Question) error {
	questions[question.ID] = question
	return nil
}

func (*InMemoryDBHandler) AddTest(test models.Test) error {
	testsMap[test.ID] = test
	return nil
}

func (*InMemoryDBHandler) AddTestAttempt(ta models.TestAttempt) error {
	testAttemptMaps[ta.ID] = ta
	return nil
}

func (*InMemoryDBHandler) AddQuestionsToTest(testID int, questionsList []models.Question) error {
	testByQuestionsMap[testID] = questionsList
	return nil
}

func (*InMemoryDBHandler) FetchTestAttemptData(testAttemptID int) (*models.TestAttempt, error) {
	testAttempt, ok := testAttemptMaps[testAttemptID]
	if !ok {
		return nil, fmt.Errorf("attempt ID not found")
	}
	return &testAttempt, nil
}

func (*InMemoryDBHandler) FetchQuestionDetails(questionID int) (*models.Question, error) {
	question, ok := questions[questionID]
	if !ok {
		return nil, fmt.Errorf("question ID not found")
	}
	return &question, nil
}
