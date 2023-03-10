namespace go api

struct ChatReq {
    1: required i64 from_user_id
    2: required i64 to_user_id
    3: required i64 latest_time
}

struct ChatResp {
    1: required i64 status_code
    2: optional string status_msg
    3: list<Msg> message_list
}

struct Msg {
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: required i64 create_time
}

struct ActionReq {
    1: required i64 from_user_id
    2: required i64 to_user_id
    3: required i64 action_type
    4: required string content
}

struct ActionResp {
    1: required i64 status_code
    2: optional string status_msg
}

struct ChatLatestReq {
    1: required i64 from_user_id
    2: required i64 to_user_id
}

struct ChatLatestResp {
    1: Msg message
}

service Message {
    ChatResp chat(1: ChatReq req)
    ActionResp action(1: ActionReq req)
    ChatLatestResp chatLatest(1: ChatLatestReq req)
}
