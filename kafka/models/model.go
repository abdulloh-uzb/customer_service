package models

type CustomerReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Bio         string `json:"bio"`
	Addresses   []*AddressRes
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Code        string `json:"code"`
	Posts       []*Post
}

type AddressRes struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	District string `json:"district"`
	Street   string `json:"street"`
}

type Post struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Medias      []*Media
	Reviews     []*Review
}

type Media struct {
	Id     int64  `json:"id"`
	PostId int64  `json:"post_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Type   string `json:"type"`
}

type Review struct {
	Id          int64  `json:"id"`
	PostId      int64  `json:"post_id"`
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	Rating      int64  `json:"rating"`
	Description string `json:"description"`
}
