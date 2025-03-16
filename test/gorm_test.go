package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// User 定义一个用户模型
type User struct {
	gorm.Model
	Name  string
	Age   int
	Email string
}

// 初始化数据库连接
func setupDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动迁移模型，创建对应的数据库表
	db.AutoMigrate(&User{})
	return db, nil
}

// TestCreateUser 测试创建用户
func TestCreateUser(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	result := db.Create(&user)
	if result.Error != nil {
		t.Errorf("Failed to create user: %v", result.Error)
	}

	if user.ID == 0 {
		t.Error("User ID should not be 0 after creation")
	}
}

// TestGetUser 测试查询用户
func TestGetUser(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user := User{
		Name:  "Jane Smith",
		Age:   25,
		Email: "janesmith@example.com",
	}
	db.Create(&user)

	var retrievedUser User
	result := db.First(&retrievedUser, user.ID)
	if result.Error != nil {
		t.Errorf("Failed to retrieve user: %v", result.Error)
	}

	if retrievedUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, retrievedUser.Name)
	}
}

// TestUpdateUser 测试更新用户信息
func TestUpdateUser(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user := User{
		Name:  "Bob Johnson",
		Age:   40,
		Email: "bobjohnson@example.com",
	}
	db.Create(&user)

	newName := "Updated Name"
	result := db.Model(&user).Update("Name", newName)
	if result.Error != nil {
		t.Errorf("Failed to update user: %v", result.Error)
	}

	var updatedUser User
	db.First(&updatedUser, user.ID)
	if updatedUser.Name != newName {
		t.Errorf("Expected name %s, got %s", newName, updatedUser.Name)
	}
}

// TestDeleteUser 测试删除用户
func TestDeleteUser(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user := User{
		Name:  "Delete Me",
		Age:   50,
		Email: "deleteme@example.com",
	}
	db.Create(&user)

	result := db.Delete(&user)
	if result.Error != nil {
		t.Errorf("Failed to delete user: %v", result.Error)
	}

	var deletedUser User
	result = db.First(&deletedUser, user.ID)
	if result.Error == nil {
		t.Error("User should not be found after deletion")
	}
}
