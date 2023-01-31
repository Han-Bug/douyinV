package rpc

// InitRPC init rpc client
// 进行rpc客户端初始化
// 调用两个客户端初始化方法
// 这个函数的调用者在main.go中
func InitRPC() {
	//创建一个UserRpc的客户端
	initUserRpc()
}
