package handlerdemo

/*
流程：
	1：解析token获得userId，获得toUserId
	2：向relation服务发送请求
	3：返回请求结果
*/

func RelationAction() {

}

/*
流程：
	1：解析token获得userId，获得toUserId
	2：向relation服务发送请求获取关注用户列表
		2.1：relation服务向user服务获取用户详细信息
	3：返回请求结果
*/

func FollowList() {

}

/*
流程：
	1：解析token获得userId，获得ToUserId
	2：向relation服务发送请求获取粉丝用户列表
		2.1：relation服务向user服务获取用户详细信息
	3：返回请求结果
*/

func FollowerList() {

}

/*
流程：
	1：解析token获得userId，获得ToUserId
	2：向relation服务发送请求获得好友用户列表（互粉）
		2.1：relation服务向user服务获取用户详细信息
	3：返回请求结果
*/

func FriendList() {

}
