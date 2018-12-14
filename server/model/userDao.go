package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// UserDao 结构体
type UserDao struct {
	pool *redis.Pool
}

//MyUserDao model包里定义了一个UserDao的声明
var MyUserDao *UserDao

// NewUserDao 工厂模式创建一个userDao
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (userDao *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {
	str, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		} else {
			fmt.Println("getUserByID fail, err =", err)
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(str), user)
	if err != nil {
		fmt.Println("UserDao json.Unmarshal fail, err = ", err)
		return
	}
	return
}

//Login aa
func (userDao *UserDao) Login(userID int, userPWD string) (user *User, err error) {
	conn := userDao.pool.Get()
	defer conn.Close()
	user, err = userDao.getUserByID(conn, userID)
	if err != nil {
		return
	}
	if user.UserPWD != userPWD {
		err = ERROR_USER_PWD
		return
	}
	return
}
