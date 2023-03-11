package mq

import (
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"testing"
)

func TestMqClient(t *testing.T)  {
	kqueuePaymentUpdatePayStatusClient := kq.NewPusher([]string{"0.0.0.0:9092"}, "payment-update-paystatus-topic")
	err := kqueuePaymentUpdatePayStatusClient.Push("1231231231231")
	if err != nil {
		fmt.Println("err", err.Error())
	}
	fmt.Println("-----")
}