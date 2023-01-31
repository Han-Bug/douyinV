namespace go user

struct RegisterRequest{
    1:string username
    2:string password
}
struct RegisterResponse{
    1:string id
    2:string token
}
struct LoginRequest{
    1:string username
    2:string password
}
struct LoginResponse{
    1:string id
    2:string token
}
struct UserInfoRequest{
    1:string id
    2:string token
}
struct UserInfoResponse{
    1:string id
    2:string token
}
