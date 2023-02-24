// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishsvr

import (
	"context"
	publish "douyinV/kitex_gen/publish"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishSvrServiceInfo
}

var publishSvrServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishSvr"
	handlerType := (*publish.PublishSvr)(nil)
	methods := map[string]kitex.MethodInfo{
		"Action": kitex.NewMethodInfo(actionHandler, newPublishSvrActionArgs, newPublishSvrActionResult, false),
		"List":   kitex.NewMethodInfo(listHandler, newPublishSvrListArgs, newPublishSvrListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "publish",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishSvrActionArgs)
	realResult := result.(*publish.PublishSvrActionResult)
	success, err := handler.(publish.PublishSvr).Action(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishSvrActionArgs() interface{} {
	return publish.NewPublishSvrActionArgs()
}

func newPublishSvrActionResult() interface{} {
	return publish.NewPublishSvrActionResult()
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishSvrListArgs)
	realResult := result.(*publish.PublishSvrListResult)
	success, err := handler.(publish.PublishSvr).List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishSvrListArgs() interface{} {
	return publish.NewPublishSvrListArgs()
}

func newPublishSvrListResult() interface{} {
	return publish.NewPublishSvrListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, req *publish.ActionRequest) (r *publish.ActionResponse, err error) {
	var _args publish.PublishSvrActionArgs
	_args.Req = req
	var _result publish.PublishSvrActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, req *publish.ListRequest) (r *publish.ListResponse, err error) {
	var _args publish.PublishSvrListArgs
	_args.Req = req
	var _result publish.PublishSvrListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}