package mock

import (
	"tiktok/config"
	"tiktok/feed/kitex_gen/api"
)

type User struct {
	Basic    api.User
	Username string
	Password string
	Token    string
}

func getUsers() []User {
	return []User{
		{
			Basic: api.User{
				Id:            1,
				Name:          "Bonnenult",
				FollowCount:   1,
				FollowerCount: 1,
				IsFollow:      false,
				Avatar:        config.ClientUrl + "/resource/avatar/1.jpeg",
			},
			Username: "ljd",
			Password: "123456",
			Token:    "1",
		},
		{
			Basic: api.User{
				Id:            2,
				Name:          "admin",
				FollowCount:   1,
				FollowerCount: 1,
				IsFollow:      false,
				Avatar:        config.ClientUrl + "/resource/avatar/2.jpeg",
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
		if v.Basic.Id == int64(id) {
			return &v
		}
	}
	return nil
}
