package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Database 连接和关闭已由db.go实现
type Database interface {
	//Connect() (*sql.DB, error)
	GetUserByStudID(studID int) (*User, error)
	AddUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(studID int) error
	//Close() error
}

type MysqlDatabase struct {
	db *sql.DB
}

// Connect 连接数据库
//func (m MysqlDatabase) Connect() (*sql.DB, error) {
//	dbConn := "root:123456@tcp(127.0.0.1:3306)/web"
//	db, err := sql.Open("mysql", dbConn)
//	if err != nil {
//		return nil, err
//	}
//	err = db.Ping()
//	if err != nil {
//		_ = m.Close()
//		return nil, err
//	}
//
//	return m.db, nil
//}

// GetUserByStudID 通过studID查询
func (m MysqlDatabase) GetUserByStudID(studID int) (*User, error) {
	query := "SELECT id, studId, username, sex, email FROM stud WHERE studId = ?"
	row := m.db.QueryRow(query, studID)

	var user User
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
func (m MysqlDatabase) AddUser(user *User) error {
	query := "INSERT INTO stud (studNo, username, sex, email) VALUES (?,?, ?, ?, ?)"
	_, err := m.db.Exec(query, user.StudID, user.Username, user.Sex, user.Email)
	if err != nil {
		return err
	}
	fmt.Println(*user)
	return err
}

// UpdateUser 更新用户信息
func (m MysqlDatabase) UpdateUser(user *User) error {
	query := "UPDATE stud SET username = ?, sex = ?, email = ? WHERE studNo = ?"
	_, err := m.db.Exec(query, user.Username, user.Sex, user.Email, user.StudID)
	return err
}

// DeleteUser 通过studID删除用户
func (m MysqlDatabase) DeleteUser(studID int) error {
	query := "DELETE FROM stud WHERE studNo = ?"
	_, err := m.db.Exec(query, studID)
	return err
}

//// Close 关闭数据库连接
//func (m MysqlDatabase) Close() error {
//	if m.db != nil {
//		return m.db.Close()
//	}
//	return nil
//}
