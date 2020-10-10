package exporter

import (
	"sync/atomic"

	"golang.org/x/net/context"

	pb "github.com/BrobridgeOrg/gravity-api/service/exporter"
	app "github.com/BrobridgeOrg/gravity-exporter-nats/pkg/app"
	"github.com/prometheus/common/log"
)

var counter uint64 = 0
var SendEventSuccess = pb.SendEventReply{
	Success: true,
}

type Service struct {
	app app.App
}

func NewService(a app.App) *Service {

	service := &Service{
		app: a,
	}
	return service
}

func (service *Service) SendEvent(ctx context.Context, in *pb.SendEventRequest) (*pb.SendEventReply, error) {

	id := atomic.AddUint64((*uint64)(&counter), 1)

	if id%1000 == 0 {
		log.Info(id)
	}

	conn := service.app.GetEventBus().GetConnection()

	err := conn.Publish(in.Channel, in.Payload)
	if err != nil {
		return &pb.SendEventReply{
			Success: false,
		}, err
	}

	return &SendEventSuccess, nil
}
