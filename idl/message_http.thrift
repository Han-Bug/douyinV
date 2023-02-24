namespace go message

struct ChatReq {
    1: required string Token (api.query="token")
    2: required i64 ToUserID (api.query="to_user_id")
}

struct ChatResp {
    1: required i64 StatusCode (api.body="status_code")
    2: optional string StatusMsg (api.body="status_msg")
    3: list<Message> MessageList (api.body="message_list")
}

struct Message {
    2: required i64 ToUserID (api.body="to_user_id")
    3: required i64 FromUserID (api.body="from_user_id")
    4: required string Content (api.body="content")
    5: required i64 CreateTime (api.body="create_time")
}

service Chat {
    ChatResp Chat(1: ChatReq req) (api.get="douyin/message/chat/")
}


struct ActionReq {
    1: required string Token (api.query="token")
    2: required i64 ToUserID (api.query="to_user_id")
    3: required i64 ActionType (api.query="action_type")
    4: required string Content (api.query="content")
}

struct ActionResp {
    1: required i64 StatusCode (api.body="status_code")
    2: optional string StatusMsg (api.body="status_msg")
}

service Action {
    ActionResp Action(1: ActionReq req) (api.post="douyin/message/action/")
}