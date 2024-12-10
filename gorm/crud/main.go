package main

import (
	"fmt"
	"gorm.io/gorm"
	"note/gorm/domain"
	"note/gorm/initialize"
	"time"
)

var db *gorm.DB

func main() {
	db = initialize.InitMysql()
	//Create()

	//Select()

	//Update()

	//Delete()

	GenerateSQL()
}

func GenerateSQL() {
	type Result struct {
		ID   int
		Name string
		Age  int
	}

	// 用于执行 SELECT 查询，返回结果
	//var result Result
	//db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)

	// 用于执行不返回结果的 SQL 语句，如 INSERT、UPDATE 或 DELETE
	//db.Exec("DROP TABLE users")
	//db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

	// DryRun 模式 在不执行的情况下生成 SQL 及其参数
	var user domain.User
	statement := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	fmt.Println(statement.SQL.String())
	fmt.Println(statement.Vars)
	// ToSQL 生成sql 不执行
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&domain.User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]domain.User{})
	})

	fmt.Println(sql)
}

func Delete() {
	// 删除一条记录
	var user domain.User
	user.ID = 1
	db.Where("name = ?", "jinzhu").Delete(&user)

	// 根据主键删除
	db.Delete(&domain.User{}, 10)

	// gorm 会阻止全局删除
	_ = db.Delete(&domain.User{}).Error // gorm.ErrMissingWhereClause

	_ = db.Delete(&[]domain.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}}).Error // gorm.ErrMissingWhereClause
}

func Update() {
	var user domain.User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	// 会保存所有的字段，即使字段是零值
	db.Save(&user)

	// Save() 保存时不包含主键，他将执行 Create, 否者执行 Update
	db.Save(&domain.User{Name: "qwe", Age: 124})

	// 更新单个列，根据条件更新
	db.Model(&domain.User{}).Where("name = ?", "qwe").Update("age", gorm.Expr("age - 10"))

	db.Model(&domain.User{}).Update("age", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "users" SET "age" = age * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
}

func Select() {
	var user domain.User
	// 获取第一条记录（主键升序）
	//result := db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	//result := db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	//result := db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	//fmt.Printf("result: %v", result)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	//fmt.Printf("err record not found %v", errors.Is(result.Error, gorm.ErrRecordNotFound))

	//var users []domain.User

	// works because destination struct is passed in
	//db.Session(&gorm.Session{QueryFields: true}).First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// works because model is specified using `db.Model()`
	result := map[string]interface{}{}
	//db.Model(&domain.User{}).First(&result)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	//result := map[string]interface{}{}
	//tx := db.Table("users").First(&result)

	// works with Take
	//result := map[string]interface{}{}
	//tx := db.Session(&gorm.Session{QueryFields: true}).Table("users").Take(&result)
	tx := db.Model(&domain.User{}).Where("id > (?)", db.Model(&domain.User{}).Select("AVG(id)")).Find(&result)

	fmt.Printf("err: %v", tx.Error)

	db.Where(domain.User{Name: "Jinzhu"}).Assign(domain.User{Age: 10}).FirstOrInit(&user)
	fmt.Printf("user: %v", user)

}

func Create() {
	//err := db.AutoMigrate(&domain.User{})
	//if err != nil {
	//	panic(err)
	//}
	//user := domain.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)

	users := []*domain.User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	//err := db.CreateInBatches(&users, len(users)).Error
	err := db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(&users, len(users)).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
