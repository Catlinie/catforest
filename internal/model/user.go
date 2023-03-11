package model

type User struct {
	Id           uint32 `json:"id" gorm:"primary_key"`
	Cid          uint32 `json:"cid"`
	Username     string `json:"username"`
	Desc         string `json:"desc"`
	MasterName   string `json:"master_name"`
	Url          string `json:"url"`
	Gender       uint8  `json:"gender"`
	Times        uint32 `json:"times"`
	Bills        uint32 `json:"bills"`
	RegisterTime int64  `json:"register_time"`
	LoginTime    int64  `json:"login_time"`
	IsBaned      uint8  `json:"is_baned"`
}

func (user *User) TableName() string {
	return "users"
}
