package handlerdemo

/*
publish/list
req:
	token: 用户鉴权，同时标识请求用户的编号
	user_id: 查询视频列表的目标用户编号
流程：
	1: 根据传入数据获取 userId 和 toUserId
	2.1: 向relation服务查询是否已对用户关注 is_follow
	2.2: 向publish服务查询视频列表 视频列表包含视频的固定信息如playUrl
	2.2.1: 根据视频列表中的视频，向favorite服务查询相关视频的favorite_count
	2.2.2: 根据视频列表中的视频，向comment服务查询视频的评论总数
	2.3: 向favorite服务查询目标用户的follow_count、follower_count
	2.4: 向user服务查询目标用户的固定信息 name等
	3: 拼接返回结果并返回
*/

func PublishList() {

}

/*

 */

func PublishAction() {

}
