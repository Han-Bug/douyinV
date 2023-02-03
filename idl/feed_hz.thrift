namespace go feed

struct FeedReq {
    1: optional i64 LatestTime (api.query="latest_time")
    2: optional string Token (api.query="token")
}

struct FeedResp {
    1: required i64 StatusCode (api.body="status_code")
    2: optional string StatusMsg (api.body="status_msg")
    3: list<Video> VideoList (api.body="video_list")
    4: optional i64 NextTime (api.body="next_time")
}

struct Video {
    1: required i64 ID (api.body="id")
    2: required User Author (api.body="author")
    3: required string PlayUrl (api.body="play_url")
    4: required string CoverUrl (api.body="cover_url")
    5: required i64 FavoriteCount (api.body="favorite_count")
    6: required i64 CommentCount (api.body="comment_count")
    7: required bool IsFavorite (api.body="is_favorite")
    8: required string Title (api.body="title")
}

struct User {
    1: required i64 ID (api.body="id")
    2: required string Name (api.body="name")
    3: required i64 FollowCount (api.body="follow_count")
    4: required i64 FollowerCount (api.body="follower_count")
    5: required bool IsFollow (api.body="is_follow")
    6: required string AvatarUrl (api.body="avatar")
}

service Feed {
    FeedResp Feed(1: FeedReq req) (api.get="douyin/feed")
}