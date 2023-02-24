package handlerdemo

/*
	user、comment、
*/

/*
流程：
	1：解析token获取userId，获取videoId
	2：向favorite服务发送请求并获得结果
	3：拼装返回结果

*/

func FavoriteAction() {

}

// 创建favorite对象，返回创建结果

func FavoriteAction_() {

}

/*
流程：
	1：解析token获取userId，获取toUserId
	2：向favorite服务发送请求获得视频列表
		2.1：favorite服务向publish服务获取视频详细信息
		2.2：favorite服务向user服务获取用户详细信息
		2.3：favorite服务填充favoriteCount信息
		2.4：favorite服务填充
	3：拼装结果并返回
*/

func FavoriteList() {

}

// 根据videoid获取视频详细信息

func GetVideoByVideoId_() {

}

func GetUserByUserIdAndToUserId_() {

}
