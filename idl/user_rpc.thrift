namespace go userRpc

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct RegisterRequest{
    1:string username
    2:string password
}
struct RegisterResponse{
    1:i64 id
    3:BaseResp base_resp
}
struct LoginRequest{
    1:string username
    2:string password
}
struct LoginResponse{
    1:i64 id
    3:BaseResp base_resp
}
struct UserInfoRequest{
    1:i64 id
    2:i64 to_id
}
struct UserInfoResponse{
    1:i64 id
    2:string name
    3:i32   follow_count
    4:i32   follower_count
    5:bool  is_follow
    6:BaseResp base_resp
}

service UserSvr{
    RegisterResponse Register(1:RegisterRequest req)
    LoginResponse Login(1:LoginRequest req)
    UserInfoResponse UserInfo(1:UserInfoRequest req)
}