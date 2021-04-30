// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/task/task.proto

package micro_service_task

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

// Api Endpoints for TaskService service

func NewTaskServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TaskService service

type TaskService interface {
	Create(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error)
	Delete(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error)
	Modify(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error)
	Finished(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) Create(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Create", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Delete(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Delete", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Modify(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Modify", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Finished(ctx context.Context, in *Task, opts ...client.CallOption) (*EditResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Finished", in)
	out := new(EditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	Create(context.Context, *Task, *EditResponse) error
	Delete(context.Context, *Task, *EditResponse) error
	Modify(context.Context, *Task, *EditResponse) error
	Finished(context.Context, *Task, *EditResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		Create(ctx context.Context, in *Task, out *EditResponse) error
		Delete(ctx context.Context, in *Task, out *EditResponse) error
		Modify(ctx context.Context, in *Task, out *EditResponse) error
		Finished(ctx context.Context, in *Task, out *EditResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) Create(ctx context.Context, in *Task, out *EditResponse) error {
	return h.TaskServiceHandler.Create(ctx, in, out)
}

func (h *taskServiceHandler) Delete(ctx context.Context, in *Task, out *EditResponse) error {
	return h.TaskServiceHandler.Delete(ctx, in, out)
}

func (h *taskServiceHandler) Modify(ctx context.Context, in *Task, out *EditResponse) error {
	return h.TaskServiceHandler.Modify(ctx, in, out)
}

func (h *taskServiceHandler) Finished(ctx context.Context, in *Task, out *EditResponse) error {
	return h.TaskServiceHandler.Finished(ctx, in, out)
}

func (h *taskServiceHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.TaskServiceHandler.Search(ctx, in, out)
}
