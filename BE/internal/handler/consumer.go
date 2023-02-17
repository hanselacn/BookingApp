// Package handler
package handler

import (
	"context"

	"BookingApp/BE/internal/appctx"
	"BookingApp/BE/internal/consts"
	uContract "BookingApp/BE/internal/ucase/contract"
	"BookingApp/BE/pkg/awssqs"
)

// SQSConsumerHandler sqs consumer message processor handler
func SQSConsumerHandler(msgHandler uContract.MessageProcessor) awssqs.MessageProcessorFunc {
	return func(decoder *awssqs.MessageDecoder) error {
		return msgHandler.Serve(context.Background(), &appctx.ConsumerData{
			Body:        []byte(*decoder.Body),
			Key:         []byte(*decoder.MessageId),
			ServiceType: consts.ServiceTypeConsumer,
		})
	}
}
