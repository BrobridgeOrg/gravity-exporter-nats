module github.com/BrobridgeOrg/gravity-exporter-nats

go 1.13

require (
	github.com/BrobridgeOrg/gravity-sdk v0.0.40
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad // indirect
	google.golang.org/grpc v1.32.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

//replace github.com/BrobridgeOrg/gravity-sdk => ../gravity-sdk
