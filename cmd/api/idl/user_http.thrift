namespace go user

struct RegisterRequest{
    1:string username
    2:string password
}
struct RegisterResponse{
    1:i64 status_code
    2:string status_msg
    3:i64 user_id
    4:string token

}
struct LoginRequest{
    1:string username
    2:string password
}
struct LoginResponse{
    1:i64 status_code
    2:string status_msg
    3:i64 user_id
    4:string token
}
struct InfoRequest{
    1:string user_id
    2:string token
}
struct InfoResponse{
    1:i64 status_code
    2:string status_msg
    3:UserParam user
}
struct UserParam{
    1:i64 id
    2:string name
    3:i32   follow_count
    4:i32   follower_count
    5:bool  is_follow
}
service UserSvr{
    RegisterResponse Register(1:RegisterRequest req)(api.post="/douyin/user/register")
    LoginResponse Login(1:LoginRequest req)(api.get="/douyin/user")
    InfoResponse Info(1:InfoRequest req)(api.post="/douyin/user/login")
}