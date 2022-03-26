package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.82.129:3306)/mylearngo"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createQuanzi(conn *gorm.DB) error {
	resp := conn.Create(&Quanzi{
		Name:    "LuLu",
		Sex:     "男",
		Tall:    1.70,
		Weight:  71,
		Age:     35,
		Fatrate: 2.5,
		Message: "今天心情不错",
	})
	if err := resp.Error; err != nil {
		fmt.Println("圈子创建失败：", err)
		return err
	}
	fmt.Println("圈子创建成功")
	return nil
}

func selectQuanzi(conn *gorm.DB) {
	output := make([]*Quanzi, 0)
	resp := conn.Where(&Quanzi{Name: "LuLu"}).Find(&output)
	if resp.Error != nil {
		fmt.Println("圈子查询失败：", resp.Error)
		return
	}

	data, _ := json.Marshal(output)
	fmt.Printf("圈子查询结果：%+v\n", string(data))
}

func updateQuanzi(conn *gorm.DB) error {
	p := &Quanzi{
		Id:      1,
		Name:    "LuLu",
		Sex:     "男",
		Tall:    1.70,
		Weight:  71,
		Age:     35,
		Fatrate: 2.9,
		Message: "今天天气不错",
	}
	resp := conn.Model(p).Select("name", "message").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Println("圈子更新失败：", err)
		return err
	}
	fmt.Println("圈子更新成功")
	return nil
}

func deleteQuanzi(conn *gorm.DB) {
	p := &Quanzi{
		Id: 1,
	}
	resp := conn.Delete(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除圈子失败：", err)
		return
	}
	fmt.Println("删除圈子成功")
}

func main() {
	//连接数据库
	conn := connectDb()
	fmt.Println(conn)

	//创建圈子
	//createQuanzi(conn)

	//查询圈子
	selectQuanzi(conn)

	//修改圈子
	//updateQuanzi(conn)

	//删除圈子
	//deleteQuanzi(conn)
}
