package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"

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

func (userDao *UserDao) getUserByID(conn redis.Conn, id int) (user *message.User, err error) {
	str, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		} else {
			fmt.Println("getUserByID fail, err =", err)
		}
		return
	}
	user = &message.User{}
	err = json.Unmarshal([]byte(str), user)
	if err != nil {
		fmt.Println("UserDao json.Unmarshal fail, err = ", err)
		return
	}
	return
}

//Login model包里userDao的Login函数，用于和redis进行校验
func (userDao *UserDao) Login(userID int, userPWD string) (user *message.User, err error) {
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

//Register model包里userDao的Register函数，用于和redis进行校验
func (userDao *UserDao) Register(user *message.User) (err error) {
	conn := userDao.pool.Get()
	defer conn.Close()
	_, err = userDao.getUserByID(conn, user.UserID)
	//
	if err != nil {
		if err == ERROR_USER_NOTEXISTS {
			data, err := json.Marshal(user)
			if err != nil {
				fmt.Println("userDao里的Register函数序列化user失败，err =", err)
				return err
			}
			_, err = conn.Do("hset", "users", user.UserID, string(data))
			if err != nil {
				fmt.Println("注册里的写入redis出错，err =", err)
				return err
			}
		}
	} else {
		err = ERROR_USER_EXISTS
	}
	return
}
