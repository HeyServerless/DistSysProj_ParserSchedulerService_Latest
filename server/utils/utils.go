package utils

import (
	"strings"

	"github.com/google/uuid"
)

func Int32Ptr(i int32) *int32 { return &i }

var precedence = map[string]int{
	"(": 0,
	")": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

type AutoScalingQueue struct {
	QueueUrl  string
	QueueName string
}

var AutoScalingQueues []AutoScalingQueue = []AutoScalingQueue{
	{
		QueueUrl:  "https://sqs.us-east-1.amazonaws.com/626995068279/AddAutoscalingQueue",
		QueueName: "add",
	},
	{
		QueueUrl:  "https://sqs.us-east-1.amazonaws.com/626995068279/MultiplyAutoscalingQueue",
		QueueName: "multiply",
	},
	{
		QueueUrl:  "https://sqs.us-east-1.amazonaws.com/626995068279/SubtractAutoscalingQueue",
		QueueName: "subtract",
	},
	{
		QueueUrl:  "https://sqs.us-east-1.amazonaws.com/626995068279/DivisionAutoscalingQueue",
		QueueName: "division",
	},
}

func GetQueueUrlByName(functionality string) AutoScalingQueue {

	for _, queue := range AutoScalingQueues {
		if strings.Contains(queue.QueueName, functionality) {
			return queue
		}
	}
	return AutoScalingQueue{}

}

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
