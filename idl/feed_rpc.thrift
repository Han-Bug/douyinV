namespace go api

struct FeedReq {
}

struct FeedResp {
    1: required i64 status_code
    2: optional string status_msg
    3: list<Video> video_list
    4: optional i64 next_time
}

struct Video {
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
    6: required string avatar
    7: required string background_image
    8: required string signature
    9: required string total_favorited
    10: required i64 work_count
    11: required i64 favorite_count
}

service Feed {
    FeedResp feed(1: FeedReq req)
}