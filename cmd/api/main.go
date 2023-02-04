package main

import (
	"context"
	"douyinV/cmd/api/constants"
	"douyinV/cmd/api/handlers/user"
	"douyinV/cmd/api/rpc"
	usersvr "douyinV/cmd/user/kitex_gen/user"
	"douyinV/pkg/errno"

	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"

	"time"
)

const (
	hp = "127.0.0.1:8080"
)

func main() {
	rpc.Init()
	//利用hertz创建一个web服务，接收http请求
	r := server.New(
		//规定好ip-port
		server.WithHostPorts(hp),
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
				return e.(errno.ErrNo).Message
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
			var loginVar user.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			//进行用户检查
			return rpc.UserLogin(context.Background(), &usersvr.LoginRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	// 打印错误
	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))
	douyinGroup := r.Group("/douyin")
	userGroup := douyinGroup.Group("/user")
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	////分组，RESTFUL API
	//note1 := douyinGroup.Group("/note")
	////这个组必须使用验证中间件的MiddlewareFunc（token验证）
	//note1.Use(authMiddleware.MiddlewareFunc())
	//note1.GET("/query", handlers.QueryNote)
	//note1.POST("", handlers.CreateNote)
	//note1.PUT("/:note_id", handlers.UpdateNote)
	//note1.DELETE("/:note_id", handlers.DeleteNote)

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
