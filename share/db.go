package share

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDB() {

	config := viper.New()
	config.SetConfigName("db")
	config.AddConfigPath(Path.Config)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		config.GetString("database.user"),
		config.GetString("database.pwd"),
		config.GetString("database.host"),
		config.GetString("database.dbname"),
		true,
		// "Asia%2FShanghai",  // 必须是 url.QueryEscape 的
		"Local",
	))

	if err != nil {
		//log.Fatalf("数据库连接失败. 数据库名字: %s. 错误信息: %s", name, err)
		fmt.Println(err)
	} else {
		//配置
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return config.GetString("database.pre") + defaultTableName
		}
		db.SingularTable(true)
	}
}

func GetDB() *gorm.DB {
	return db
}
