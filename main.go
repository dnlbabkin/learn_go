package main

import ("fmt"
        "net/http"
        "html/template")

//создание структуры
type User struct {
  Name string
  Age uint16 //целое не отрицательное число
  Money int16
  AvgGrades, Happiness float64
  Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d and he " +
    "has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
  u.Name = newName
}

func main() {
  //создание объектов структуры
  // var bob User =
  // bob := User{name: "Bob", age: 33, money: -50, avg_grades: 4.2, happiness: 0.8}

  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", homePage) //принимает 2 парамаетра. Позволяет отследить url и при переходе на него вызывает определенный метод
  http.HandleFunc("/contacts/", contactsPage)

  //создание локального сервера
  http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
  bob := User{"Bob", 33, -50, 4.2, 0.8, []string {"Football", "Skate", "Masturbating"}}
  // bob.name = "NeBob"
  // bob.setNewName("NeBob")
  // fmt.Fprintf(w, bob.getAllInfo())
  // fmt.Fprintf(w, `<h1>Main text</h1>
  //   <b>Not main text</b>`)
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "То шо какие-то контакты")
}
