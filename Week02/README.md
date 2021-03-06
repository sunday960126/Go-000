#####我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

> 在dao层出现错误后，应该 Wrap 这个 error，抛给上层。

### 代码

1. homework.go

```go

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
```

   

