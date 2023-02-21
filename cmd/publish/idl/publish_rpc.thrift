namespace go publish


struct ActionRequest{
    1: required binary   Data
    2: required string   UserId
    3: required string   Title
}

struct ActionResponse{
    1: required i64  StatusCode
    2: required string StatusMsg
}

struct ListRequest{
    1: required string UserId
    2: required string ToUserId
}

struct ListResponse{
    1: required i64  StatusCode
    2: required string StatusMsg
    3: list<Video> VideoList
}

struct Video{
    1: required i64 ID
    2: required User Author
    3: required string PlayUrl
    4: required string CoverUrl
    5: required i64 FavoriteCount
    6: required i64 CommentCount
    7: required bool IsFavorite
    8: required string Title
}

struct User {
    1: required i64 ID
    2: required string Name
    3: required i64 FollowCount
    4: required i64 FollowerCount
    5: required bool IsFollow
    6: required string AvatarUrl
    7: required string BackgroundImage
    8: required string Signature
    9: required string TotalFavorited
    10: required i64 WorkCount
    11: required i64 FavoriteCount
}

service Publish{
    ActionResponse Action(1: ActionRequest req)
    ListResponse List(1: ListRequest req)
}