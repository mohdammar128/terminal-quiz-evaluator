package main

import (
	"calculator/quiz"
	"fmt"
	"strconv"
)

func main() {
	// file, err := os.Open("quiz.csv")
	// if err != nil {
	// 	log.Fatal("File is not able to open")
	// }

	// defer file.Close()
	// expression := csv.NewReader(file)
	// //it contains record (expr ,ans)
	// records, err := expression.ReadAll()
	// expr := "5+4*3/4"
	st := quiz.NewStack()
	st.Push(1)
	st.Push(3)
	size := st.Size()
	for i := 0; i < size; i++ {
		fmt.Println(st.Pop())
	}
	/*
		3 4 + 2 * 7 -
		7
	*/
	q := quiz.NewQueue()
	q.Add(1)
	q.Add(2)
	q.Add(89)
	q.Add(5)
	q.Add(10)
	fmt.Printf("%d", q.Size())
	// size2 := q.Size()
	for !q.Empty() {
		v, _ := q.Poll()
		fmt.Printf("%d\n", v)
	}
	// trimExpr := strings.TrimSpace(expr)
	// res1, res2 := evaluateExpression(trimExpr)
	// res := applyOperator(trimExpr)
	// fmt.Println(res1, res2, res)

}

func evaluateExpression(expr string) ([]int, []rune) {
	num := []int{}
	opt := []rune{}
	// var operator map[rune]int
	prev := -1
	for i, e := range expr {
		switch e {
		case '*':
			val, _ := strconv.Atoi(expr[prev+1 : i])
			num = append(num, val)
			opt = append(opt, '*')
			// operator['*'] = 1
			prev = i
		case '/':
			val, _ := strconv.Atoi(expr[prev+1 : i])
			num = append(num, val)
			opt = append(opt, '/')

			// operator['/'] = 1
			prev = i
		case '+':
			val, _ := strconv.Atoi(expr[prev+1 : i])
			num = append(num, val)
			opt = append(opt, '+')

			// operator['+'] = 2
			prev = i
		case '-':
			val, _ := strconv.Atoi(expr[prev+1 : i])
			num = append(num, val)
			opt = append(opt, '-')

			// operator['-'] = 2
			prev = i

		}
	}
	val, _ := strconv.Atoi(expr[prev+1:])
	num = append(num, val)
	return num, opt
}

func applyOperator(a int, b int, c string) int {
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
	//just for sake of returnning
	return -1
}

func calulate() {

}

//read csv file
//perform some logic
//store the answer
//when give two things based on option (result ,result with answer)
