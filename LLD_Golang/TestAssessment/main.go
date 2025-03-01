package main

import (
	"fmt"
	"testassessment/data"
	"testassessment/models"
	"testassessment/svc"
)

func main() {
	fmt.Println("Test Assessment System - Driver Program")
	fmt.Println("--------------------------------------")

	// Create an instance of the in-memory database handler
	dbHandler := data.InMemoryDBHandler{}

	// Add students
	fmt.Println("\nAdding students...")
	students := []models.Student{
		{ID: 101, Name: "Alice Johnson", Email: "alice@example.com"},
		{ID: 102, Name: "Bob Smith", Email: "bob@example.com"},
		{ID: 103, Name: "Carol Davis", Email: "carol@example.com"},
		{ID: 104, Name: "David Wilson", Email: "david@example.com"},
	}

	for _, student := range students {
		err := dbHandler.AddStudent(student)
		if err != nil {
			fmt.Printf("Error adding student %s: %v\n", student.Name, err)
		} else {
			fmt.Printf("Added student: %s (ID: %d, Email: %s)\n",
				student.Name, student.ID, student.Email)
		}
	}

	// Add questions
	fmt.Println("\nAdding questions...")
	questions := []models.Question{
		{
			ID:            1,
			Text:          "What is 2+2?",
			Options:       []string{"1", "2", "3", "4"},
			CorrectAnswer: "4",
			MaxScore:      5,
			Type:          "MCQ",
		},
		{
			ID:            2,
			Text:          "What is the capital of France?",
			Options:       []string{"Mumbai", "Delhi", "ABC", "Paris"},
			CorrectAnswer: "Paris",
			MaxScore:      5,
			Type:          "MCQ",
		},
		{
			ID:            3,
			Text:          "What is the powerhouse of cell?",
			Options:       []string{"Protein", "RBC", "WBC", "Mitochondria"},
			CorrectAnswer: "Mitochondria",
			MaxScore:      10,
			Type:          "MCQ",
		},
	}

	for _, question := range questions {
		err := dbHandler.AddQuestion(question)
		if err != nil {
			fmt.Printf("Error adding question ID %d: %v\n", question.ID, err)
		} else {
			fmt.Printf("Added question ID %d: %s (MaxScore: %d)\n",
				question.ID, question.Text, question.MaxScore)
		}
	}

	// Add tests
	fmt.Println("\nAdding tests...")
	mathTest := models.Test{
		ID:           1,
		Title:        "Basic Math Test",
		Subject:      "Mathematics",
		TimeLimit:    60,
		Instructions: "Answer all questions. Each question is worth 5 points.",
		Questions:    []models.Question{questions[0]},
	}

	csTest := models.Test{
		ID:           2,
		Title:        "Biology Fundamentals",
		Subject:      "Biology",
		TimeLimit:    90,
		Instructions: "Answer all questions. Points vary per question.",
		Questions:    []models.Question{questions[2]},
	}

	generalTest := models.Test{
		ID:           3,
		Title:        "General Knowledge",
		Subject:      "Mixed",
		TimeLimit:    120,
		Instructions: "Answer all questions to the best of your ability.",
		Questions:    []models.Question{questions[0], questions[1], questions[2]},
	}

	tests := []models.Test{mathTest, csTest, generalTest}

	for _, test := range tests {
		err := dbHandler.AddTest(test)
		if err != nil {
			fmt.Printf("Error adding test %s: %v\n", test.Title, err)
		} else {
			fmt.Printf("Added test: %s (ID: %d, Subject: %s, Time Limit: %d minutes)\n",
				test.Title, test.ID, test.Subject, test.TimeLimit)
		}

		// Add questions to test
		dbHandler.AddQuestionsToTest(test.ID, test.Questions)
		fmt.Printf("  Added %d questions to test ID %d\n", len(test.Questions), test.ID)
	}

	// feeding the test attempts to the system to calculate scores afterwards
	fmt.Println("\nAdding test attempts...")
	testAttempts := []models.TestAttempt{
		{
			ID:        1,
			StudentID: 101,
			TestID:    3,
			Answers: []models.Answer{
				{
					QuestionID: 1,
					Answer:     "3",
					IsCorrect:  false,
				},
				{
					QuestionID: 2,
					Answer:     "Paris",
					IsCorrect:  true,
				},
			},
			Completed: true,
		},
		{
			ID:        2,
			StudentID: 102,
			TestID:    1,
			Answers: []models.Answer{
				{
					QuestionID: 1,
					Answer:     "3",
					IsCorrect:  false,
				},
			},
			Completed: true,
		},
	}

	for _, attempt := range testAttempts {
		err := dbHandler.AddTestAttempt(attempt)
		if err != nil {
			fmt.Printf("Error adding test attempt ID %d: %v\n", attempt.ID, err)
		} else {
			status := "Completed"
			if !attempt.Completed {
				status = "In Progress"
			}
			fmt.Printf("Added test attempt: ID %d, Student ID %d, Test ID %d, Status: %s\n",
				attempt.ID, attempt.StudentID, attempt.TestID, status)
		}
	}

	fmt.Println("\nAll data successfully added to in-memory database!")

	// Calculate Score for each testAttempt
	testSvc := svc.NewTestSvcHandler(&dbHandler)

	for _, testAttempt := range testAttempts {
		score, err := testSvc.CalculateScore(testAttempt.ID)
		if err != nil {
			fmt.Printf("Error in Calculating Score: %v \n", err)
		}
		fmt.Printf("Score For TestAttempt: %+v, score: %f \n", testAttempt, score)
	}
}
