package main

import (
    "fmt"
    "log"
	"os"
    "net/http"
    "github.com/gorilla/mux"
	"database/sql"	
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

func connect() *sql.DB {
	// change with your db configurastion
	var db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/restdb")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}

func getJSON(sqlString string) (string, error) {
    var db = connect()
	defer db.Close() 
	rows, err := db.Query(sqlString)
	if err != nil {
	  return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
	  return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
	  for i := 0; i < count; i++ {
		  valuePtrs[i] = &values[i]
	  }
	  rows.Scan(valuePtrs...)
	  entry := make(map[string]interface{})
	  for i, col := range columns {
		  var v interface{}
		  val := values[i]
		  b, ok := val.([]byte)
		  if ok {
			  v = string(b)
		  } else {
			  v = val
		  }
		  entry[col] = v
	  }
	  tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
	  return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil 
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/employee", EmployeeIndex)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome Rest!")
}

func EmployeeIndex(w http.ResponseWriter, r *http.Request) {
    data , err := getJSON("SELECT * FROM employee")
	if err != nil {return;}
	fmt.Fprintln(w, data)
}