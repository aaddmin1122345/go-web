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
	GetUserByStudID(studID int) (*model.User, error)
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(studID int) error
}

type MyDatabaseImpl struct {
	Db *sql.DB
}

// GetUserByStudID 通过studID查询
func (m *MyDatabaseImpl) GetUserByStudID(studID int) (*model.User, error) {
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
func (m *MyDatabaseImpl) AddUser(user *model.User) error {
	query := "INSERT INTO stud (studNo, username, sex, email) VALUES (?,?, ?, ?, ?)"
	_, err := m.Db.Exec(query, user.StudID, user.Username, user.Sex, user.Email)
	if err != nil {
		return err
	}
	fmt.Println(*user)
	return err
}

// UpdateUser 更新用户信息
func (m *MyDatabaseImpl) UpdateUser(user *model.User) error {
	query := "UPDATE stud SET username = ?, sex = ?, email = ? WHERE studNo = ?"
	_, err := m.Db.Exec(query, user.Username, user.Sex, user.Email, user.StudID)
	return err
}

// DeleteUser 通过studID删除用户
func (m *MyDatabaseImpl) DeleteUser(studID int) error {
	query := "DELETE FROM stud WHERE studNo = ?"
	_, err := m.Db.Exec(query, studID)
	return err
}
