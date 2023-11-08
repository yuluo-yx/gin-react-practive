package services

import (
	"backend/dao"
	"backend/model"
	"backend/pkg/e"
	"backend/serializer"
	"context"
	"fmt"
	"sync"
)

type MsgService struct {
	ID     uint   `json:"id" form:"id"`
	UserId uint   `json:"user_id" form:"user_id"`
	Text   string `json:"text" form:"text"`

	model.BasePage
}

func (msgService MsgService) Publish(ctx context.Context) serializer.Response {

	var msg = model.Message{
		Text:   msgService.Text,
		UserId: msgService.UserId,
	}

	// 初始化响应状态码变量值
	code := e.Success

	// 获取userDao对象
	msgDao := dao.NewMsgDao(ctx)

	err := msgDao.SaveMsg(&msg)
	if err != nil {
		code = e.ErrorMsgSave
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "留言保存失败！",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (msgService MsgService) Get(ctx context.Context) serializer.Response {
	// 初始化响应状态码变量值
	code := e.Success
	// 获取userDao对象
	msgDao := dao.NewMsgDao(ctx)

	msgList, err := msgDao.GetMsg()
	if err != nil {
		code = e.ErrorMsgGet
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "留言查询失败！",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildMsgs(msgList),
	}
}

// ListMsg 分页查询
func (msgService MsgService) ListMsg(ctx context.Context) serializer.Response {
	// 初始化响应状态码变量值
	code := e.Success

	var msgs []*model.Message

	fmt.Printf("%v", msgService.BasePage)

	// 分页信息
	if msgService.BasePage.PageSize == 0 {
		msgService.BasePage.PageSize = 15
	}

	msgDao := dao.NewMsgDao(ctx)
	// 获取留言信息总数
	total, err := msgDao.CountMsg()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		msgDao := dao.NewMsgDaoByDB(msgDao.DB)
		msgs, _ = msgDao.MsgListByCondition(msgService.BasePage)
		wg.Done()
	}()
	wg.Wait()

	if err != nil || msgs == nil {
		code = e.ErrorMsgGet
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "留言分页查询失败！",
		}
	}

	// 返回
	return serializer.BuildListResponse(serializer.BuildMsgs(msgs), uint(total))

}

func (msgService MsgService) Delete(ctx context.Context) serializer.Response {
	// 初始化响应状态码变量值
	code := e.Success
	// 获取userDao对象
	msgDao := dao.NewMsgDao(ctx)

	err := msgDao.DeleteMsg(msgService.ID, msgService.UserId)
	if err != nil {
		code = e.ErrorMsgDelete
		// 返回
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "留言删除失败！",
		}
	}

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
