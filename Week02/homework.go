package Week02

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_errors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

var (
	ErrNotFound = &Error{
		Code: 1001,
		Msg:  "record not found",
	}
)

var db *gorm.DB

type User struct {
}

func Dao(id int) error {
	user := &User{}
	err := db.Table("t_user").Where("id = ?", id).Find(user).Error
	if err != nil {
		return _errors.Wrap(err, fmt.Sprintf("find user by id err: %v", id))
	}
	return nil
}

func Service() error {
	return Dao(0)
}

func Api() error {
	err := Service()
	if err != nil {
		logrus.Errorf("original error:%T %v\n", _errors.Cause(err), _errors.Cause(err))
		logrus.Errorf("stack trace:\n%+v\n", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}
