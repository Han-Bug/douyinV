package main

import (
	"context"
	"douyin_mini/cmd/api/handlers"

	"douyin_mini/cmd/api/rpc"
	"douyin_mini/kitex_gen/user"
	"douyin_mini/pkg/constants"
	"douyin_mini/pkg/errno"
	"douyin_mini/pkg/tracer"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/hertz-contrib/jwt"
	"time"
)

// Init api服务的初始化
func Init() {
	//初始化分布式跟踪系统
	tracer.InitJaeger(constants.ApiServiceName)
	//初始化rpc客户端
	rpc.InitRPC()
}

// http层的服务入口！！！！
// api服务的主启动函数
func main() {
	//进行api层的服务初始化
	Init()
	//利用hertz创建一个web服务，接收http请求
	r := server.New(
		//规定好ip-port
		server.WithHostPorts("127.0.0.1:8080"),
		//如果开启，当当前路径不能被匹配上时
		//server 会去检查其他方法是否注册了当前路径的路由
		//如果存在则会响应"Method Not Allowed"
		//并返回状态码405;
		//如果没有，
		//则会用 NotFound 的 handler 进行处理。默认关闭
		server.WithHandleMethodNotAllowed(true),
	)
	//创建认证中间件
	//利用hertz的jwt框架，创建一个HertzJWTMiddleware
	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"code":   errno.SuccessCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		//登陆验证流程
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			//拦截前端的请求，进行用户名密码校验
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			//进行用户检查
			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{
				Username: loginVar.Username,
				Password: loginVar.Password,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "",
		TimeFunc:      time.Now,
	})

	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))
	//创建组，后续所有的请求路径以v1/user/开头
	base := r.Group("/douyin")

	//user分组

	user1 := base.Group("/user")
	//定义http请求的具体路径
	//登陆逻辑，直接走验证中间件的LoginHandler
	user1.POST("/login", authMiddleware.LoginHandler)
	//执行注册逻辑
	user1.POST("/register", handlers.Register)
	//获取用户信息
	user1.GET("", handlers.GetUserInfo)

	//comment分组

	//匹配不到对应接口
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	//匹配不到处理请求的方法
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	//等待、自旋
	r.Spin()
}
