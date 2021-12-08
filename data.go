package main

import (
  "fmt"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
  id int16 `json:"id"`
  name string `json:"name"`
  age uint16 `json:"age"`
}

//проверка на ошибки
func checkError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") //подключение к базе
  checkError(err)

  defer db.Close() //закрытие подключения к базе

  //Установка данных
  // insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Sho', 53)")
  // checkError(err)
  //
  // defer insert.Close()

  //Выборка данных
  res, err := db.Query("SELECT * FROM `users`")
  checkError(err)

  for res.Next() {
    var user users
    err = res.Scan(&user.id, &user.name, &user.age)

    checkError(err)

    fmt.Println(fmt.Sprintf("User: %s with age %d and with id equal %d", user.name, user.age, user.id))
  }

  defer res.Close()

  fmt.Println("Подключено к MySQL")
}
