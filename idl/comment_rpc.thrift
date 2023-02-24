namespace go comment

struct ActionRequest{
    1: required i64 user_id
    2: required i64 video_id
    3: required i64 action_type
    4: optional string comment_text
    5: optional i64 comment_id
}

struct ActionResponse{
    1: required i64 status_code
    2: required string status_message
    3: required Comment comment
}

struct ListRequest{
    1: required i64 user_id
    2: required i64 video_id
}

struct ListResponse{
    1: required i64 status_code
    2: required string status_message
    3: list<Comment> comment_list
}
struct CountRequest{
    1: required i64 video_id
}
struct CountResponse{
    1: required i64 status_code
    2: required string status_message
    3: required i64 comment_count
}

struct CountInBatchesRequest{
    1: required list<i64> video_id_list
}

struct CountInBatchesResponse{
    1: required i64 status_code
    2: required string status_message
    3: required list<i64> comment_count_list
    4: required list<string> error_message_list
}

struct Comment{
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date
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

service CommentSvr{
    ActionResponse Action(1: ActionRequest req)
    ListResponse List(1: ListRequest req)
    CountResponse Count(1: CountRequest req)
    CountInBatchesResponse CountInBatches(1: CountInBatchesRequest req)
}