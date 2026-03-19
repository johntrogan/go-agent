module github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus

go 1.24.0

require (
	github.com/newrelic/go-agent/v3 v3.42.0
	github.com/sirupsen/logrus v1.8.3
)

require (
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/grpc v1.79.3 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/newrelic/go-agent/v3 => ../../..
