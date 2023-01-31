namespace go user

//基础的响应格式
struct BaseResp{
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct RegisterUserRequest{
    1: string username
    2: string password
}

struct RegisterUserResponse{
    1: BaseResp base_resp
    2: required i64 user_id
}

struct CheckUserRequest{
    1: string username
    2: string password
}

struct CheckUserResponse{
    1:i64 user_id
    2: BaseResp base_resp
}

struct User{
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_followed
}

struct  UserInfoRequest{
    1: required list<i64> user_ids
}

struct UserInfoResponse{
    1: BaseResp base_resp
    2: list<User> users
}

service UserService{
    RegisterUserResponse  RegisterUser(1: RegisterUserRequest req)
    CheckUserResponse     CheckUser (1:CheckUserRequest req)
    UserInfoResponse      GetUserInfo(1: UserInfoRequest req)
}