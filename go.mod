module github.com/BrobridgeOrg/gravity-exporter-nats

go 1.13

require (
	github.com/BrobridgeOrg/gravity-api v0.0.0-20200808191818-646e409ed0b8
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/prometheus/common v0.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20200807164945-d3e3e7a46f57 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

//replace github.com/BrobridgeOrg/gravity-api v0.0.0-20200808075207-3326e6e4eea1 => ../gravity-api
