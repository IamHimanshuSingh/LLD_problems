# Online Assessment Platform

## 90-Minute Machine Coding Problem

### 📚 Problem Definition  
Design an **Online Assessment Platform** where students can take tests with various question types. The platform should automatically evaluate responses for certain question types and calculate scores based on question weights.  

---

## ✅ Required Requirements  

### **1. Supported Question Types:**  
- Multiple Choice (single correct answer)  
- True/False  
- Fill in the Blank (exact match text answer)  

### **2. Basic Functionalities:**  
- **Add Student** (id, name, email)  
- **Create Test** (id, title, subject, time_limit_minutes)  
- **Add Question to Test** (test_id, question_text, question_type, options, correct_answer, max_score)  
- **Submit Test Attempt** (student_id, test_id, answers)  
- **Get Test Results** (student_id, test_id)  
- **Get Student Test History** (student_id)  

### **3. Score Calculation:**  
- **Correct Answer** → Student receives the **maximum score** for that question.  
- **Incorrect Answer** → Student receives **0 points**.  
- **Final Score Calculation:**  
  ```
  (Total points earned / Total possible points) * 100
  ```

---

## 🎯 Bonus Requirements (If Time Permits)  

### **1. Additional Question Types:**  
- Multiple Select (multiple correct answers)  
- Numeric Answer (with tolerance range)  

### **2. Partial Credit System:**  
- For **Multiple Select** questions, award **partial credit** based on correct selections.  
- **Penalty** for incorrect options (-25% of max score per wrong option).  

### **3. Additional Functionalities:**  
- **Time-limited tests** with automatic submission.  
- **Question difficulty levels** (Easy, Medium, Hard) with auto-scoring adjustments.  
- **Item analysis** (success rate per question).  
- **Performance analytics** (average score, highest score, lowest score, etc.).  

---

## 🔹 Assumptions  
1. All tests contain at least **one question**.  
2. Each question has a **positive integer max score**.  
3. Tests can be **taken only once** by each student.  
4. Time limits are specified in **minutes**.  

---

## 📝 Test Cases  

### **1. Basic Multiple Choice Test**  
- **Scenario:** Student S takes a test with **3 multiple-choice questions** (max score **5 each**).  
- **Input:** S submits answers for `test_id 101`:  
  ```json
  { "question1": "B", "question2": "A", "question3": "D" }
  ```
- **Correct Answers:** `B, C, D`  
- **Output:**  
  ```
  Score: 10/15 (66.67%)
  Question 1: Correct (5/5)
  Question 2: Incorrect (0/5)
  Question 3: Correct (5/5)
  ```

### **2. Mixed Question Types**  
- **Scenario:** Student S takes a test with **1 multiple-choice, 1 true/false, and 1 fill-in-the-blank question**.  
- **Input:** S submits answers for `test_id 102`:  
  ```json
  { "question1": "A", "question2": "True", "question3": "Python" }
  ```
- **Correct Answers:** `A, False, Python`  
- **Output:**  
  ```
  Score: 7/12 (58.33%)
  Question 1: Correct (5/5)
  Question 2: Incorrect (0/3)
  Question 3: Correct (2/2)
  ```

### **3. (Bonus) Multiple Select Question**  
- **Scenario:** Student S answers a **multiple select question** with **4 options**, where `A` and `C` are correct.  
- **Input:** S selects **A and B**.  
- **Output:**  
  ```
  Score: 2.5/10 (25%)
  Correct options selected: 1/2
  Incorrect options selected: 1/2
  ```

---

## 🚀 Expectations  
1. Create **minimal sample data** for students and tests.  
2. The code should be **demo-able** via a main method or simple test cases.  
3. Follow **basic Object-Oriented Design principles**.  
4. Handle **basic edge cases gracefully**.  
5. Code should be **legible and easy to understand**.  

---

## 🛠️ Guidelines  
1. **Do not access the internet** (except for syntax lookup).  
2. Use **any language** and **IDE of your choice**.  
3. The entire code should be **your own**.  
4. **Focus on required functionality first**, then attempt bonus features.  

---

## ⚠️ Basic Edge Cases to Handle  
1. Student attempting to **retake** a test.  
2. Submitting a test with **missing answers**.  
3. Adding a question to a **non-existent test**.  
4. Calculating the **score for a test with no questions**.  