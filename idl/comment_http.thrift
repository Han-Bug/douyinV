namespace go comment

struct ActionRequest{
    1: required string token
    2: required string video_id
    3: required string action_type
    4: optional string comment_text
    5: optional string comment_id
}

struct ActionResponse{
    1: required i64 status_code
    2: required string status_msg
    3: optional Comment comment
}

struct ListRequest{
    1: required string token
    2: required string video_id
}

struct ListResponse{
    1: required i64 status_code
    2: required string status_msg
    3: required list<Comment> comment_list
}

struct Comment{
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date
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

service CommentSvr{
    ActionResponse Action(1: ActionRequest req) (api.post="douyin/comment/action")
    ListResponse List(1: ListRequest req) (api.get="douyin/comment/list")
}