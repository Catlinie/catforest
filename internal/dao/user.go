package dao

import (
	"catlinie_test01/internal/model"
)

// GetLoginInfo 用户登陆获取密码
func (d *Dao) GetLoginInfo(account, itype string) (string, error) {
	userAuth, err := d.getSpecificAuth(account, itype)
	if err != nil {
		return "", nil
	}
	return userAuth.Credential, nil
}

// CreateUserInfo 创建一个user与对应的auth
func (d *Dao) CreateUserInfo(user *model.User, account, password, itype string) error {
	// Create()里面要一个指针
	return d.beginTrans(func() error {
		var count int
		err := d.Engine.Model(user).Where("account = ? && identty_type = ?", account, itype).Count(&count).Error
		if err != nil || count > 0 {
			return err
		}
		err = d.Engine.Create(user).Error
		if err != nil {
			return err
		}
		userAuth := &model.UserAuth{
			Cid:          user.Cid,
			Identifier:   account,
			IdentityType: itype,
			Credential:   password,
		}
		err = d.Engine.Create(userAuth).Error
		return err
	})
}

// GetEntireInfo 获取一个用户的完整资料信息
func (d *Dao) GetEntireInfo(account, itype string) (*model.User, error) {
	userAuth, err := d.getSpecificAuth(account, itype)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	// Model与First里面要的都是指针
	err = d.Engine.Model(user).Where("cid = ?", userAuth.Cid).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Dao) SetUserInfo(username, desc, masterName, url, account, itype string, gender uint8) error {
	return d.beginTrans(func() error {
		auth, err := d.getSpecificAuth(account, itype)
		if err != nil {
			return err
		}
		return d.Engine.Table("users").Where("cid = ?", auth.Cid).Update(
			map[string]interface{}{
				"username":    username,
				"gender":      gender,
				"desc":        desc,
				"master_name": masterName,
				"url":         url,
			}).Error
	})
}

func (d *Dao) getSpecificAuth(account, itype string) (*model.UserAuth, error) {
	var userAuth model.UserAuth
	err := d.Engine.Model(&userAuth).Where("identifier = ? AND identity_type = ?", account, itype).First(&userAuth).Error
	if err != nil {
		return nil, err
	}
	return &userAuth, nil
}
