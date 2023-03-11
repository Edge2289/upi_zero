package handle

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"upi/app/mqueue/cmd/job/internal/svc"
)

type Test struct {
	svcCtx *svc.ServiceContext
}

func NewTest(svcCtx *svc.ServiceContext) *Test {
	return &Test{
		svcCtx: svcCtx,
	}
}

func (t *Test) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	fmt.Println("scheduler job test demo ----> entity exec \n ")
	return nil
}