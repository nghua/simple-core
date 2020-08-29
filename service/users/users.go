package users

import (
	"simple-core/database"
	"simple-core/graph/model"
	"simple-core/service/errmsg"
	"simple-core/service/token"
	"simple-core/service/validate"

	json "github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
)

type userEmail struct {
	Email string `validate:"required,email" label:"用户邮箱"`
}

type userPassword struct {
	Password string `validate:"required,gte=6" label:"密码"`
}

type userInfo struct {
	userEmail
	userPassword
}

// GetUserInfo 函数通过传入用户的id，获取cols字段中的所需要的数据。
func GetUserInfo(id int64, cols ...string) (*model.User, error) {
	u := &database.Users{Id: id}
	has, err := database.Engine().Cols(cols...).Get(u)
	if err != nil {
		return nil, errmsg.InternalError
	}

	if !has {
		return nil, errmsg.UserNotFoundError
	}

	um := &model.UserMeta{}
	if u.Meta != "" {
		err := json.Unmarshal([]byte(u.Meta), um)
		if err != nil {
			return nil, errmsg.InternalError
		}
	}

	user := &model.User{
		ID:    id,
		Email: u.Email,
		Role:  u.Role,
		Meta:  um,
	}

	return user, nil
}

// RegisterUser 函数为用户注册的具体逻辑实现
func RegisterUser(email, password string, meta *model.UserMeta) (bool, error) {
	ui := &userInfo{
		userEmail:    userEmail{Email: email},
		userPassword: userPassword{Password: password},
	}

	err := validate.Struct(ui)
	if err != nil {
		return false, err
	}

	u := &database.Users{Email: email}
	has, err := database.Engine().Exist(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if has {
		return false, errmsg.UserEmailExistError
	}

	pwBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, errmsg.InternalError
	}

	u = &database.Users{
		Email:    email,
		Password: string(pwBytes),
	}

	if meta != nil {
		metaBytes, err := json.Marshal(meta)
		if err != nil {
			return false, errmsg.InternalError
		}

		u.Meta = string(metaBytes)
	}

	_, err = database.Engine().InsertOne(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// UserLogin 函数为用户登录的具体实现
func UserLogin(email, password string) (string, error) {
	ui := &userInfo{
		userEmail:    userEmail{Email: email},
		userPassword: userPassword{Password: password},
	}

	err := validate.Struct(ui)
	if err != nil {
		return "", err
	}

	u := &database.Users{Email: email}
	has, err := database.Engine().Cols("id", "password").Get(u)
	if err != nil {
		return "", errmsg.InternalError
	}

	if !has {
		return "", errmsg.UserEmailOrPasswordWrongError
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", errmsg.UserEmailOrPasswordWrongError
	}

	userToke, err := token.Sign(u.Id, email)
	if err != nil {
		return "", errmsg.InternalError
	}
	return userToke, nil
}

// UpdateUserEmail 函数为用户更新邮箱的具体实现
func UpdateUserEmail(id int64, email string) (bool, error) {
	ue := &userEmail{Email: email}
	err := validate.Struct(ue)
	if err != nil {
		return false, err
	}
	u := &database.Users{Email: email}
	has, err := database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if has {
		return false, errmsg.UserEmailExistError
	}

	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// UpdateUserPassword 函数为用户更新密码的具体实现
func UpdateUserPassword(id int64, password string) (bool, error) {
	up := &userPassword{Password: password}
	err := validate.Struct(up)
	if err != nil {
		return false, err
	}
	u := &database.Users{}
	has, err := database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.UserNotFoundError
	}

	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// UpdateUserMeta 函数为用户跟新字段的具体实现
func UpdateUserMeta(id int64, meta *model.UserMeta) (bool, error) {
	u := &database.Users{}
	has, err := database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.UserNotFoundError
	}
	um, err := json.MarshalToString(meta)
	if err != nil {
		return false, errmsg.InternalError
	}

	u.Meta = um
	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// InsertUser 为后台插入用户的具体实现
func InsertUser(email, password string, role int, meta *model.UserMeta) (bool, error) {
	ui := &userInfo{
		userEmail:    userEmail{Email: email},
		userPassword: userPassword{Password: password},
	}

	err := validate.Struct(ui)
	if err != nil {
		return false, err
	}

	um, err := json.MarshalToString(meta)
	if err != nil {
		return false, errmsg.InternalError
	}

	u := &database.Users{
		Email:    email,
		Password: password,
		Role:     role,
		Meta:     um,
	}

	_, err = database.Engine().InsertOne(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// AlterUserInfo 为后台更改用户信息的具体实现
func AlterUserInfo(id int64, email, password string, role int, meta *model.UserMeta) (bool, error) {
	u := &database.Users{}
	has, err := database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.UserNotFoundError
	}

	if email != "" {
		ue := &userEmail{Email: email}
		err = validate.Struct(ue)
		if err != nil {
			return false, err
		}
		u.Email = email
	}

	if password != "" {
		up := &userPassword{Password: password}
		err = validate.Struct(up)
		if err != nil {
			return false, err
		}
		u.Password = password
	}

	if role != 0 {
		u.Role = role
	}

	if meta != nil {
		um, err := json.MarshalToString(meta)
		if err != nil {
			return false, errmsg.InternalError
		}
		u.Meta = um
	}

	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// DeleteUser 为后台删除用户的具体实现
func DeleteUser(id int64) (bool, error) {
	u := &database.Users{}
	_, err := database.Engine().ID(id).Delete(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}
