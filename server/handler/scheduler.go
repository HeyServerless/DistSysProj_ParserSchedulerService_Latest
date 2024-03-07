package handler

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	// "encoding/json"
	"github.com/Knetic/govaluate"

	"net/http"

	"github.com/gin-gonic/gin"
	// appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/util/intstr"
	// "k8s.io/client-go/util/retry"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ScheduleAndCalculate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  "token",
	})
}

// Expression represents a mathematical expression to be evaluated
type Expression struct {
	Expression string `json:"expression"`
}

// Result represents the result of a mathematical expression
type Result struct {
	Expression string  `json:"expression"`
	Value      float64 `json:"value"`
}

func RunTheSceduler(c *gin.Context) {
	// Load the Kubernetes config

	// Create a new cron job that runs every minute

	// c.Start()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  "token",
	})
}

// divideExpression divides a given expression into multiple sub-expressions
func divideExpression(expression string) []string {
	// Divide the expression into sub-expressions
	subExpressions := make([]string, 0)
	exprLen := len(expression)
	startIndex := 0
	for i := 0; i < exprLen; i++ {
		if i == exprLen-1 {
			subExpressions = append(subExpressions, expression[startIndex:])
			break
		}
		// Check if the current character is an operator
		isOperator := false
		switch expression[i] {
		case '+', '-', '*', '/':
			isOperator = true
		}

		// If the current character is an operator, add the previous part as a sub-expression
		if isOperator {
			subExpressions = append(subExpressions, expression[startIndex:i])
			startIndex = i + 1
		}
	}

	return subExpressions
}

// sendMessage sends a message to the SQS queue
func sendMessage(svc *sqs.SQS, message string) error {
	queueURL := "https://sqs.us-east-1.amazonaws.com/626995068279/outputqueue"
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	}

	_, err := svc.SendMessage(params)
	if err != nil {
		return err
	}

	return nil
}

// receiveMessage receives a message from the SQS queue
func receiveMessage(svc *sqs.SQS) (*Result, error) {
	queueURL := "https://sqs.us-east-1.amazonaws.com/626995068279/outputqueue"
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(0),
	}

	resp, err := svc.ReceiveMessage(params)
	if err != nil {
		return nil, err
	}

	if len(resp.Messages) == 0 {
		return nil, nil
	}

	// Parse the message as an Expression
	var expression Expression
	err = json.Unmarshal([]byte(*resp.Messages[0].Body), &expression)
	if err != nil {
		return nil, err
	}

	// Evaluate the Expression and create a Result
	value, err := evaluateExpressionSCH(expression.Expression)
	if err != nil {
		return nil, err
	}
	result := &Result{
		Expression: expression.Expression,
		Value:      value,
	}

	// Delete the message from the queue
	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// evaluateExpressionSCH evaluates a mathematical expression and returns the result
func evaluateExpressionSCH(expression string) (float64, error) {
	// evaluate the expression
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}
	returnValue, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return returnValue.(float64), nil

	// return value, nil
}

// aggregateResults aggregates the results of multiple sub-expressions
func aggregateResults(results []*Result) *Result {
	// Sum the values of all results
	sum := 0.0
	for _, result := range results {
		sum += result.Value
	}
	// Create a new Result with the final value
	finalResult := &Result{
		Expression: "Total",
		Value:      sum,
	}

	return finalResult
}
