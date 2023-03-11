package main

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"strconv"
	"time"
)

type cc struct {
	Brokers []string
	Topic   string
}

func main() {
	var c cc
	conf.MustLoad("test.yaml", &c)
	//
	//q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {
	//	fmt.Printf("=> %s\n", v)
	//	return nil
	//}))
	//defer q.Stop()
	//q.Start()

	//fmt.Println("c", c)
	//kqueuePaymentUpdatePayStatusClient := kq.NewPusher(c.Brokers, c.Topic)
	//err := kqueuePaymentUpdatePayStatusClient.Push("小星星")
	//fmt.Println("err", err)
	//if err != nil {
	//	fmt.Println("err", err.Error())
	//}
	//fmt.Println("-----")
	type message struct {
		Keyq     string `json:"Keyq"`
		Value   string `json:"value"`
		//Payload string `json:"message"`
	}


	pusher := kq.NewPusher([]string{
		"kafka:9092",
		"kafka:9092",
		"kafka:9092",
	}, "payment-update-paystatus-topic")
	//m := message{
	//	Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
	//	Value:   fmt.Sprintf("%d,%d", 1, 100),
	//	Payload: fmt.Sprintf("%d,%d", 2, 100),
	//}
	//body, _ := json.Marshal(m)
	//_ = pusher.Push(string(body))
	//time.Sleep(time.Second * 10)

	//ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 1; round++ {
		//select {
		//case <-ticker.C:
			count := 201
			m := message{
				Keyq:     strconv.FormatInt(time.Now().UnixNano(), 10),
				Value:   fmt.Sprintf("%d,%d", round, count),
				//Payload: fmt.Sprintf("%d,%d", 121, count),
			}
			body, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body))
			//_ = pusher.Push(string(body))
			if err := pusher.Push(string(body)); err != nil {
				log.Fatal(err)
			}
		//}
	}
	cmdline.EnterToContinue()
}
