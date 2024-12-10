package main

import (
	"gorm.io/gorm"
	"note/gorm/initialize"
)

var db *gorm.DB

func main() {
	db = initialize.InitMysql()
	//db.AutoMigrate(&User{}, &CreditCard{})
	//user := User{
	//	CreditCards: []CreditCard{{
	//		Number: "123",
	//	}},
	//}
	//db.Create(&user)
	//db.Create(&CreditCard{UserID: 4, Number: "1234"})
	//var user User
	//db.Preload("CreditCards").Find(&user, 4)
	//fmt.Println(user)

	// many2many
	db.AutoMigrate(&User{}, &Language{})
	//db.Omit("languages").Create(&User{})
	users := []User{
		{Username: "wa"},
		{Username: "li"},
	}
	db.Create(&users)
	db.Model(&users[0]).Association("Languages").Replace([]Language{{Name: "cn"}, {Name: "en"}})
	// 关联删除
	//user.ID = 2
	//db.Select(clause.Associations).Delete(&user)
	// 追加关联
	//db.Model(&user).Association("CreditCards").Append([]CreditCard{{Number: "asdfg"}, {Number: "hj"}})

	// 替换关联
	//db.Model(&user).Association("CreditCards").Replace([]CreditCard{{Number: "12345"}, {Number: "7890"}})
}
