package user

import (
	"fmt"
	"main/app/model"
	"main/share"
	"main/share/entity"
	"main/share/tool/unique"
	"main/unit"
)

func Result(b *unit.Context) string {
	return entity.Result(1, "abc", nil)
}

func addUser(b *unit.Context) string {

	user := model.User{
		NickName: "dakesolo",
		TrueName: "张老虎",
		UserName: "qissen",
	}
	//database.DB.Self.AutoMigrate(&User{})
	share.GetDB().Create(&user)
	fmt.Println("111")
	return "GetIndex"
}

func GetNav(b *unit.Context) string {
	//fmt.Println(share.Path.Config)
	tables, err := share.TableStoreClient.ListTable()
	if err != nil {
		fmt.Println("Failed to list table")
	} else {
		fmt.Println("List table result is")
		for _, table := range tables.TableNames {
			fmt.Println("TableName: ", table)
		}
	}

	return "getNav"
}

func GetBanner(context *unit.Context) string {
	/*share.Log.WithFields(logrus.Fields{
		"request": context.Request,
	}).Info("A group of walrus emerges from the ocean")*/
	//fmt.Println(unique.GetMD5())
	fmt.Println(unique.GetMD5())
	return "getBanner"
}
