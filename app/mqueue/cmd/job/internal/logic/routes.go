package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"upi/app/mqueue/cmd/job/internal/handle"
	"upi/app/mqueue/cmd/job/internal/svc"
	"upi/app/mqueue/cmd/jobType"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//scheduler job
	mux.Handle(jobType.SchedulerSettleTest, handle.NewTest(l.svcCtx))
	return mux
}

