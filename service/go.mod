module go-grpc/service

go 1.18

require (
	go-grpc/commons v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220706132729-d86698d07c53 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace go-grpc/commons => ../commons
