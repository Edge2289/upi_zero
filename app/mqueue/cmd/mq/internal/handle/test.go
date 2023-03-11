package handle

import (
	"context"
	"fmt"
	"upi/app/mqueue/cmd/mq/internal/svc"
	"upi/app/unionpay/com/rpc/pb"
)

type PaymentUpdateStatus struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateStatus {
	return &PaymentUpdateStatus{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentUpdateStatus) Consume(_, val string) error {
	fmt.Println("----consume ---" , val)
	//time.Sleep(time.Second * 10)
	resp, _ := l.svcCtx.Unionpay.SavePayOrder(l.ctx, &pb.SavePayOrderReq{TradeNo: val})
	fmt.Println("resp", resp.Id)
	//logx.Infof("123123131231")
	return nil //errors.New("xxxxx")
}

