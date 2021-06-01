// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: consignment.proto

package consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Shipping service

func NewShippingEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Shipping service

type ShippingService interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// 创建一个新的方法
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingService struct {
	c    client.Client
	name string
}

func NewShippingService(name string, c client.Client) ShippingService {
	return &shippingService{
		c:    c,
		name: name,
	}
}

func (c *shippingService) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Shipping.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Shipping.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shipping service

type ShippingHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	// 创建一个新的方法
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingHandler(s server.Server, hdlr ShippingHandler, opts ...server.HandlerOption) error {
	type shipping interface {
		CreateConsignment(ctx context.Context, in *Consignment, out *Response) error
		GetConsignments(ctx context.Context, in *GetRequest, out *Response) error
	}
	type Shipping struct {
		shipping
	}
	h := &shippingHandler{hdlr}
	return s.Handle(s.NewHandler(&Shipping{h}, opts...))
}

type shippingHandler struct {
	ShippingHandler
}

func (h *shippingHandler) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingHandler.CreateConsignment(ctx, in, out)
}

func (h *shippingHandler) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingHandler.GetConsignments(ctx, in, out)
}
