package services

import (
	"errors"
	"fmt"
	"net/smtp"
	"strconv"
	"time"

	"math/rand"

	"github.com/jordan-wright/email"
	"github.com/r1005410078/meida-admin-server/internal/app/repository"
	"github.com/r1005410078/meida-admin-server/internal/domain/user/events"
	"github.com/r1005410078/meida-admin-server/internal/infrastructure/dao/model"
	"go.uber.org/zap"
)

type UserServices struct {
	repo repository.IUserRepository
	logger *zap.Logger
}

func NewUserServices(repo repository.IUserRepository, logger *zap.Logger) *UserServices {
	return &UserServices{
		repo,
		logger,
	}
}

func (u *UserServices) FindById(userId string) (*model.User, error) {
	return u.repo.FindById(userId)
}

func (u *UserServices) List() ([]model.User, error) {
	return u.repo.List()
}

// 关联角色
func (u *UserServices) AssoicatedRolesEventHandle(event *events.AssoicatedRolesEvent) error {
	return u.repo.AssoicatedRoles(event)
}

// 关联角色失败
func (u *UserServices) AssoicatedRolesFailedEventHandle(event *events.AssoicatedRolesFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 删除用户
func (u *UserServices) DeleteUserHandle(event *events.UserDeletedEvent) error {
	return u.repo.DeleteUser(event)
}

// 删除用户失败	
func (u *UserServices) DeleteUserFailedEventHandle(event *events.UserDeleteFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 保存用户
func (u *UserServices) SaveUserEventHandle(event *events.SaveUserEvent) error {
	return u.repo.SaveUser(event)
}

// 保存用户失败
func (u *UserServices) SaveUserFailedEventHandle(event *events.SaveUserFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 更改用户状态
func (u *UserServices) SaveUserStatusEventHandle(event *events.UserStatusEvent) error {
	return u.repo.SaveUserStatus(event)
}

// 更改用户状态失败
func (u *UserServices) SaveUserStatusFailedEventHandle(event *events.UserStatusFailedEvent) error {
	u.logger.Error(event.Err.Error())
	return event.Err
}

// 根据邮箱获取验证码
func (u *UserServices) FindUserByEmail(email string) (*model.User, error) {
	return u.repo.FindUserByEmail(email)
}

// 给邮箱发送验证码
func (u *UserServices) SendEmailCode(username, input_email, code string) (*string, error) {
	if input_email == "" {
		return nil, errors.New("email is empty")
	}

	emailCode := generateCaptcha(6)
	targetMailBox := "1005410788@qq.com"          // 目标邮箱
	smtpServer := "smtp.163.com" 									// smtp服务器
	emailAddr := "rongtaosheng88@163.com"         // 要发件的邮箱地址
	smtpKey := "DHUfbbvFd97D5PnU"                 // 获取的smtp密钥

	em := email.NewEmail()
	em.From = fmt.Sprintf("Go-Cloud-Disk <%s>", emailAddr) // 发件人
	em.To = []string{targetMailBox}                        // 目标邮箱

	// email title
	em.Subject = "Email Confirm Test" // 标题
	// build email content
	em.Text = []byte(emailCode) // 内容

	// 调用接口发送邮件
  // 此处端口号不一定为25使用对应邮箱时需要具体更换
	em.Send(smtpServer+":25", smtp.PlainAuth("", emailAddr, smtpKey, smtpServer))

	// 保存验证码
  if err := u.repo.SaveEmailCode(input_email, emailCode); err != nil {
		return nil, err
	}

	return &emailCode, nil
}

// 登陆成功事件
func (u *UserServices) LoginSuccessEventHandle(event *events.LoggedInEvent) error {
	if err := u.repo.SaveLoginToken(*event.ID, generateCaptcha(32)); err != nil {
		return err
	}

	u.logger.Info(fmt.Sprintf("user %s login success", *event.Username))	
	return nil
}

// 登陆失败
func (u *UserServices) LoginFailedEventHandle(event *events.LoginFailedEvent) error {
	u.logger.Error(fmt.Sprintf("user %s login failed %s", event.Username, event.Error))
	return nil
}

// 退出登录
func (u *UserServices) LogoutEventHandle(event *events.LoggedOutEvent) error {
	if err := u.repo.DeleteLoginToken(&event.UserId); err != nil {
		return err
	}
	return nil
}

// 退出登录失败
func (u *UserServices) LogoutFailedEventHandle(event *events.LoggedOutFailedEvent) error {
	u.logger.Error(fmt.Sprintf("user %s logout failed %s", event.UserId, event.Err))
	return nil
}

// 注册成功
func (u *UserServices) RegisterCommandEventHandle(event *events.RegisteredEvent) error {
	u.logger.Info(fmt.Sprintf("user %s register success", event.Username))
	if err := u.repo.SaveUser(&events.SaveUserEvent {
		ID:        &event.ID,
		Username:  &event.Username,
		Email:     &event.Email,
	}); err != nil {
		return err
	}
	return nil
}

// 注册失败
func (u *UserServices) RegisterFailedCommandEventHandle(event *events.RegisterFailedEvent) error {
	u.logger.Error(fmt.Sprintf("user %s register failed %s", event.Username, event.Error))
	return nil
}

func generateCaptcha(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))// 设置随机数种子
	var captcha string
	for i := 0; i < length; i++ {
			captcha += strconv.Itoa(rand.Intn(10)) // 随机生成数字
	}
	return captcha
}