package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	// gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DBマイグレート
func dbInit() {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず!（dbInit）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// DB追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// DB全取得
func dbGetAll() []Todo {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DB一つ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

// DB更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DB削除
func dbDelete(id int) {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// func userPage(w http.ResponseWriter, r *http.Request) {
// 	users := getUsers()

// 	fmt.Println("Endpoint Hit: usersPage")
// 	json.NewEncoder(w).Encode(users)
// }

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	// dbInit()

	//Index
	r.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	//Create
	r.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	r.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	//削除確認
	r.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	r.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")

	})

	r.Run(":8080")
}
