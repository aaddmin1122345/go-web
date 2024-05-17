package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-web/model"
)

// Database 连接和关闭已由db.go实现
type Database interface {
	GetUserByUserName(username string) ([]*model.User, error)
	GetUserByStudID(studID int) (*model.User, error)
	AddUser(user *model.Register) error
	UpdateUser(user *model.User) error
	DeleteUser(studID int) error
	Login(login *model.Login) (*model.Login, error)
}

type MyDatabaseImpl struct {
	Db *sql.DB
}

// GetUserByUserName 通过用户名进行模糊查询
func (m MyDatabaseImpl) GetUserByUserName(username string) ([]*model.User, error) {
	query := "SELECT ID, StudID, Username, Sex, Email FROM stud WHERE Username LIKE ?"
	rows, err := m.Db.Query(query, "%"+username+"%")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.StudID, &user.Username, &user.Sex, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// GetUserByStudID 通过studID查询
func (m MyDatabaseImpl) GetUserByStudID(studID int) (*model.User, error) {
	//query := "SELECT id, studId, username, sex, email FROM stud WHERE studId = ?"
	query := "SELECT ID, StudID, Username, Sex, Email FROM stud WHERE StudID = ?"
	row := m.Db.QueryRow(query, studID)
	var user model.User
	err := row.Scan(&user.ID, &user.StudID, &user.Username, &user.Sex, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// AddUser 添加用户
func (m MyDatabaseImpl) AddUser(user *model.Register) error {
	query := "INSERT INTO stud (StudID, Username, Sex, Email,Password) VALUES (?,?, ?, ?, ?)"
	_, err := m.Db.Exec(query, user.StudID, user.Username, user.Sex, user.Email, user.Password)
	if err != nil {
		return err
	}
	fmt.Println(*user)
	return err
}

// UpdateUser 更新用户信息
func (m MyDatabaseImpl) UpdateUser(user *model.Register) error {
	query := "UPDATE stud SET StudID = ?, Username = ?, Sex = ?, Email = ? ,Password = ? ,WHERE ID = ?"
	_, err := m.Db.Exec(query, user.Username, user.Sex, user.Email, user.StudID)
	return err
}

// DeleteUser 通过studID删除用户
func (m MyDatabaseImpl) DeleteUser(studID int) error {
	query := "DELETE FROM stud WHERE ID = ?"
	_, err := m.Db.Exec(query, studID)
	return err
}

func (m MyDatabaseImpl) Login(login *model.Login) (*model.Login, error) {
	//query := "SELECT Username, Password FROM stud WHERE " + ""

	query := "SELECT Username, Password FROM stud WHERE Username = ?"

	// 执行查询
	var storedUsername, storedPassword string
	err := m.Db.QueryRow(query, login.Username).Scan(&storedUsername, &storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	//fmt.Println(login.Username, login.Password)

	if login.Password != storedPassword {
		return nil, errors.New("用户名或密码错误")
	}

	// 返回登录成功的用户信息
	return &model.Login{
		Username: storedUsername,
		Password: storedPassword,
	}, nil
}
