package users

import (
	"simple-core/database"
	"simple-core/graph/model"
	"simple-core/service/errmsg"
	"simple-core/service/token"
	"simple-core/service/uuid"
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

// Get 函数通过传入用户的id，获取cols字段中的所需要的数据。
func Get(id int64, cols ...string) (*model.User, error) {
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

// Register 函数为用户注册的具体逻辑实现
func Register(email, password string, meta *model.UserMeta) (bool, error) {
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

	uid, err := uuid.GenerateUid()
	if err != nil {
		return false, errmsg.InternalError
	}

	u = &database.Users{
		Id:       uid,
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

// Login 函数为用户登录的具体实现
func Login(email, password string) (string, error) {
	ui := &userInfo{
		userEmail:    userEmail{Email: email},
		userPassword: userPassword{Password: password},
	}

	err := validate.Struct(ui)
	if err != nil {
		return "", err
	}

	u := &database.Users{Email: email}
	has, err := database.Engine().Cols("id", "password", "role").Get(u)
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

	userToke, err := token.Sign(u.Id, email, u.Role)
	if err != nil {
		return "", errmsg.InternalError
	}
	return userToke, nil
}

// UpdateEmail 函数为用户更新邮箱的具体实现
func UpdateEmail(id int64, email string) (bool, error) {
	ue := &userEmail{Email: email}
	err := validate.Struct(ue)
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

	u = &database.Users{Id: id}
	has, err = database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.UserNotFoundError
	}

	u = &database.Users{Email: email, Version: u.Version}

	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// UpdatePassword 函数为用户更新密码的具体实现
func UpdatePassword(id int64, password string) (bool, error) {
	up := &userPassword{Password: password}
	err := validate.Struct(up)
	if err != nil {
		return false, err
	}
	u := &database.Users{Id: id}
	has, err := database.Engine().Cols("version").Get(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.UserNotFoundError
	}

	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, errmsg.InternalError
	}

	u = &database.Users{Password: string(p), Version: u.Version}
	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// UpdateMeta 函数为用户跟新字段的具体实现
func UpdateMeta(id int64, meta *model.UserMeta) (bool, error) {
	u := &database.Users{Id: id}
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

	u = &database.Users{
		Meta:    um,
		Version: u.Version,
	}

	_, err = database.Engine().ID(id).Update(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// Insert 为后台插入用户的具体实现
func Insert(email, password string, role int, meta *model.UserMeta) (bool, error) {
	ui := &userInfo{
		userEmail:    userEmail{Email: email},
		userPassword: userPassword{Password: password},
	}

	err := validate.Struct(ui)
	if err != nil {
		return false, err
	}

	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, errmsg.InternalError
	}

	um := ""
	if meta != nil {
		um, err = json.MarshalToString(meta)
		if err != nil {
			return false, errmsg.InternalError
		}
	}

	uid, err := uuid.GenerateUid()
	if err != nil {
		return false, errmsg.InternalError
	}

	u := &database.Users{
		Id:       uid,
		Email:    email,
		Password: string(p),
		Role:     role,
		Meta:     um,
	}

	_, err = database.Engine().InsertOne(u)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

// Alter 为后台更改用户信息的具体实现
func Alter(id int64, email, password string, role int, meta *model.UserMeta) (bool, error) {
	u := &database.Users{}
	has, err := database.Engine().ID(id).Cols("version").Get(u)
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
		p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return false, errmsg.InternalError
		}
		u.Password = string(p)
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

// Delete 为后台删除用户的具体实现
func Delete(id int64) (bool, error) {
	_, err := database.Engine().ID(id).Delete(new(database.Users))
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

func GetList(offset, row int, cols ...string) ([]*model.User, error) {
	var usersList []database.Users
	err := database.Engine().Cols(cols...).Limit(row, offset).Find(&usersList)
	if err != nil {
		return nil, errmsg.InternalError
	}

	users := make([]*model.User, len(usersList))
	meta := &model.UserMeta{}

	for i, u := range usersList {
		if u.Meta != "" {
			err = json.UnmarshalFromString(u.Meta, meta)
			if err != nil {
				return nil, errmsg.InternalError
			}
		}

		users[i] = &model.User{
			ID:    u.Id,
			Email: u.Email,
			Role:  u.Role,
			Meta:  meta,
		}
	}

	return users, nil
}
