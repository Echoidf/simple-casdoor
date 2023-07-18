package object

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Name      string    `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedAt time.Time `xorm:"varchar(100) index" json:"created_at"`

	UserId       string `xorm:"varchar(255)" json:"user_id"`
	Password     string `xorm:"varchar(100)" json:"password"`
	PasswordSalt string `xorm:"varchar(100)" json:"passwordSalt,omitempty"`
	PasswordType string `xorm:"varchar(100)" json:"passwordType"`
	Phone        string `xorm:"varchar(20) index" json:"phone,omitempty"`
	CountryCode  string `xorm:"varchar(6)" json:"countryCode"`
}

type Userinfo struct {
	Name  string `json:"preferred_username,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func AddUser(user *User) (int64, error) {
	var err error
	if user.UserId == "" {
		user.UserId = uuid.NewString()
	}

	affected, err := adapter.Engine.Insert(user)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func GetUserByPhone(phone string) (*User, error) {
	if phone == "" {
		return nil, nil
	}

	user := User{Phone: phone}
	existed, err := adapter.Engine.Get(&user)
	if err != nil {
		return nil, err
	}

	if existed {
		return &user, nil
	} else {
		return nil, nil
	}
}
