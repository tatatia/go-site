package main

//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type User struct {
//	 Name string `json:"name"`
//	 Age  uint16 `json:"age"`
//}
//
//func main() {
//	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mydb")
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	//установка даних
//	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 45)")
//	//if err != nil{
//	//	panic(err)
//	//}
//	//defer insert.Close()
//
//	//виборка даних
//	res, err := db.Query("SELECT `name`, `age` FROM `users`")
//	if err != nil {
//		panic(err)
//	}
//	for res.Next() {
//		var user User
//		err = res.Scan(&user.Name, &user.Age)
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
//	}
//}
