package mysqlDB

import (
	"context"
	"gorm.io/gorm"
	"time"
)

var (
	// ErrUser 使用既定User值
	ErrUser = User{
		ID:       0,
		Name:     "err",
		Password: "",
		CUDTime: CUDTime{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
)

func CreateUser(ctx context.Context, username string, password string) (int64, error) {
	newUser := User{
		Name:     username,
		Password: password,
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&newUser).Error
	})
	if err != nil {
		return 0, err
	}
	// 获取新用户的id
	u := User{}
	DB.Where(&User{Name: username}).Limit(1).Find(&u)
	return u.ID, nil
}

// TODO 将返回类型user.User更改为数据库模型User，并在上级调用层组装成user.User

// UserInfo 用户编号为id的用户查询编号为toId的用户信息 错误上抛
func UserInfo(ctx context.Context, id int64) (*User, error) {
	toUser := User{}
	// 查询目标用户基本信息
	result := DB.WithContext(ctx).Where(&User{ID: id}).First(&toUser)
	if result.Error != nil {
		// 找不到目标用户
		return nil, result.Error
	}
	return &toUser, nil

}

// TODO 将返回类型user.User更改为数据库模型User，并在上级调用层组装成user.User

// UserInfoInBatches 批量查询用户信息
func UserInfoInBatches(ctx context.Context, ids []int64) (users []*User, errMsgs []string) {
	for i, id := range ids {
		var u User
		if id <= 0 {
			users[i] = nil
			errMsgs[i] = "invalid user_id"
		} else if err := DB.WithContext(ctx).Where(&User{ID: id}).First(&u).Error; err != nil {
			users[i] = nil
			errMsgs[i] = err.Error()
		}
		users[i] = &u
		errMsgs[i] = ""
	}
	return
	//err := DB.WithContext(ctx).Model(&User{}).Where("id in (?)", toIds).Find(&users).Error
	//if err != nil {
	//	return nil, err
	//}
	//res := make([]*user.User, len(toIds))
	//for i, u := range users {
	//	isFollow := false
	//	if id != 0 {
	//		// TODO 采用图数据库存储关系
	//		// 此处为使用MySql存储关系（暂定方案）
	//		relation := Relation{UserId: id, FollowerId: u.ID}
	//		result := DB.First(&relation)
	//		if err := result.Error; err != nil {
	//			// 如果找不到相关relation
	//			if errors.Is(err, gorm.ErrRecordNotFound) {
	//				isFollow = false
	//			} else {
	//				// 其他错误
	//				klog.Error("批量查询时出错")
	//				return nil, err
	//			}
	//		}
	//		isFollow = true
	//	}
	//	// 若在查询关注数或粉丝数时出错，直接将相关值置0
	//	followCount, err := FollowCount(u.ID)
	//	if err != nil {
	//		followCount = 0
	//	}
	//	followerCount, err := FollowerCount(u.ID)
	//	if err != nil {
	//		followerCount = 0
	//	}
	//	res[i] = &user.User{
	//		Id:              u.ID,
	//		Name:            u.Name,
	//		FollowCount:     followCount,
	//		FollowerCount:   followerCount,
	//		IsFollow:        isFollow,
	//		Avatar:          "https://p3-passport.byteimg.com/img/mosaic-legacy/3793/3114521287~64x64.awebp",
	//		BackgroundImage: "",
	//		Signature:       "test signature",
	//		TotalFavorited:  "0",
	//		WorkCount:       0,
	//		FavoriteCount:   0,
	//	}
	//}
	//return res, nil
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
