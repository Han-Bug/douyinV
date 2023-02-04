package mysqlDB

import (
	"context"
	"douyinV/cmd/user/kitex_gen/user"
	. "douyinV/cmd/user/model"
	"errors"
	"gorm.io/gorm"
)

// 错误上抛而不处理

func CreateUser(ctx context.Context, username string, password string) (int64, error) {
	newUser := User{
		Name:     username,
		Password: password,
	}
	tx := DB.Begin()
	if err := tx.WithContext(ctx).Create(&newUser).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	// 获取新用户的id
	u := User{}
	DB.Where(&User{Name: username}).First(&u)
	return u.ID, nil
}

// UserInfo 用户编号为id的用户查询编号为toId的用户信息 错误上抛
func UserInfo(ctx context.Context, id int64, toId int64) (*user.UserInfoResponse, error) {
	toUser := User{}
	// 查询目标用户基本信息
	result := DB.WithContext(ctx).Where(&User{ID: toId}).First(&toUser)
	if result.Error != nil {
		// 找不到目标用户
		return nil, result.Error
	}
	// TODO 查询当前用户是否关注目标用户(redis)
	relation := Relation{UserId: id, FollowId: toId}
	result = DB.First(&relation)
	var isFollow bool
	if err := result.Error; err != nil {
		// 如果找不到相关relation
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isFollow = false

		} else {
			// 其他错误
			return nil, err
		}
	}
	isFollow = true
	resParam := user.UserInfoResponse{
		Id:            toUser.ID,
		Name:          toUser.Name,
		FollowCount:   toUser.FollowCount,
		FollowerCount: toUser.FollowerCount,
		IsFollow:      isFollow,
	}
	return &resParam, nil

}

// UserInfoAnonymous 查询编号为toId的用户（不考虑是否关注）
func UserInfoAnonymous(ctx context.Context, toId int64) (*user.UserInfoResponse, error) {
	u := User{}
	// 查询目标用户基本信息
	result := DB.WithContext(ctx).Where(&User{ID: toId}).First(&u)
	if err := result.Error; err != nil {
		// 找不到目标用户
		return nil, err
	}
	resParam := user.UserInfoResponse{
		Id:            u.ID,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      false,
	}
	return &resParam, nil
}

// UserLogin 登录 返回用户id与err
func UserLogin(ctx context.Context, username string, password string) (int64, error) {
	u := User{}
	result := DB.WithContext(ctx).Where(&User{Name: username, Password: password}, "name", "password").First(&u)
	if result.Error != nil {
		return 0, result.Error
	}
	return u.ID, nil
}

// Unscoped 彻底删除（仅测试用）永久删除指定用户
func Unscoped(ctx context.Context, id int64) error {
	u := User{ID: id}
	tx := DB.Begin()
	if err := tx.WithContext(ctx).Unscoped().Delete(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
