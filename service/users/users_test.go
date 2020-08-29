package users

import (
	"os"
	"simple-core/graph/model"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

var (
	email       = "wanyamodi@163.com"
	password    = "123456"
	nickname    = "滑滑"
	newNickname = "哈哈哈哈"
)

func TestRegisterUser(t *testing.T) {
	_, err := RegisterUser(email, password, &model.UserMeta{Nickname: &nickname})
	if err != nil {
		t.Errorf("注册出错：%v", err)
	}
}

func TestUserLogin(t *testing.T) {
	token, err := UserLogin(email, password)
	if err != nil {
		t.Errorf("登录出错：%v", err)
	}

	t.Logf("用户的token： %s", token)
}

func TestGetUserInfo(t *testing.T) {
	user, err := GetUserInfo(1, "email")
	if err != nil {
		t.Errorf("获取用户信息出错：%v", err)
	}
	t.Logf("用户信息：%#v", user)
}

func TestUpdateUserEmail(t *testing.T) {
	_, err := UpdateUserEmail(1, "wanyamodi@outlook.com")
	if err != nil {
		t.Errorf("更新用户邮箱出错：%v", err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	_, err := UpdateUserPassword(1, "2345678")
	if err != nil {
		t.Errorf("更新用户密码出错：%v", err)
	}
}

func TestUpdateUserMeta(t *testing.T) {
	_, err := UpdateUserMeta(1, &model.UserMeta{Nickname: &newNickname})
	if err != nil {
		t.Errorf("更新用户字段出错：%v", err)
	}
}

func TestInsertUser(t *testing.T) {
	_, err := InsertUser("abc@qwrr.com", "123456", 1, &model.UserMeta{Nickname: &newNickname})
	if err != nil {
		t.Errorf("插入用户出错：%v", err)
	}
}

func TestAlterUserInfo(t *testing.T) {
	_, err := AlterUserInfo(2, "", "", 4, nil)
	if err != nil {
		t.Errorf("更改用户信息出错：%v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	_, err := DeleteUser(2)
	if err != nil {
		t.Errorf("删除用户出错：%v", err)
	}
}
