namespace go video


struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
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
struct FeedRequest{
    1: optional i64 latest_time
    2: optional i64 user_id
}

struct FeedResponse{
    1: BaseResp base_resp
    2: list<Video> video_list
    3: optional i64 next_time
}

struct VideoPublishActionRequest{
    1: required i64 user_id
    2: required byte data
    3: required string title
}

struct VideoPublishActionResponse{
    1: BaseResp base_resp
}


struct VideoListRequest{
    1: required i64 user_id
}

struct VideoListResponse{
    1: BaseResp base_resp
    2: list<Video> video_list
}

service VideoService{
    FeedResponse Feed(1: FeedRequest req)
    VideoPublishActionResponse PublishVideo(1: VideoPublishActionRequest req)
    VideoListResponse ListVideo(1: VideoPublishActionRequest req)
}

