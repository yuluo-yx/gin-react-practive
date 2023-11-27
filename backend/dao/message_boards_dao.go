package dao

import (
	"backend/model"
	"context"
	"gorm.io/gorm"
)

type MsgDao struct {
	*gorm.DB
}

// NewMsgDao 创建 MsgDao 通过上下文的方式创建
func NewMsgDao(ctx context.Context) *MsgDao {
	return &MsgDao{NewDBClient(ctx)}
}

// NewMsgDaoByDB 通过 DB 去创建 MsgDao
func NewMsgDaoByDB(db *gorm.DB) *MsgDao {
	return &MsgDao{db}
}

func (dao *MsgDao) GetMsg() (msgList []*model.Message, err error) {

	err = dao.Find(&msgList).Error
	return
}

func (dao *MsgDao) DeleteMsg(id uint, uId uint) error {

	return dao.DB.Where("id=?", id).Where("user_Id=?", uId).Delete(&model.Message{}).Error
}

func (dao *MsgDao) SaveMsg(msg *model.Message) error {

	return dao.DB.Model(&model.Message{}).Create(&msg).Error
}

type MsgResult struct {
	model.Message
	UserName string
}

func (dao *MsgDao) MsgListByCondition(page model.BasePage,
	userDao *UserDao) (msgRes []*MsgResult, err error) {

	var msgs []*model.Message
	err = dao.DB.Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).Find(&msgs).Error

	for _, msg := range msgs {
		var msgR MsgResult
		user, _ := userDao.GetUserById(msg.UserId)
		msgR.UserName = user.UserName
		msgR.Text = msg.Text
		msgR.UserId = msg.UserId

		msgRes = append(msgRes, &msgR)
	}

	return
}

func (dao *MsgDao) CountMsg() (total int64, err error) {
	err = dao.DB.Model(&model.Message{}).Count(&total).Error
	return
}
