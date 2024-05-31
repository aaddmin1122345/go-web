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
type Database interface {
	HashPassword(password string) ([]byte, error)
	CheckUser(user *model.Register) error
	CheckDelUser(id int) error
	GetUserByKeyword(username string) ([]*model.User, error)
	GetUserByPhoneNum(phoneNum string) (*model.User, error)
	AddUser(user *model.Register) error
	UpdateUser(user *model.Register) error
	DeleteUser(id int) error
	Login(login *model.Login) (*model.Login, error)
	SetDb(db *sql.DB)
}

type MyDatabaseImpl struct {
	db *sql.DB
}

func (m *MyDatabaseImpl) SetDb(db *sql.DB) {
	m.db = db
}

func (m *MyDatabaseImpl) GetUserByKeyword(username string) ([]*model.User, error) {
	//db, err := m.db.Conn()
	//if err != nil {
	//	return nil, err
	//}
	//defer db.Close()

	if m.db == nil {
		fmt.Println("数据库为空")

		return nil, nil
	}
	query := "SELECT ID, PhoneNum, Username, Sex, Email, Password,UserType,CreateTime FROM user WHERE PhoneNum LIKE ? OR Username LIKE ? OR Sex LIKE ? OR Email LIKE ? OR Password LIKE ? OR UserType LIKE ? OR CreateTime LIKE ?;"
	rows, err := m.db.Query(query, "%"+username+"%", "%"+username+"%", "%"+username+"%", "%"+username+"%", "%"+username+"%", "%"+username+"%", "%"+username+"%")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			// handle error
		}
	}(rows)

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.PhoneNum, &user.Username, &user.Sex, &user.Email, &user.Password, &user.UserType, &user.CreateTime)
		if err != nil {
			return nil, err
		}

		//fmt.Println(user)
		//fmt.Println(&user)
		//fmt.Printf("%T\n", user)
		//fmt.Printf("%T\n", &user)
		users = append(users, &user)
	}

	return users, nil
}

func (m *MyDatabaseImpl) CheckUser(user *model.Register) error {
	// 定义用户类型和性别的数组
	userTypeArray := [...]string{"reader", "author"}
	sexArray := [...]string{"男", "女", "其它", "不便于透露"}

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
	query := "DELETE FROM user WHERE ID = ?"
	_, err := m.db.Exec(query, id)
	return err
}

func (m *MyDatabaseImpl) Login(login *model.Login) (*model.Login, error) {
	query := "SELECT Username, Password,UserType FROM user WHERE Username = ? OR PhoneNum = ? OR Email = ?"

	// 执行查询
	var storedUsername, storedPassword, storedUserType string
	//sql用了 or 语句,所以两个参数可以设置为Username,Username没匹配到会去匹配PhoneNum,缺点是用户不存在会去查询两次
	err := m.db.QueryRow(query, login.Username, login.Username, login.Username).Scan(&storedUsername, &storedPassword, &storedUserType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("用户名或手机号或邮箱不存在或密码不对")
		}
		return nil, err
	}

	// 验证密码是否匹配,只能使用这个库自带的语句进行验证
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(login.Password))
	if err != nil {
		return nil, errors.New("用户名或手机号或邮箱不存在或密码不对")
	}

	// 便于测试返回登陆的用户信息
	return &model.Login{
		Username: storedUsername,
		Password: "none",
		UserType: storedUserType,
	}, nil

}
