package rabbitmq

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// NewMessageLog ...
func NewMessageLog(queueName string) {
	fmt.Println(aurora.BgBlue("*** New message from queue: " + queueName))
}

// PublishErrorLog ...
func PublishErrorLog(queueName string, err error) {
	fmt.Println(aurora.Red(
		fmt.Sprintf(
			"*** Error when publish message to %s: %s",
			queueName,
			err.Error(),
			),
		),
	)
}
