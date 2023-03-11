package trade

import (
	"context"
	"fmt"
	"testing"
	"upi/expand/unionpay"
	"upi/expand/unionpay/core"
)

func TestUnifiedPay(t *testing.T) {
	c := context.Background()

	client := core.NewClient()

	w := WebPay {
		Client: client,
		Config: unionpay.Config{
			Host: "",
			Account: "",
			SysCode: "",
			PublicKey: "",
			PrivateKey: "",
		},
	}
	pay, c2, err := w.UnifiedPay(c, UnifiedPayRequest{

	})
	fmt.Println(pay, c2, err)
}

func TestName(t *testing.T) {
	go func() {
		for k := range [10]int{1, 2, 3, 4, 5, 6, 7 ,8, 9, 10} {
			fmt.Println("k", k)
		}
	}()

	go func() {
		for k := range [10]int{11, 12, 13, 14, 15, 16, 17 , 18, 19, 20} {
			fmt.Println("c", k)
		}
	}()
}