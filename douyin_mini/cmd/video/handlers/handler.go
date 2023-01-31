package handlers

import (
	"context"
	video "douyin_mini/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.VideoPublishActionRequest) (resp *video.VideoPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) ListVideo(ctx context.Context, req *video.VideoPublishActionRequest) (resp *video.VideoListResponse, err error) {
	// TODO: Your code here...
	return
}
