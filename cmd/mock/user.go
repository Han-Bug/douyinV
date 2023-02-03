package mock

import (
	"tiktok/biz/model/feed"
	"tiktok/config"
)

type User struct {
	Basic    feed.User
	Username string
	Password string
	Token    string
}

func getUsers() []User {
	return []User{
		{
			Basic: feed.User{
				ID:            1,
				Name:          "Bonnenult",
				FollowCount:   1,
				FollowerCount: 1,
				IsFollow:      false,
				AvatarUrl:     config.ClientUrl + "/resource/avatar/1.jpeg",
			},
			Username: "ljd",
			Password: "123456",
			Token:    "1",
		},
		{
			Basic: feed.User{
				ID:            2,
				Name:          "admin",
				FollowCount:   1,
				FollowerCount: 1,
				IsFollow:      false,
				AvatarUrl:     config.ClientUrl + "/resource/avatar/2.jpeg",
			},
			Username: "admin",
			Password: "123456",
			Token:    "2",
		},
	}
}

func GetUserByUsername(un string) *User {
	users := getUsers()
	for _, v := range users {
		if v.Username == un {
			return &v
		}
	}
	return nil
}

func GetUserByID(id int) *User {
	users := getUsers()
	for _, v := range users {
		if v.Basic.ID == int64(id) {
			return &v
		}
	}
	return nil
}
