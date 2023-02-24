namespace go user

struct RegisterRequest{
    1:required string username
    2:required string password
}
struct RegisterResponse{
    1:required i64 status_code
    2:required string status_msg
    3:required i64 user_id
    4:required string token

}
struct LoginRequest{
    1:required string username
    2:required string password
}
struct LoginResponse{
    1:required i64 status_code
    2:required string status_msg
    3:required i64 user_id
    4:required string token
}
struct InfoRequest{
    1:required string user_id
    2:required string token
}
struct InfoResponse{
    1:required i64 status_code
    2:required string status_msg
    3:required User user
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


service UserSvr{
    RegisterResponse Register(1:RegisterRequest req)(api.post="/douyin/user/register")
    LoginResponse Login(1:LoginRequest req)(api.get="/douyin/user")
    InfoResponse Info(1:InfoRequest req)(api.post="/douyin/user/login")
}

struct ActionRequest{
    1: required string token
    2: required string to_user_id
    3: required string action_type
}

struct ActionResponse{
    1: required i64 status_code
    2: required string status_msg
}

struct FollowListRequest{
    1: required string user_id
    2: required string token
}

struct FollowListResponse{
    1: required i64 status_code
    2: required string status_msg
    3: list<User> user_list
}

struct FollowerListRequest{
    1: required string user_id
    2: required string token
}

struct FollowerListResponse{
    1: required i64 status_code
    2: required string status_msg
    3: list<User> user_list
}

struct FriendListRequest{
    1: required string user_id
    2: required string token
}
struct FriendListResponse{
    1: required i64 status_code
    2: required string status_msg
    3: list<User> user_list
}

service RelationSvr{
    ActionResponse Action(1: ActionRequest req)(api.post="/douyin/relation/action")
    FollowListResponse FollowList(1: FollowListRequest req)(api.get="/douyin/relation/follow/list")
    FollowerListResponse FollowerList(1: FollowerListRequest req)(api.get="/douyin/relation/follower/list")
    FriendListResponse FriendList(1: FriendListRequest req)(api.get="/douyin/relation/friend/list")
}
