package services

import (
	"context"
	"douyinV/user/dal/mysqlDB"
	"douyinV/user/kitex_gen/user"
	"douyinV/user/pkg/errno"
	utils "douyinV/user/pkg/md5"
	"github.com/go-sql-driver/mysql"
	"strings"
)

// 此层错误处理：打印错误日志，上传自定义错误（用于给客户端返回错误信息）

// TODO 常量整理

const (
	Md5key = "douyin"
)

func Register(ctx context.Context, req user.RegisterRequest) (int64, error) {

	// 获取参数
	username := req.GetUsername()
	password := req.GetPassword()
	// TODO 合法性校验
	// 帐号长度应位于4至16

	// 密码长度应位于6至16
	// 密码加密
	pd := utils.EncryptMd5(strings.Join([]string{Md5key, password}, ""))
	// 写入数据库
	id, err := mysqlDB.CreateUser(ctx, username, pd)
	if err != nil {
		// 如果是因为主键冲突导致的错误（即已存在该用户名）
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			if errMySQL.Number == 1062 {
				// 账号已存在
				return 0, errno.UserAlreadyExistErr
			} else {
				// TODO 如果是其他数据库错误
				return 0, err
			}
		} else {
			// TODO 如果是其他数据库错误
			return 0, err
		}
	}
	// 返回结果
	return id, nil
}
