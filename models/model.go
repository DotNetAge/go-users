package models

type User struct {
	ID     string
	Name   string
	Code   string `json:"code,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

func NewUser(name string) *User {
	return &User{
		ID:     "",
		Name:   name,
		Code:   "",
		Mobile: "",
	}
}

type Identity struct {
	ID       string
	UID      string
	Name     string
	Password string            `json:"password,omitempty"`
	Source   int               `json:"source,omitempty"`
	Claims   map[string]string `json:"claims,omitempty"`
}

type RefreshToken struct {
	Key  string
	UserName string
	UserID   string
	IID      string
	Roles    string
	Scope    string
}

func NewIdentity(name string) *Identity {
	return &Identity{
		ID:       "",
		UID:      "",
		Name:     name,
		Password: "",
		Claims:   nil,
	}
}

type Address struct {
}

type Contact struct {
}

type Role struct {
	Id   string
	Name string
	Desc string
}

//
//type User struct {
//	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
//	Name         string             `json:"name" bson:"name"`
//	Bio          string             `json:"desc" bson:"desc"`
//	Code         string             `json:"code" bson:"code"`
//	Credit       float32            `json:"credit" bson:"credit"`
//	Certified    bool               `json:"certified" bson:"certified"`
//	Mobile       string             `json:"mobile" bson:"mobile"`
//	Sex          int                `json:"sex" bson:"sex"`
//	Location     []float64          `json:"lngLat,omitempty" bson:"lngLat,omitempty"`
//	DateCreated  time.Time          `json:"dateCreated" bson:"date_created"`
//	LastModified time.Time          `json:"lastModified" bson:"last_modified"`
//	Avatar       string             `json:"imgUrl" bson:"img_url"`
//}
//
//type Identity struct {
//	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
//	UID         primitive.ObjectID `json:"uid" bson:"uid"`
//	Name        string             `json:"name" bson:"name"`
//	Source      int                `json:"source" bson:"source"`
//	Password    string             `json:"password" bson:"password"`
//	DateCreated time.Time          `json:"dateCreated" bson:"date_created"`
//}

/*
Organization 组织
Enterprise 企业
Company 公司
Merchant 商家/档主
*/
