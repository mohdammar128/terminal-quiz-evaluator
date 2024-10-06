package quiz

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"unicode"
)

var (
	Instruction = `Instruction:
1. Each question is worth one mark.
2. If the answer is an integer value (e.g., 23, 45), write it as it is.
3. If the answer contains decimal digits more than two (e.g., 234.4567, 234.50, 345.6), write only up to two decimal places (e.g., 234.45, 234.5, 345.6).
4. Once you complete the quiz, you can print your score summary using the command (get score).
`
	TotalQuest = 0
)

type Quiz struct {
	instruction   string
	correct       int
	totalQuestion int
}

func trimDownUptoTwoDecimalPlaces(value float64) float64 {
	factor := math.Pow(10, 2)
	newVal := math.Floor(value*factor) / factor
	return newVal
}

func applyOperator(a float64, b float64, c string) float64 {
	switch c {
	case "*":
		return a * b
	case "/":
		return a / b
	case "+":
		return a + b
	case "-":
		return a - b
	}
	// just for sake of returning
	return -1
}

func precedence(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	if op == "*" || op == "/" {
		return 2
	}
	return 0
}

func convertToPostFixExpression(expr string) []string {
	opt := NewStack()
	opt.Push("$")
	var postFixExpression []string
	numBuffer := ""
	for _, e := range expr {
		c := string(e)
		if unicode.IsDigit(e) {
			numBuffer += c
		} else if numBuffer != "" {
			postFixExpression = append(postFixExpression, numBuffer)
			numBuffer = ""
		}

		if c == "+" || c == "-" || c == "*" || c == "/" {
			scannedPriority := precedence(c)
			val, _ := opt.Peek()
			stackOptPriority := precedence(val)
			for scannedPriority <= stackOptPriority {
				postFixExpression = append(postFixExpression, val)
				opt.Pop()
				val, _ = opt.Peek()
				stackOptPriority = precedence(val)
			}
			opt.Push(c)
		}
	}

	if numBuffer != "" {
		postFixExpression = append(postFixExpression, numBuffer)
	}

	val, _ := opt.Peek()
	for val != string("$") {
		postFixExpression = append(postFixExpression, val)
		opt.Pop()
		val, _ = opt.Peek()
	}
	return postFixExpression
}

func evaluatePostFixExpression(postFixExpression []string) float64 {
	evalStack := NewStack()
	for _, val := range postFixExpression {
		switch val {
		case "+", "-", "*", "/":
			num1Str, _ := evalStack.Pop()
			num2Str, _ := evalStack.Pop()

			num1, err1 := strconv.ParseFloat(num1Str, 64)
			num2, err2 := strconv.ParseFloat(num2Str, 64)

			if err1 != nil || err2 != nil {
				panic("error converting stack values to float")
			}

			res := applyOperator(num2, num1, val)
			evalStack.Push(strconv.FormatFloat(res, 'f', -1, 64))

		default:
			evalStack.Push(val)
		}
	}

	finalResStr, _ := evalStack.Peek()
	finalResFloat, _ := strconv.ParseFloat(finalResStr, 64)
	return finalResFloat
}

func finalRes(expr string) float64 {
	postFixExpr := convertToPostFixExpression(expr)
	res1 := evaluatePostFixExpression(postFixExpr)
	return res1
}

func ReadCsv(filePath string) {
	file, _ := os.Open(filePath)
	r := csv.NewReader(file)
	records, _ := r.ReadAll()
	TotalQuest = len(records)

	// fmt.Println("Hello Dear")
	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	defer file.Close()
}

func NewQuiz() *Quiz {
	return &Quiz{
		instruction:   Instruction,
		correct:       0,
		totalQuestion: TotalQuest,
	}
}

func (q *Quiz) GetTotalNumberOfQuestions() int {
	return q.totalQuestion
}
func (q *Quiz) GetNumberOfCorrectQuestion() int {
	return q.correct
}
