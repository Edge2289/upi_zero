package trade

import (
	"context"
	"upi/app/unionpay/com/rpc/pb"
	"upi/common/turn"
	"upi/common/xerr"

	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	PayOrderNoOnly = "pay:trade:order:"
)

// Pay 支付请求
func (l *PayLogic) Pay(req *types.PayReq) (resp *types.PayResp, err error) {

	logx.WithContext(l.ctx).Info("支付请求数据:" + turn.StructToJson(req))
	cOrder, err := l.svcCtx.Cache.Get(PayOrderNoOnly + req.TradeNo)
	if err != nil {
		return nil, err
	}
	if cOrder != "" {
		return nil, xerr.NewErrMsg("请勿重复下单")
	}
	_ = l.svcCtx.Cache.Set(PayOrderNoOnly + req.TradeNo, "1")
	_ = l.svcCtx.Cache.Expire(PayOrderNoOnly + req.TradeNo, 60 * 3)

	_, err = l.svcCtx.UnionPay.SavePayOrder(l.ctx, &pb.SavePayOrderReq{
		TradeNo:       req.TradeNo,
		OaOrderAmount: req.Amount,
		MallOrderCode: req.TradeNo,
		MallAmount:    req.Amount,
		MallType:      req.PayMode,
	})

	if err != nil {
		_, _ = l.svcCtx.Cache.Del(PayOrderNoOnly + req.TradeNo, "1")
		return nil, err
	}

	_, _ = l.svcCtx.Cache.Del(PayOrderNoOnly + req.TradeNo, "1")
	return &types.PayResp{}, nil
}
