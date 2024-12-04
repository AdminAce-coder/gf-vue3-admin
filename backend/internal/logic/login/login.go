package login

import (
	"context"
	v1 "gf-vue3-admin/api/auth/v1"
	logcode "gf-vue3-admin/internal/consts/code/login"
	modecode "gf-vue3-admin/internal/consts/code/mode"
	"gf-vue3-admin/internal/dao"
	"gf-vue3-admin/internal/model/do"
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/dataprocess"
	"gf-vue3-admin/utility/docode"
	"gf-vue3-admin/utility/jwt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(new())
}

func new() *sLogin {
	return &sLogin{}
}

// 注册，写入数据库
func (s *sLogin) Register(ctx context.Context, req *v1.RegisterReq) error {
	// 查询用户
	ture, err := s.IsUser(req.Username)
	if err != nil {
		return err
	}
	if ture {
		return docode.NewError(modecode.MODE_SQLFailed, "用户已存在，请修改用户名")
	}
	glog.New().Info(ctx, "用户名不存在，正在创建...")

	// 加密密码
	req.Password = dataprocess.AesEcbEncrypt(ctx, req.Password)
	glog.New().Infof(ctx, "加密后的密码是...%s", req.Password)
	// 写入数据库,开启事务
	err = dao.User.DB().Model("user").Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.User.DB().Model("user").Data(do.User{
			UserName: req.Username,
			Password: req.Password,
		}).Insert()
		return err
	})
	if err != nil {
		return docode.NewError(modecode.MODE_SQLFailed, err.Error())
	}
	return nil
}

// 登录

func (s *sLogin) Lonin(ctx context.Context, req *v1.LoginReq) (string, error) {
	// 查询用户是否存在
	true, err := s.IsUser(req.Username)
	if err != nil {
		return "", docode.NewError(logcode.LOGIN_NOREPEAT_ERROR, "用户不存在，请先创建用户")
	}
	if true {

		// 查询存储的密码
		var user do.User
		err := dao.User.DB().Model("user").
			Where(dao.User.Columns().UserName, req.Username).
			Scan(&user)
		if err != nil {
			return "", docode.NewError(modecode.MODE_SQLFailed, "查询用户密码失败")
		}
		// 解析密码
		enpasswd := dataprocess.AesEcbdecrypted(ctx, gconv.String(user.Password))

		// 比较密码是否匹配
		if req.Password != enpasswd {
			return "", docode.NewError(logcode.LOGIN_NOREPEAT_ERROR, "密码错误")
		}
		// 返回token
		token, err := jwt.MakeToken(req.Username, enpasswd)
		if err != nil {
			return "", docode.NewError(7, err.Error())
		}
		return token, nil
	}
	return "", nil
}

// 查询用户
func (s *sLogin) IsUser(username string) (bool, error) {
	num, err := dao.User.DB().Model("user").Where(dao.User.Columns().UserName, username).Count()
	if err != nil {
		return false, err
	}
	if num > 0 {
		return true, nil
	}
	return false, nil

}

// 查询密码是否正确

func (s *sLogin) IsPasswdCorrect(username string) (bool, error) {
	// 查询存储的密码
	var user do.User
	err := dao.User.DB().Model("user").
		Where(dao.User.Columns().UserName, username).
		Scan(&user)
	if err != nil {
		return false, docode.NewError(modecode.MODE_SQLFailed, "查询用户密码失败")
	}
	//
	return true, nil
}
