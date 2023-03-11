package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"upi/app/business/com/api/internal/config"
	"upi/app/business/com/rpc/splitbillssummary"
)

type ServiceContext struct {
	Config config.Config
	SplitBillsSummaryRpc splitbillssummary.SplitBillsSummary
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		SplitBillsSummaryRpc: splitbillssummary.NewSplitBillsSummary(zrpc.MustNewClient(c.SplitBillsSummaryRpcConfig)),
	}
}
