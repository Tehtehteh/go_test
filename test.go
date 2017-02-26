package main
import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"math/rand"
	"log"
)
const conn_string = "root:imonomy@/goblog?charset=utf8"
const driver_name = "mysql"
var db, top_error = sql.Open(driver_name, conn_string)

func indexPage(resp http.ResponseWriter, req *http.Request){
	log.Println(req.Method)
	t, err := template.ParseFiles("D:\\Golangprojects\\test\\src\\main\\home")
	var (
		id int
		title string
		text string
		posts []map[string]string
	)
	var post = make(map[string]string)
	if err != nil {
		panic("Something went wrong with template")
	}
	res, err := db.Query("SELECT id, title, text FROM blog_post")
	if err != nil {
		panic("Something went wrong with DB query")
	}
	defer res.Close()
	for res.Next(){
		err := res.Scan(&id, &title, &text)
		if err != nil { log.Fatal(err)}
		post["title"] = title
		post["text"] = text
		posts = append(posts, post)
	}
	err = res.Err()
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(resp, "home", posts )
}

func contactPage(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp, "You request is %s", req.Method)
}



func main() {
	if top_error != nil{
		panic("Error connecting to database")
	}
	var port = ":8000"
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/contact", contactPage)
	fmt.Printf("Server started, port %s", port)
	http.ListenAndServe(port, nil)
}
