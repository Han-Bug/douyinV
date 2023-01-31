namespace go comment

//基础的响应格式
struct BaseResp{
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User{
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_followed
}

struct Comment{
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date
}

struct CommentActionRequest{
    1: required i64 user_id//用户id
    2: required i64 video_id// 视频id
    3: required i32 action_type // 1-发布评论，2-删除评论
    4: optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse{
    1: required BaseResp base_resp
    2: optional Comment comment//评论成功返回评论内容，不需要重新拉取整个列表
}


struct CommentListRequest{
    1: required i64 user_id
    2: required i64 video_id
}

struct CommentListResponse{
    1: required BaseResp base_resp
    2: list<Comment> comment_list
}

service CommentService{
    CommentActionResponse ActionComment(1: CommentActionRequest req)
    CommentListResponse ListComment(1: CommentActionRequest req)
}