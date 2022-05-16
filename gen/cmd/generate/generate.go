package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func generateModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
	})

	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	//读取数据库中的users表，生成People结构题（这个只生成表，没得方法
	// g.GenerateModelAs("users", "People")
	//这个有方法
	g.ApplyBasic(g.GenerateModelAs("users", "Person"))
	g.Execute()
}
func main() {
	generateModel()
}
