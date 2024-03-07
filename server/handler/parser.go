package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	grpcservice "github.com/parserSchedulerService/server/grpcservice"
	models "github.com/parserSchedulerService/server/model"
	services "github.com/parserSchedulerService/server/services"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenParenthesis
)

type Token struct {
	Type  TokenType
	Value string
}

func TRIGER_PARSING(c *gin.Context) {
	// expression := "(3 + 4) * 2"
	expression := models.ExpressionRequest{}

	if err := c.ShouldBindJSON(&expression); err != nil {
		log.Println("Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := EvaluateExpression(expression.Expression)

	if err != nil {
		log.Println("Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	str := strconv.FormatFloat(result, 'f', 2, 64)
	// put the result in the outbound sqs
	services.EnqueueRequestToOutboundSqs(c, str, expression.Uuid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}

func EvaluateExpression(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	rpn, err := shuntingYard(tokens)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}
func tokenize(expression string) ([]Token, error) {
	tokens := []Token{}
	var buffer strings.Builder
	previousToken := Token{Type: TokenOperator, Value: "+"} // Initialize with a dummy operator
	tokenType := TokenOperator
	for _, char := range expression {
		if unicode.IsSpace(char) {
			continue
		}

		if unicode.IsDigit(char) || char == '.' {
			buffer.WriteRune(char)
		} else {
			if buffer.Len() > 0 {
				tokens = append(tokens, Token{Type: TokenNumber, Value: buffer.String()})
				buffer.Reset()
			}

			tokenType = TokenOperator
			if char == '(' || char == ')' {
				tokenType = TokenParenthesis
			}

			// Check if the current character is a minus sign and the previous token is an operator or an opening parenthesis
			if char == '-' && (previousToken.Type == TokenOperator || previousToken.Value == "(") {
				buffer.WriteRune(char) // Treat as a negative sign and continue with the next character
				continue
			}

			tokens = append(tokens, Token{Type: tokenType, Value: string(char)})
		}

		previousToken = Token{Type: tokenType, Value: string(char)}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, Token{Type: TokenNumber, Value: buffer.String()})
	}

	return tokens, nil
}

// func tokenize(expression string) ([]Token, error) {
// 	tokens := []Token{}
// 	var buffer strings.Builder
// 	for _, char := range expression {
// 		if unicode.IsSpace(char) {
// 			continue
// 		}
// 		if unicode.IsDigit(char) || char == '.' {
// 			buffer.WriteRune(char)
// 		} else {
// 			if buffer.Len() > 0 {
// 				tokens = append(tokens, Token{Type: TokenNumber, Value: buffer.String()})
// 				buffer.Reset()
// 			}
// 			tokenType := TokenOperator
// 			if char == '(' || char == ')' {
// 				tokenType = TokenParenthesis
// 			}
// 			tokens = append(tokens, Token{Type: tokenType, Value: string(char)})
// 		}
// 	}
// 	if buffer.Len() > 0 {
// 		tokens = append(tokens, Token{Type: TokenNumber, Value: buffer.String()})
// 	}
// 	return tokens, nil
// }

func shuntingYard(tokens []Token) ([]Token, error) {
	output := []Token{}
	stack := []Token{}

	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			output = append(output, token)
		case TokenOperator:
			for len(stack) > 0 && stack[len(stack)-1].Type != TokenParenthesis &&
				precedence(token.Value) <= precedence(stack[len(stack)-1].Value) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case TokenParenthesis:
			if token.Value == "(" {
				stack = append(stack, token)
			} else {
				found := false
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					if top.Value == "(" {
						found = true
						break
					}
					output = append(output, top)
				}
				if !found {
					return nil, errors.New("mismatched parentheses")
				}
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Value == "(" || top.Value == ")" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, top)
	}

	return output, nil
}

func precedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func add(x, y float64) float64 {
	log.Printf("Adding %f and %f", x, y)
	return x + y
}

func subtract(x, y float64) float64 {
	log.Printf("Subtracting %f and %f", x, y)
	return x - y
}

func multiply(x, y float64) float64 {
	log.Printf("Multiplying %f and %f", x, y)
	return x * y
}

func divide(x, y float64) (float64, error) {
	log.Printf("Dividing %f and %f", x, y)
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func evaluateRPN(tokens []Token) (float64, error) {
	stack := []float64{}
	operations := make(chan float64)
	errorsChan := make(chan error)

	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			value, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, value)
		case TokenOperator:
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}

			rightOperand := stack[len(stack)-1]
			leftOperand := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			go func() {
				var result float64
				var temp float32
				var err error

				switch token.Value {
				case "+":
					// temp, _ = grpcservice.CallGrpcAdd(30001, float32(leftOperand), float32(rightOperand))
					// result = float64(temp)
					services.EnqueueRequestToAutoScalingSqs("add")
					count, _ := services.GetCountOfMessagesInQueue("add")
					if count > 4 {
						// fmt.Println("count is greater than 0")
						// fmt.Println("count is ", count)
						go services.UpdatePodsInDeployment("add", int32(count+4))
					}
					// temp, _ = grpcservice.CallGrpcAdd(30001, float32(leftOperand), float32(rightOperand))

					result = leftOperand + rightOperand

				case "-":

					services.EnqueueRequestToAutoScalingSqs("subtract")
					count, _ := services.GetCountOfMessagesInQueue("subtract")
					if count > 4 {
						// fmt.Println("count is greater than 0")
						// fmt.Println("count is ", count)
						go services.UpdatePodsInDeployment("subtract", int32(count+4))
					}
					temp, _ = grpcservice.CallGrpcSubtract(30002, float32(leftOperand), float32(rightOperand))
					result = float64(temp)

				case "*":

					services.EnqueueRequestToAutoScalingSqs("multiply")
					count, _ := services.GetCountOfMessagesInQueue("multiply")
					if count > 4 {
						go services.UpdatePodsInDeployment("multiply", int32(count+4))
					}
					temp, _ = grpcservice.CallGrpcMultiply(30003, float32(leftOperand), float32(rightOperand))
					result = float64(temp)

				case "/":
					services.EnqueueRequestToAutoScalingSqs("divide")
					count, _ := services.GetCountOfMessagesInQueue("divide")
					if count > 4 {
						go services.UpdatePodsInDeployment("divide", int32(count+4))
					}
					temp, _ = grpcservice.CallGrpcDivision(30004, float32(leftOperand), float32(rightOperand))
					result = float64(temp)

				default:
					err = errors.New("unsupported operator")
				}

				if err != nil {
					errorsChan <- err
				} else {
					operations <- result
				}
			}()

			select {
			case result := <-operations:
				stack = append(stack, result)
			case err := <-errorsChan:
				return 0, err
			}
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}
