package mysqlDB

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

func CreateVideo(ctx context.Context, authorId int64, title string, playUrl string, coverUrl string) error {
	video := Video{
		UserId:   authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	if err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(video).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println("创建视频对象时出错")
		fmt.Printf("userId=%s ;title=%s ;playurl=%s ;coverurl=%s \n")
		return err
	}
	return nil
}

func VideoListByUserId(ctx context.Context, id int64) ([]*Video, error) {
	var videos []*Video
	res := DB.WithContext(ctx).Where(&Video{UserId: id}).Find(&videos)
	if res.Error != nil {
		fmt.Println("查询视频列表时出错")
		fmt.Printf("id=%s \n", id)
		return nil, res.Error
	}
	return videos, nil

}
