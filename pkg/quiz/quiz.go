package quiz

import (
	"fmt"
	"strings"
)

type QAP struct {
	Question string
	Answer   string
}

type QuestionBank struct {
	Questions        []QAP
	RightAnswerCount int
}

type Quiz interface {
	GetQuestions() []string
	AskQuestion(index int) (question string, err error)
	ValidateAnswer(index int, answer string) (correct bool, err error)
	GetResults() (RightAnswerCount int, qCount int)
	RunQuiz()
}

func (q *QuestionBank) GetQuestions() []string {
	questions := make([]string, len(q.Questions))
	for data := range q.Questions {
		questions[data] = q.Questions[data].Question
	}
	return questions
}
func (q *QuestionBank) AskQuestion(index int) (question string, err error) {
	question = q.Questions[index].Question
	fmt.Printf("%s = ", question)
	return
}

func (q *QuestionBank) ValidateAnswer(index int, answer string) (correct bool, err error) {
	isCorrect := q.Questions[index].Answer == answer
	if isCorrect {
		q.RightAnswerCount++
	}
	return isCorrect, nil

}

func (q *QuestionBank) ShowResults() (RightAnswerCount int, qCount int) {
	qCount = len(q.Questions)
	fmt.Printf("You got %d out of %d correct.\n", q.RightAnswerCount, qCount)
	return q.RightAnswerCount, qCount

}

func MakeQuestionBank(data []string) *QuestionBank {
	QAPs := []QAP{}
	for _, v := range data {
		qap := QAP{}
		qaPair := strings.Split(v, ",")
		qap.Question = qaPair[0]
		qap.Answer = qaPair[1]
		QAPs = append(QAPs, qap)

	}
	return &QuestionBank{
		Questions:        QAPs,
		RightAnswerCount: 0,
	}
}

func (q *QuestionBank) RunQuiz(quizChan chan<- string) {
	for i := 0; i < len(q.GetQuestions()); i++ {
		_, err := q.AskQuestion(i)
		if err != nil {
			fmt.Println(err)
			quizChan <- "Error!"
		}
		var answer string
		_, err = fmt.Scanln(&answer)
		if err != nil {
			fmt.Println(err)
			quizChan <- "Error!"
		}
		_, err = q.ValidateAnswer(i, answer)
		if err != nil {
			fmt.Println(err)
			quizChan <- "Error!"
		}
	}
	quizChan <- "Done!"
}
