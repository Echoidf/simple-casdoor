package object

import "fmt"

func CheckUserPassword(username string, password string) (user *User, err error) {
	// 获取user的密码
	user = &User{}
	existed, err := adapter.Engine.Where("name=?", username).Get(user)
	if err != nil {
		return nil, err
	}

	if !existed {
		return nil, fmt.Errorf("该用户不存在")
	}

	if user.Password == password {
		return user, nil
	}

	return nil, fmt.Errorf("密码错误")
}
