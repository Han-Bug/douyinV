package tracer

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// InitJaeger Jaeger是一个分布式跟踪系统
// 这边是初始化Jaeger的函数
func InitJaeger(service string) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return
}
