package mysqlDB

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// FollowCount 查询关注数
func FollowCount(userId int64) (int64, error) {
	var count int64
	if userId <= 0 {
		return 0, nil
	}
	err := DB.Where(&Relation{
		FollowerId: userId,
	}).Count(&count).Error
	if err != nil {
		klog.Error("查询关注数出错： id=%v \n", userId)
	}
	return count, err
}

// FollowerCount 查询粉丝数
func FollowerCount(userId int64) (int64, error) {
	var count int64
	if userId <= 0 {
		return 0, nil
	}
	err := DB.Where(&Relation{
		UserId: userId,
	}).Count(&count).Error
	if err != nil {
		klog.Error("查询粉丝数出错： id=%v \n", userId)
	}
	return count, err
}

// TODO 采用图数据库存储关系

func GetRelationByUser(ctx context.Context, id int64, u *User) (isFollow bool, followCount int64, followerCount int64, err error) {
	isFollow = false
	if id > 0 {
		relation := Relation{
			UserId:     u.ID,
			FollowerId: id,
		}
		res := DB.WithContext(ctx).First(&relation)
		if res.Error != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				isFollow = false
			} else {
				klog.Error(err)
				return false, 0, 0, res.Error
			}
		} else {
			isFollow = true
		}
	}
	if u == nil || u.ID <= 0 {
		followCount = 0
		followerCount = 0
		err = nil
		return
	}
	followCount, err = FollowCount(u.ID)
	if err != nil {
		return
	}
	followerCount, err = FollowerCount(u.ID)
	if err != nil {
		return
	}
	return

}
func GetRelationsByUsers(ctx context.Context, id int64, users []*User) (isFollows []bool, followCounts []int64, followerCounts []int64) {
	for i, u := range users {
		isFollows[i] = false
		// 查询isFollow
		if id > 0 {
			relation := Relation{
				UserId:     u.ID,
				FollowerId: id,
			}
			err := DB.WithContext(ctx).First(&relation).Error
			if err == nil {
				isFollows[i] = true
			}

		}
		// 查询followCount、followerCount
		if u == nil || u.ID <= 0 {
			followCounts[i] = 0
			followerCounts[i] = 0
		} else {
			// 忽略查询中产生的错误
			var err error
			followCounts[i], err = FollowCount(u.ID)
			if err != nil {
				followCounts[i] = 0
			}
			followerCounts[i], err = FollowerCount(u.ID)
			if err != nil {
				followerCounts[i] = 0
			}
		}
	}
	return
}
