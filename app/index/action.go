package index

import (
	"fmt"
	"main/share"
	"main/share/tool/unique"
	"main/unit"
)

func Hello(b *unit.Context) string {
	return "hello world"
}

func GetIndex(b *unit.Context) string {

	/*model := model.User{
		ExpressName: "99999999",
		OrderNumber: "123",
	}
	//database.DB.Self.AutoMigrate(&User{})
	share.GetDB().Create(&model)
	fmt.Println("111")*/
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
