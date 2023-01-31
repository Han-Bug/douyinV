namespace go favorite

struct BaseResp{
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct FavoriteActionRequest{
    1: required string user_id //用户id
    2: required i64 video_id //视频id
    3: i32 action_type // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse{
    1:BaseResp base_rep
}
struct FavoriteListRequest{
    1: required i64 user_id
}

struct FavoriteListResponse{
    1:BaseResp base_resp
    2: list<Video> video_list
}


struct Video{
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct User{
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_followed
}


service FavoriteService{
    FavoriteActionResponse FavoriteAction (1:FavoriteActionRequest req)
    FavoriteListResponse FavoriteList (1: FavoriteActionRequest req)
}