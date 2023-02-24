namespace go user

struct RegisterRequest{
    1: required string username
    2: required string password
}
struct RegisterResponse{
    1: required i64 status_code
    2: required string status_message
    3: required i64 user_id
}
struct LoginRequest{
    1: required string username
    2: required string password
}
struct LoginResponse{
    1: required i64 status_code
    2: required string status_message
    3: required i64 user_id
}
struct InfoRequest{
    1: required i64 user_id
    2: required i64 to_user_id
}
struct InfoResponse{
    1: required i64 status_code
    2: required string status_message
    3: required User user
}
struct InfoInBatchesRequest{
    1: required i64 user_id
    2: required list<i64> to_user_id_list
}
struct InfoInBatchesResponse{
    1: required i64 status_code
    2: required string status_message
    3: required list<User> user_list
    4: required list<string> error_message_list
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

service UserSvr{
    LoginResponse Login(1: LoginRequest req)
    RegisterResponse Register(1: RegisterRequest req)
    InfoResponse Info(1: InfoRequest req)
    InfoInBatchesResponse InfoInBatches(1: InfoInBatchesRequest req)
}
