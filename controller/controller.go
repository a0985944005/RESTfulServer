package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	database "../database"
	"../services"
	"github.com/gorilla/mux"
)

type Todo struct {
	Id   int64
	Item string
}
type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

var TodoList []Todo
var TestList []database.Test

func Insert(w http.ResponseWriter, r *http.Request) {

	log.Println("<controller Test.Insert func>")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var addtest database.Test
	_ = json.Unmarshal(body, &addtest) //轉為json
	addtest.Time = time.Now().Format("2006-01-02 15:04:05")
	addtest.Insert()
	defer r.Body.Close()
	TestList = append(TestList, addtest)

	response := ApiResponse{"200", TestList}

	services.ResponseWithJson(w, http.StatusOK, response) //回傳
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.Delete func>")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var deltest database.Test
	_ = json.Unmarshal(body, &deltest) //轉為json
	defer r.Body.Close()
	deltest.Delete()
	response := ApiResponse{"200", "Delete sucess"}
	services.ResponseWithJson(w, http.StatusOK, response) //回傳
}

func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.Update func>")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var updatetest database.Test
	_ = json.Unmarshal(body, &updatetest) //轉為json
	defer r.Body.Close()
	updatetest.Update()
	response := ApiResponse{"200", "Update sucess"}
	services.ResponseWithJson(w, http.StatusOK, response) //回傳
}

func Query(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.Query func>")
	var querytest database.Test
	querydata := querytest.Query()

	defer r.Body.Close()
	response := ApiResponse{"200", querydata}

	services.ResponseWithJson(w, http.StatusOK, response) //回傳
}

func GetQueryById(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.GetQueryById func>")
	vars := mux.Vars(r)   //會將帶入的get內容轉乘map["id":"get內容"]
	queryId := vars["id"] //獲取url參數

	var querytest database.Test
	querydata := querytest.QueryById(queryId)

	defer r.Body.Close()
	response := ApiResponse{"200", querydata}

	services.ResponseWithJson(w, http.StatusOK, response) //回傳
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.AddTodo func>")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var addTodo Todo
	_ = json.Unmarshal(body, &addTodo) //轉為json
	log.Println(body)
	log.Println(&addTodo)
	TodoList = append(TodoList, addTodo)
	defer r.Body.Close()
	response := ApiResponse{"200", TodoList}

	services.ResponseWithJson(w, http.StatusOK, response) //回傳

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	log.Println("<controller Test.GetTodoById func>")
	vars := mux.Vars(r)   //會將帶入的get內容轉乘map["id":"get內容"]
	queryId := vars["id"] //獲取url參數

	var targetTodo Todo
	for _, Todo := range TodoList { //比對TodoList內是否有符合的Todo
		intQueryId, _ := strconv.ParseInt(queryId, 10, 64) //將string轉為int64
		if Todo.Id == intQueryId {
			targetTodo = Todo
		}
	}
	response := ApiResponse{"200", targetTodo}
	services.ResponseWithJson(w, http.StatusOK, response)

}
