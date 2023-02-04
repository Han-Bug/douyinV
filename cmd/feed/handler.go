package main

import (
	"context"
	api "tiktok/feed/kitex_gen/api"
	"tiktok/mock"
)

// FeedImpl implements the last service interface defined in the IDL.
type FeedImpl struct{}

// Feed implements the FeedImpl interface.
func (s *FeedImpl) Feed(ctx context.Context, req *api.FeedReq) (resp *api.FeedResp, err error) {
	resp = &api.FeedResp{
		StatusCode: 0,
		VideoList:  mock.GetVideos(),
	}
	return
}
