package service

import (
	"errors"
	"gorm.io/gorm"
	"vk-im/internal/db"
	"vk-im/internal/model"
	//"vk-im/internal/model"
)

/*
   1. 注册功能：用户名，密码，邮箱等信息的验证和存储，生成token（基于某种自定义规则）
   2. 登录功能：用户id/邮箱+密码+token校验，实现单点登录
   3. 用户信息修改
   4. 获取用户信息
   5. 添加好友
   6. 删除好友
   7. 获取好友列表
*/

type userService struct{}

var UserService userService

// Register TODO 注册： 必须要有用户名，邮箱，密码，生成token
func (u *userService) Register(user *model.User) error {
	// 1. 验证用户名是否重复
	_db := db.GetDb() //获取数据库实例， 目的是查这个数据库中的user表里的username字段，防止新注册的用户名重复
	tx := _db.Begin()
	var count int64
	_db.Model(&model.User{}).Where("username = ?", user.Username).Count(&count) // 查username表里的值，看有没有和传进来的username重复
	if count > 0 {
		tx.Rollback()
		return errors.New("用户已存在")
	}
	// 2. 验证邮箱是否重复
	_db.Model(&model.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		tx.Rollback()
		return errors.New("邮箱重复")
	}
	// 3. 密码加密
	user.Password = EncryptPassword(user.Password)

	// 4. 生成token
	user.Token, _ = GenerateToken(user.Username, user.Email)

	// 5. 存储用户信息
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Login TODO 登录验证： 校验用户名/邮箱+密码+token
func (u *userService) Login(user *model.User) (bool, error) {
	// 1. 验证用户名/邮箱+密码
	_db := db.GetDb()
	// 如果数据库里没有user表，就自动创建一个
	if _db.Migrator().HasTable(&model.User{}) == false {
		err := db.CreateTable(&model.User{})
		if err != nil {
			return false, err
		}
	}
	var queryUser *model.User
	_db.First(&queryUser, "username = ? or email = ?", user.Username, user.Email)

	return queryUser.Password == user.Password, nil
}

// TODO 信息修改
func (u *userService) ModifyUserInfo(user *model.User) (bool, error) {
	var queryUser *model.User
	_db := db.GetDb()
	_db.First(&queryUser, "id = ?", user.Id)

	var nullId int32 = 0
	if queryUser.Id == nullId {
		return false, errors.New("用户不存在")
	}
	// 修改逻辑
	queryUser.Username = user.Username
	queryUser.Sex = user.Sex
	queryUser.Description = user.Description

	// 不可修改项：id
	if queryUser.Id != user.Id {
		return false, errors.New("用户id不能修改")
	}

	// TODO 特殊修改项：密码，邮箱

	//保存
	_db.Save(queryUser)
	return true, nil
}

func (u *userService) GetUserInfo(id string) (*model.User, interface{}) {
	var queryUser *model.User
	_db := db.GetDb()
	if err := _db.First(&queryUser, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return queryUser, true
}

func (u *userService) GetUserFriends(user_id string) ([]*model.UserFriends, bool) {
	var friends []*model.UserFriends
	_db := db.GetDb()
	_db.Table("user").Select("user.*").
		Joins("left join user_friends on user.id = user_friends.friend_id").
		Where("user_friends.user_id = ?", user_id).
		Scan(&friends)
	return friends, true
}

func EncryptPassword(password string) string {
	// TODO 实现： 密码按照hash 算法加密，返回加密后的密码

	return password
}

// GenerateToken TODO Token生成策略： 用sha256加密用户名的最后3位，邮箱的@符号前倒数3位，生成token
func GenerateToken(username string, email string) (string, error) {
	// TODO 实现
	return "1234", nil
}
