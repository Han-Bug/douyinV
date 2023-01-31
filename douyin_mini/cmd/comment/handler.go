package comment

import (
	"context"
	comment "douyin_mini/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// ActionComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ActionComment(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
