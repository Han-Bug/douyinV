package mysqlDB

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func CreateComment(ctx context.Context, userId int64, videoId int64, content string) (*Comment, error) {
	newComment := Comment{
		VideoId: videoId,
		UserId:  userId,
		Content: content,
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&newComment).Error
	})
	if err != nil {
		return nil, err
	}
	// 获取评论
	c := Comment{}
	if err = DB.Where(&Comment{VideoId: videoId, UserId: userId}).Limit(1).Find(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCommentsByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CommentCount(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Where(&Comment{
		VideoId: videoId,
	}).Count(&count).Error; err != nil {
		klog.Error("查询评论数出错： id=%v \n", videoId)
		return 0, err
	}
	return count, nil

}

func CommentCountInBatches(ctx context.Context, videoIds []int64) (counts []int64, errMsgs []string) {
	for i, id := range videoIds {
		var count int64
		if id <= 0 {
			errMsgs[i] = ""
			counts[i] = 0
		} else if err := DB.WithContext(ctx).Where(&Comment{VideoId: id}).Count(&count).Error; err != nil {
			errMsgs[i] = err.Error()
			counts[i] = 0
		} else {
			errMsgs[i] = ""
			counts[i] = count
		}
	}
	return counts, errMsgs
}

func DeleteComment(ctx context.Context, commentId int64) (int64, error) {
	c := Comment{
		ID: commentId,
	}
	var delCount int64
	err := DB.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Delete(&c)
		delCount = res.RowsAffected
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	return delCount, err
}
