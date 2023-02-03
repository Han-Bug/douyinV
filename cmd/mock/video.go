package mock

import (
	"strconv"
	"tiktok/biz/model/feed"
	"tiktok/config"
)

func GetVideos() (vs []*feed.Video) {
	for i := 1; i <= 9; i++ {
		vs = append(vs, &feed.Video{
			ID:            int64(i),
			Author:        &GetUserByID(1).Basic,
			PlayUrl:       config.ClientUrl + "/resource/video/" + strconv.Itoa(i) + ".mp4",
			CoverUrl:      config.ClientUrl + "/resource/cover/" + strconv.Itoa(i) + ".jpeg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         "测试视频" + strconv.Itoa(i) + "号",
		})
	}
	return
}
