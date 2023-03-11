package listen

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"upi/app/mqueue/cmd/mq/internal/config"
	"upi/app/mqueue/cmd/mq/internal/handle"
	"upi/app/mqueue/cmd/mq/internal/svc"
)

func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	services = append(services, KqMqs(ctx, c, svcContext)...)

	return services
}

func KqMqs(ctx context.Context, c config.Config, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.PaymentUpdateStateConf, handle.NewPaymentUpdateStatusMq(ctx, svcContext)),
	}
}