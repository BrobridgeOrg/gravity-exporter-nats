package app

import (
	"github.com/BrobridgeOrg/gravity-exporter-nats/pkg/eventbus"
)

type App interface {
	GetEventBus() eventbus.EventBus
}
