package dao

import (
	"backend/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao 创建 UserDao 通过上下文的方式创建
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// NewUserDaoByDB 通过 DB 去创建UserDao
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 查看注册的用户名是否存在
func (dao *UserDao) ExistOrNotByUserName(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", name).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}

	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {

	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById 根据用户id查询用户信息
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// UpdateUserById 通过用户id 更新用户信息
func (dao *UserDao) UpdateUserById(id uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
}

func (dao *UserDao) GetUserList() (userList []*model.User, err error) {
	err = dao.DB.Find(&userList).Error
	return
}
