package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-web/model"
	"golang.org/x/crypto/bcrypt"
)

//var userTypeArray [...]string

//var sexArray [...]string

// Database 连接和关闭已由db.go实现

type MyDatabaseImpl struct {
	db *sql.DB
}

func (m *MyDatabaseImpl) SetDb(db *sql.DB) {
	m.db = db
}

func (m *MyDatabaseImpl) GetUserByKeyword(username string) ([]*model.User, error) {
	var query string
	var args []interface{}

	if username == "" {
		query = "SELECT ID, PhoneNum, Username, Sex, Email, Password, UserType, CreateTime, IsDelete FROM user WHERE IsDelete = 0;"
	} else {
		query = "SELECT ID, PhoneNum, Username, Sex, Email, Password, UserType, CreateTime, IsDelete FROM user WHERE PhoneNum LIKE ? OR Username LIKE ? OR Sex LIKE ? OR Email LIKE ? OR Password LIKE ? OR UserType LIKE ? OR CreateTime LIKE ? AND IsDelete = 0;"
		args = []interface{}{"%" + username + "%", "%" + username + "%", "%" + username + "%", "%" + username + "%", "%" + username + "%", "%" + username + "%", "%" + username + "%"}
	}

	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.PhoneNum, &user.Username, &user.Sex, &user.Email, &user.Password, &user.UserType, &user.CreateTime, &user.IsDelete)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (m *MyDatabaseImpl) CheckUser(user *model.Register) error {
	// 定义用户类型和性别的数组
	userTypeArray := [...]string{"reader", "author"}
	sexArray := [...]string{"男", "女", "其它", "不便透露"}

	// 检查用户类型是否有效
	userTypeValid := false
	for _, userType := range userTypeArray {
		if user.UserType == userType {
			userTypeValid = true
			break
		}
	}
	if !userTypeValid {
		return errors.New("无效的用户类型")
	}

	// 检查性别是否有效
	sexValid := false
	for _, sex := range sexArray {
		if user.Sex == sex {
			sexValid = true
			break
		}
	}
	if !sexValid {
		return errors.New("无效的性别")
	}

	return nil
}

func (m *MyDatabaseImpl) HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("生成密码失败:", err)
		return nil, err
	}

	// 打印密码哈希

	// 验证密码
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("密码验证失败:", err)
		return nil, err
	}

	fmt.Println("密码验证通过")
	return hashedPassword, err
}

// GetUserByKeyword 通过用户名进行模糊查询,使用指针就得用切片,不然只会返回一条数据

// GetUserByPhoneNum  通过studID查询
func (m *MyDatabaseImpl) GetUserByPhoneNum(phoneNum string) (*model.User, error) {
	//query := "SELECT id, studId, username, sex, email FROM stud WHERE studId = ?"
	query := "SELECT ID, PhoneNum, Username, Sex, Email FROM user WHERE PhoneNum = ?"
	row := m.db.QueryRow(query, phoneNum)
	var user model.User
	err := row.Scan(&user.ID, &user.PhoneNum, &user.Username, &user.Sex, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// AddUser 添加用户
func (m *MyDatabaseImpl) AddUser(user *model.Register) error {
	// 验证用户类型和性别是否有效
	//前边有if,后面的err就可以简写
	if err := m.CheckUser(user); err != nil {
		return err
	}

	// 生成密码哈希
	hashPassword, err := m.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// 执行插入操作
	query := "INSERT INTO user (PhoneNum, Username, Sex, Email, Password, UserType) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = m.db.Exec(query, user.PhoneNum, user.Username, user.Sex, user.Email, hashPassword, user.UserType)
	if err != nil {
		err = errors.New("邮箱/手机号/用户名被人使用了")
		return err
	}

	fmt.Println(*user)
	return nil
}

// UpdateUser 更新用户信息
func (m *MyDatabaseImpl) UpdateUser(user *model.Register) error {
	//前边有if,后面的err就可以简写
	if err := m.CheckUser(user); err != nil {
		return err
	}
	query := "UPDATE user SET PhoneNum = ?, Username = ?, Sex = ?, Email = ?, Password = ? WHERE ID = ?"
	hashPassword, _ := m.HashPassword(user.Password)
	_, err := m.db.Exec(query, user.PhoneNum, user.Username, user.Sex, user.Email, hashPassword, user.ID)
	return err
}

func (m *MyDatabaseImpl) CheckDelUser(id int) error {
	if id == 1 {
		return errors.New("无法删除默认用户")
	}
	return nil
}

// DeleteUser 通过studID删除用户
func (m *MyDatabaseImpl) DeleteUser(id int) error {
	//user := &model.Register{} // 创建一个 model.Register 的指针
	if err := m.CheckDelUser(id); err != nil {
		return err
	}
	query := "update user set IsDelete = 1 where ID = ?"
	_, err := m.db.Exec(query, id)
	return err
}

func (m *MyDatabaseImpl) Login(login *model.Login) (*model.Login, error) {
	// 检查密码长度
	if len(login.Password) < 6 {
		return nil, errors.New("密码太短了")
	}

	// 查询语句
	query := "SELECT ID, Username, Password, UserType FROM user WHERE (Username = ? OR PhoneNum = ? OR Email = ?) AND IsDelete = 0"

	// 执行查询
	var storedUsername, storedPassword, storedUserType string
	var storedUserID int
	err := m.db.QueryRow(query, login.Username, login.Username, login.Username).Scan(&storedUserID, &storedUsername, &storedPassword, &storedUserType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("用户或密码不正确")
		}
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(login.Password))
	if err != nil {
		return nil, errors.New("用户或密码不正确")
	}

	// 构造并返回登录信息
	return &model.Login{
		ID:       storedUserID,
		Username: storedUsername,
		Password: "*********", // 或者使用空字符串或其他占位符
		UserType: storedUserType,
	}, nil
}
