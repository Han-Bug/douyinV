package dal

import "github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"

// Init init dal
// 初始化，调用者是main.go
func Init() {
	db.Init() // mysql init
}
