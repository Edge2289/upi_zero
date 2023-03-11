// Code generated by goctl. DO NOT EDIT!
// Source: splitbillssummary.proto

package server

import (
	"context"

	"upi/app/business/com/rpc/internal/logic"
	"upi/app/business/com/rpc/internal/svc"
	"upi/app/business/com/rpc/pd"
)

type SplitBillsSummaryServer struct {
	svcCtx *svc.ServiceContext
	pd.UnimplementedSplitBillsSummaryServer
}

func NewSplitBillsSummaryServer(svcCtx *svc.ServiceContext) *SplitBillsSummaryServer {
	return &SplitBillsSummaryServer{
		svcCtx: svcCtx,
	}
}

func (s *SplitBillsSummaryServer) List(ctx context.Context, in *pd.SplitBillsSummaryListReq) (*pd.SplitBillsSummaryListResp, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}