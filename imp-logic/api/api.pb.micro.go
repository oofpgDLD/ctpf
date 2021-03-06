// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/oofpgDLD/ctpf/imp-comet/api"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Logic service

type LogicService interface {
	// Ping Service
	Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingReply, error)
	// Close Service
	Close(ctx context.Context, in *CloseReq, opts ...client.CallOption) (*CloseReply, error)
	// Connect
	Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectReply, error)
	// Disconnect
	Disconnect(ctx context.Context, in *DisconnectReq, opts ...client.CallOption) (*DisconnectReply, error)
	// Heartbeat
	Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...client.CallOption) (*HeartbeatReply, error)
	// RenewOnline
	RenewOnline(ctx context.Context, in *OnlineReq, opts ...client.CallOption) (*OnlineReply, error)
	// Receive
	Receive(ctx context.Context, in *ReceiveReq, opts ...client.CallOption) (*ReceiveReply, error)
	//ServerList
	Nodes(ctx context.Context, in *NodesReq, opts ...client.CallOption) (*NodesReply, error)
}

type logicService struct {
	c    client.Client
	name string
}

func NewLogicService(name string, c client.Client) LogicService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ctpf.logic"
	}
	return &logicService{
		c:    c,
		name: name,
	}
}

func (c *logicService) Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Ping", in)
	out := new(PingReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Close(ctx context.Context, in *CloseReq, opts ...client.CallOption) (*CloseReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Close", in)
	out := new(CloseReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Connect", in)
	out := new(ConnectReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Disconnect(ctx context.Context, in *DisconnectReq, opts ...client.CallOption) (*DisconnectReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Disconnect", in)
	out := new(DisconnectReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...client.CallOption) (*HeartbeatReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Heartbeat", in)
	out := new(HeartbeatReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) RenewOnline(ctx context.Context, in *OnlineReq, opts ...client.CallOption) (*OnlineReply, error) {
	req := c.c.NewRequest(c.name, "Logic.RenewOnline", in)
	out := new(OnlineReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Receive(ctx context.Context, in *ReceiveReq, opts ...client.CallOption) (*ReceiveReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Receive", in)
	out := new(ReceiveReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicService) Nodes(ctx context.Context, in *NodesReq, opts ...client.CallOption) (*NodesReply, error) {
	req := c.c.NewRequest(c.name, "Logic.Nodes", in)
	out := new(NodesReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logic service

type LogicHandler interface {
	// Ping Service
	Ping(context.Context, *PingReq, *PingReply) error
	// Close Service
	Close(context.Context, *CloseReq, *CloseReply) error
	// Connect
	Connect(context.Context, *ConnectReq, *ConnectReply) error
	// Disconnect
	Disconnect(context.Context, *DisconnectReq, *DisconnectReply) error
	// Heartbeat
	Heartbeat(context.Context, *HeartbeatReq, *HeartbeatReply) error
	// RenewOnline
	RenewOnline(context.Context, *OnlineReq, *OnlineReply) error
	// Receive
	Receive(context.Context, *ReceiveReq, *ReceiveReply) error
	//ServerList
	Nodes(context.Context, *NodesReq, *NodesReply) error
}

func RegisterLogicHandler(s server.Server, hdlr LogicHandler, opts ...server.HandlerOption) error {
	type logic interface {
		Ping(ctx context.Context, in *PingReq, out *PingReply) error
		Close(ctx context.Context, in *CloseReq, out *CloseReply) error
		Connect(ctx context.Context, in *ConnectReq, out *ConnectReply) error
		Disconnect(ctx context.Context, in *DisconnectReq, out *DisconnectReply) error
		Heartbeat(ctx context.Context, in *HeartbeatReq, out *HeartbeatReply) error
		RenewOnline(ctx context.Context, in *OnlineReq, out *OnlineReply) error
		Receive(ctx context.Context, in *ReceiveReq, out *ReceiveReply) error
		Nodes(ctx context.Context, in *NodesReq, out *NodesReply) error
	}
	type Logic struct {
		logic
	}
	h := &logicHandler{hdlr}
	return s.Handle(s.NewHandler(&Logic{h}, opts...))
}

type logicHandler struct {
	LogicHandler
}

func (h *logicHandler) Ping(ctx context.Context, in *PingReq, out *PingReply) error {
	return h.LogicHandler.Ping(ctx, in, out)
}

func (h *logicHandler) Close(ctx context.Context, in *CloseReq, out *CloseReply) error {
	return h.LogicHandler.Close(ctx, in, out)
}

func (h *logicHandler) Connect(ctx context.Context, in *ConnectReq, out *ConnectReply) error {
	return h.LogicHandler.Connect(ctx, in, out)
}

func (h *logicHandler) Disconnect(ctx context.Context, in *DisconnectReq, out *DisconnectReply) error {
	return h.LogicHandler.Disconnect(ctx, in, out)
}

func (h *logicHandler) Heartbeat(ctx context.Context, in *HeartbeatReq, out *HeartbeatReply) error {
	return h.LogicHandler.Heartbeat(ctx, in, out)
}

func (h *logicHandler) RenewOnline(ctx context.Context, in *OnlineReq, out *OnlineReply) error {
	return h.LogicHandler.RenewOnline(ctx, in, out)
}

func (h *logicHandler) Receive(ctx context.Context, in *ReceiveReq, out *ReceiveReply) error {
	return h.LogicHandler.Receive(ctx, in, out)
}

func (h *logicHandler) Nodes(ctx context.Context, in *NodesReq, out *NodesReply) error {
	return h.LogicHandler.Nodes(ctx, in, out)
}
