package api

//go:generate protoc -I schema/ -I ../vendor -I ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis schema/profile.proto --go_out=plugins=grpc:.
//go:generate protoc -I schema/ -I ../vendor -I ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis schema/profile.proto --grpc-gateway_out=logtostderr=true:.
//go:generate protoc -I schema/ -I ../vendor -I ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis schema/profile.proto --swagger_out=logtostderr=true:.
