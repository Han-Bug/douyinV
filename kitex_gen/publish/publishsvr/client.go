// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishsvr

import (
	"context"
	publish "douyinV/kitex_gen/publish"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Action(ctx context.Context, req *publish.ActionRequest, callOptions ...callopt.Option) (r *publish.ActionResponse, err error)
	List(ctx context.Context, req *publish.ListRequest, callOptions ...callopt.Option) (r *publish.ListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPublishSvrClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPublishSvrClient struct {
	*kClient
}

func (p *kPublishSvrClient) Action(ctx context.Context, req *publish.ActionRequest, callOptions ...callopt.Option) (r *publish.ActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Action(ctx, req)
}

func (p *kPublishSvrClient) List(ctx context.Context, req *publish.ListRequest, callOptions ...callopt.Option) (r *publish.ListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.List(ctx, req)
}
