// Package constants 定义一些用得着的常量
package constants

const (
	CommentTableName           = "comment"
	UserTableName              = "user"
	SecretKey                  = "secret key"
	IdentityKey                = "id"
	Total                      = "total"
	Notes                      = "notes"
	CommentID                  = "comment_id"
	UserID                     = "user_id"
	ApiServiceName             = "api"
	CommentServiceName         = "comment"
	VideoServiceName           = "video"
	UserServiceName            = "user"
	MySQLDefaultDSN            = "root:password@tcp(47.108.237.99:3311)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                = "127.0.0.1:2379"
	CPURateLimit       float64 = 80.0
	DefaultLimit               = 10
)
