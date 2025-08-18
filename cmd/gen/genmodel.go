package main

// gorm gen configure
// go build cmd/gen/genmodel.go
// ref: https://www.liwenzhou.com/posts/Go/gen/
// ref: https://www.cnblogs.com/jeffid/articles/16701279.html

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"gorm.io/gen"
)

const MySQLDSN = "webstackgoai:webstackgoai@tcp(127.0.0.1:13306)/webstackgoai?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func WithInt11Annotation() gen.ModelOpt {
	return gen.FieldType("int", "int32")
}

func WithGormTags() gen.ModelOpt {

	return gen.FieldGORMTag("age", func(tag field.GormTag) field.GormTag {
		tag.Set("type", "int(11)")
		tag.Set("not null", "")
		return tag
	})

	// return gen.FieldGORMTag(func(field *schema.Field) (tag string) {
	// 	if field.DataType == "int" {
	// 		return `gorm:"type:int(11)"`
	// 	}
	// 	return ""
	// })
}

func IntTagConverter(tag field.GormTag) field.GormTag {
	tag.Set("type", "int(11)")
	tag.Set("not null", "")
	return tag
}

func main() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: "./internal/dal/query",
		/* ModelPkgPath: "dal/model"*/
		ModelPkgPath: "./internal/dal/model",
		// Repository:

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		FieldNullable:     false,
		FieldCoverable:    false,
		FieldWithIndexTag: true,
		FieldSignable:     true,
		FieldWithTypeTag:  true,
	})

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(connectDB(MySQLDSN))

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"int": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "*int"
			}
			return "int"
		},

		// bool mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "bool"
			}
			return "int8"
		},
		// datetime mapping
		"datetime": func(columnType gorm.ColumnType) (dataType string) {
			return "*time.Time"
		},
	}

	g.WithDataTypeMap(dataMap)

	// 从连接的数据库为所有表生成Model结构体和CRUD代码
	// 也可以手动指定需要生成代码的数据表
	// g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(
		g.GenerateModel("st_site",
			gen.FieldGORMTag("category_id", IntTagConverter),
			gen.FieldGORMTag("main_category_id", IntTagConverter),
			gen.FieldGORMTag("sort", IntTagConverter),
			gen.FieldGORMTag("view_count", IntTagConverter),
		),
		g.GenerateModel("st_category",
			gen.FieldGORMTag("parent_id", IntTagConverter),
			gen.FieldGORMTag("level", IntTagConverter),
			gen.FieldGORMTag("sort", IntTagConverter),
			gen.FieldGORMTag("count", IntTagConverter),
			gen.FieldGORMTag("free_count", IntTagConverter),
		),
	)

	// 执行并生成代码
	g.Execute()
}
