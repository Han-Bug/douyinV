namespace go relation
//基础的响应格式
struct BaseResp{
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}
struct RelationActionRequest{
    1: required i64 user_id //自己
    2: required i64 to_user_id //对方
    3: required i32 action_type // 1-关注，2-取消关注
}

struct RelationActionResponse{
   1: BaseResp base_resp
}


//查看我关注了谁
struct RelationFollowListRequest{
     1: required i64 user_id
}


struct RelationFollowListResponse{
    1:BaseResp base_resp
    2:list<User> user_list
}

//查看谁关注了我
struct RelationFollowerListRequest{
    1: required i64 user_id

}
struct RelationFollowerListResponse{
    1: BaseResp base_resp
    2: list<User> user_list
}

struct User{
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_followed
}

service RelationService{
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    //查看我关注了谁
    RelationFollowListResponse GetRelationFollowList (1: RelationFollowListRequest req)
    //查看谁关注了我
    RelationFollowerListResponse GetRelationFollowerList(1: RelationFollowerListRequest req)
}