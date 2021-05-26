package instance

import (
	"runtime"
	"time"

	eventbus "github.com/BrobridgeOrg/gravity-exporter-nats/pkg/eventbus/service"
	subscriber "github.com/BrobridgeOrg/gravity-exporter-nats/pkg/subscriber/service"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppInstance struct {
	done       chan bool
	eventBus   *eventbus.EventBus
	subscriber *subscriber.Subscriber
}

func NewAppInstance() *AppInstance {

	a := &AppInstance{
		done: make(chan bool),
	}

	return a
}

func (a *AppInstance) Init() error {

	log.WithFields(log.Fields{
		"max_procs": runtime.GOMAXPROCS(0),
	}).Info("Starting application")

	// Initializing modules
	a.eventBus = eventbus.NewEventBus(
		a,
		viper.GetString("nats.host"),
		eventbus.EventBusHandler{
			Reconnect: func(natsConn *nats.Conn) {
				log.Warn("re-connected to event server")
				//a.eventBus.InitSubscription()
			},
			Disconnect: func(natsConn *nats.Conn) {
				log.Error("event server was disconnected")
			},
		},
		eventbus.Options{
			PingInterval:        time.Duration(viper.GetInt64("nats.pingInterval")),
			MaxPingsOutstanding: viper.GetInt("nats.maxPingsOutstanding"),
			MaxReconnects:       viper.GetInt("nats.maxReconnects"),
		},
	)
	a.subscriber = subscriber.NewSubscriber(a)

	// Initializing EventBus
	err := a.initEventBus()
	if err != nil {
		return err
	}

	err = a.subscriber.Init()
	if err != nil {
		return err
	}

	return nil
}

func (a *AppInstance) Uninit() {
}

func (a *AppInstance) Run() error {

	err := a.subscriber.Run()
	if err != nil {
		return err
	}

	<-a.done

	return nil
}
