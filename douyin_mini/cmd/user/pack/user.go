package pack

import (
	"douyin_mini/cmd/user/dal/db"
	"douyin_mini/kitex_gen/user"
)

// User pack user info
// 将db里边的User类型转换成thrift里边定义的User类型
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollowed:    false,
	}
}

// Users pack list of user info
// 将一大堆的db的User转换成thrift中的User
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
