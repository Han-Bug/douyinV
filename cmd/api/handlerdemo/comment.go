package handlerdemo

/*
流程：
	1：获取token解析得userId，获取videoId
	2.1：向comment服务发送请求创建/删除评论，获得评论comment的详细数据
	2.2：向user服务发送请求获取user信息
	3：拼装返回结果并返回
*/

func CommentAction() {

}

/*
流程：
	1：获取token解析得userId，获取videoId
	2：向comment服务请求获取视频的评论列表
	2.1：comment服务根据评论列表的authorId向user服务请求获取comment作者的用户信息
	2.2：comment服务返回拼装好的结果
	3：拼装返回结果并返回
*/

func CommentList() {

}
