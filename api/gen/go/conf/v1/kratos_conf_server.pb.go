// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: conf/v1/kratos_conf_server.proto

package confv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 服务器
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http     *Server_HTTP     `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`          // HTTP服务
	Grpc     *Server_GRPC     `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`          // gRPC服务
	Rabbitmq *Server_RabbitMQ `protobuf:"bytes,10,opt,name=rabbitmq,proto3" json:"rabbitmq,omitempty"` // RabbitMQ服务
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Server) GetRabbitmq() *Server_RabbitMQ {
	if x != nil {
		return x.Rabbitmq
	}
	return nil
}

// REST
type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`                                   // 网络
	Addr          string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                                         // 服务监听地址
	Timeout       *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`                                   // 超时时间
	Cors          *Server_HTTP_CORS    `protobuf:"bytes,4,opt,name=cors,proto3" json:"cors,omitempty"`                                         // 跨域配置
	Middleware    *Middleware          `protobuf:"bytes,5,opt,name=middleware,proto3" json:"middleware,omitempty"`                             // 中间件
	EnableSwagger bool                 `protobuf:"varint,6,opt,name=enable_swagger,json=enableSwagger,proto3" json:"enable_swagger,omitempty"` // 启用SwaggerUI
	EnablePprof   bool                 `protobuf:"varint,7,opt,name=enable_pprof,json=enablePprof,proto3" json:"enable_pprof,omitempty"`       // 启用pprof
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetCors() *Server_HTTP_CORS {
	if x != nil {
		return x.Cors
	}
	return nil
}

func (x *Server_HTTP) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Server_HTTP) GetEnableSwagger() bool {
	if x != nil {
		return x.EnableSwagger
	}
	return false
}

func (x *Server_HTTP) GetEnablePprof() bool {
	if x != nil {
		return x.EnablePprof
	}
	return false
}

// gPRC
type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network    string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"` // 网络
	Addr       string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`       // 服务监听地址
	Timeout    *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"` // 超时时间
	Middleware *Middleware          `protobuf:"bytes,4,opt,name=middleware,proto3" json:"middleware,omitempty"`
	Tls        *Server_TLS          `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty"` // TLS配置
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_GRPC) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Server_GRPC) GetTls() *Server_TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

type Server_TLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	CertFile string `protobuf:"bytes,2,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"` // 证书文件
	KeyFile  string `protobuf:"bytes,3,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`    // 私钥文件
	CaFile   string `protobuf:"bytes,4,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty"`       // CA证书文件
}

func (x *Server_TLS) Reset() {
	*x = Server_TLS{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_TLS) ProtoMessage() {}

func (x *Server_TLS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_TLS.ProtoReflect.Descriptor instead.
func (*Server_TLS) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Server_TLS) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Server_TLS) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *Server_TLS) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *Server_TLS) GetCaFile() string {
	if x != nil {
		return x.CaFile
	}
	return ""
}

// RabbitMQ
type Server_RabbitMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"` // 对端网络地址
}

func (x *Server_RabbitMQ) Reset() {
	*x = Server_RabbitMQ{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_RabbitMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_RabbitMQ) ProtoMessage() {}

func (x *Server_RabbitMQ) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_RabbitMQ.ProtoReflect.Descriptor instead.
func (*Server_RabbitMQ) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Server_RabbitMQ) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type Server_HTTP_CORS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
	Origins []string `protobuf:"bytes,3,rep,name=origins,proto3" json:"origins,omitempty"`
}

func (x *Server_HTTP_CORS) Reset() {
	*x = Server_HTTP_CORS{}
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP_CORS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP_CORS) ProtoMessage() {}

func (x *Server_HTTP_CORS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP_CORS.ProtoReflect.Descriptor instead.
func (*Server_HTTP_CORS) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Server_HTTP_CORS) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Server_HTTP_CORS) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Server_HTTP_CORS) GetOrigins() []string {
	if x != nil {
		return x.Origins
	}
	return nil
}

var File_conf_v1_kratos_conf_server_proto protoreflect.FileDescriptor

var file_conf_v1_kratos_conf_server_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe4, 0x06, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04,
	0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x28, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x12, 0x34, 0x0a, 0x08, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x4d, 0x51, 0x52, 0x08, 0x72, 0x61,
	0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x1a, 0xed, 0x02, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x2e, 0x43, 0x4f, 0x52, 0x53, 0x52, 0x04, 0x63, 0x6f, 0x72,
	0x73, 0x12, 0x33, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x70, 0x72, 0x6f, 0x66,
	0x1a, 0x54, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0xc5, 0x01, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x1a, 0x6e,
	0x0a, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x28,
	0x0a, 0x08, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x4d, 0x51, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x9c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x65, 0x63, 0x34, 0x30, 0x34, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43,
	0x6f, 0x6e, 0x66, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43,
	0x6f, 0x6e, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_kratos_conf_server_proto_rawDescOnce sync.Once
	file_conf_v1_kratos_conf_server_proto_rawDescData = file_conf_v1_kratos_conf_server_proto_rawDesc
)

func file_conf_v1_kratos_conf_server_proto_rawDescGZIP() []byte {
	file_conf_v1_kratos_conf_server_proto_rawDescOnce.Do(func() {
		file_conf_v1_kratos_conf_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_kratos_conf_server_proto_rawDescData)
	})
	return file_conf_v1_kratos_conf_server_proto_rawDescData
}

var file_conf_v1_kratos_conf_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_conf_v1_kratos_conf_server_proto_goTypes = []any{
	(*Server)(nil),              // 0: conf.v1.Server
	(*Server_HTTP)(nil),         // 1: conf.v1.Server.HTTP
	(*Server_GRPC)(nil),         // 2: conf.v1.Server.GRPC
	(*Server_TLS)(nil),          // 3: conf.v1.Server.TLS
	(*Server_RabbitMQ)(nil),     // 4: conf.v1.Server.RabbitMQ
	(*Server_HTTP_CORS)(nil),    // 5: conf.v1.Server.HTTP.CORS
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
	(*Middleware)(nil),          // 7: conf.v1.Middleware
}
var file_conf_v1_kratos_conf_server_proto_depIdxs = []int32{
	1, // 0: conf.v1.Server.http:type_name -> conf.v1.Server.HTTP
	2, // 1: conf.v1.Server.grpc:type_name -> conf.v1.Server.GRPC
	4, // 2: conf.v1.Server.rabbitmq:type_name -> conf.v1.Server.RabbitMQ
	6, // 3: conf.v1.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	5, // 4: conf.v1.Server.HTTP.cors:type_name -> conf.v1.Server.HTTP.CORS
	7, // 5: conf.v1.Server.HTTP.middleware:type_name -> conf.v1.Middleware
	6, // 6: conf.v1.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	7, // 7: conf.v1.Server.GRPC.middleware:type_name -> conf.v1.Middleware
	3, // 8: conf.v1.Server.GRPC.tls:type_name -> conf.v1.Server.TLS
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_conf_v1_kratos_conf_server_proto_init() }
func file_conf_v1_kratos_conf_server_proto_init() {
	if File_conf_v1_kratos_conf_server_proto != nil {
		return
	}
	file_conf_v1_kratos_conf_middleware_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_v1_kratos_conf_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_kratos_conf_server_proto_goTypes,
		DependencyIndexes: file_conf_v1_kratos_conf_server_proto_depIdxs,
		MessageInfos:      file_conf_v1_kratos_conf_server_proto_msgTypes,
	}.Build()
	File_conf_v1_kratos_conf_server_proto = out.File
	file_conf_v1_kratos_conf_server_proto_rawDesc = nil
	file_conf_v1_kratos_conf_server_proto_goTypes = nil
	file_conf_v1_kratos_conf_server_proto_depIdxs = nil
}
