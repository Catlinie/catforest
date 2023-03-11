package model

type UserAuth struct {
	Id           uint32 `json:"id"  gorm:"primary_key"`
	Cid          uint32 `json:"cid"`
	IdentityType string `json:"identity_type"`
	Identifier   string `json:"identifier"`
	Credential   string `json:"credential"`
}

func (u UserAuth) TableName() string {
	return "user_auths"
}
