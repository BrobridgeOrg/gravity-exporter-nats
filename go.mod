module github.com/BrobridgeOrg/gravity-exporter-nats

go 1.13

require (
	github.com/BrobridgeOrg/gravity-api v0.2.14
	github.com/BrobridgeOrg/gravity-sdk v0.0.8
	github.com/BrobridgeOrg/gravity-transmitter-postgres v0.0.0-20210521202000-7255c776322a
	github.com/jinzhu/copier v0.3.0
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/prometheus/common v0.15.0
	github.com/sirupsen/logrus v1.7.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/viper v1.7.1
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	google.golang.org/grpc v1.32.0
	google.golang.org/grpc/examples v0.0.0-20200807164945-d3e3e7a46f57 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

//replace github.com/BrobridgeOrg/gravity-api v0.0.0-20200808075207-3326e6e4eea1 => ../gravity-api
