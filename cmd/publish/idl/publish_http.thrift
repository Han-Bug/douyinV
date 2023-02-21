namespace go publish

struct ActionRequest{
    1: required binary   Data (api.query="data")
    2: required string   Token (api.query="token")
    3: required string   Title (api.query="title")
}

struct ActionResponse{
    1: required i64  StatusCode (api.query="status_code")
    2: required string StatusMsg (api.body="status_msg")
}

struct ListRequest{
    1: required string Token (api.query="token")
    2: required string UserId (api.query="user_id")
}

struct ListResponse{
    1: required i64  StatusCode (api.query="status_code")
    2: required string StatusMsg (api.body="status_msg")
    3: list<Video> VideoList (api.body="video_list")
}

struct Video{
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
    7: required string BackgroundImage (api.body="background_image")
    8: required string Signature (api.body="signature")
    9: required string TotalFavorited (api.body="total_favorited")
    10: required i64 WorkCount (api.body="work_count")
    11: required i64 FavoriteCount (api.body="favorite_count")
}

service Publish{
    ActionResponse Action(1: ActionRequest req) (api.post="douyin/publish/action")
    ListResponse List(1: ListRequest req) (api.get="douyin/publish/list")
}