package gen

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

// GenEntity 生成数据库表实体 @param output 生成实体存储路径 @db gorm的数据库连接实例
func GenEntity(output string, db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: output,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // reuse your gorm db

	// 为结构模型生成基本的类型安全DAO API
	//g.ApplyBasic(models.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	g.GenerateAllTable()

	// Generate the code
	g.Execute()
}
