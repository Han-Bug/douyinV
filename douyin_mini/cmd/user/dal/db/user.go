package db

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	//包含了gorm四个参数
	gorm.Model
	Name          string `json:"username"`
	Password      string `json:"password"`
	FollowerCount int64  `json:"followerCount"`
	FollowCount   int64  `json:"followCount"`
}

// TableName 定义表名
// 这是User这个结构的一个方法
func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
// 根据UserIDs批量获取用户
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	//创建一个User切片
	res := make([]*User, 0)
	//参数校验
	if len(userIDs) == 0 {
		return res, nil
	}
	//进行db查询
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser 创建用户们
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser 查询用户
func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("name = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
