package routes

import (
	"log"
	"net/http"

	controller "../controller"
	"github.com/gorilla/mux" //路由分配 輕量.方便.穩定
)

type Route struct { //定義一個新的資料類別Route
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route //宣告變數routes賦予資料類別Route

func init() {
	log.Println("<router init>")
	register("POST", "/PlatformServer/add", controller.AddTodo, nil)
	register("POST", "/PlatformServer/insert", controller.Insert, nil)
	register("POST", "/PlatformServer/delete", controller.Delete, nil)
	register("POST", "/PlatformServer/update", controller.Update, nil)
	register("GET", "/PlatformServer/query/{id}", controller.GetQueryById, nil)
	register("GET", "/PlatformServer/query", controller.Query, nil)

}

//初始化路由
func NewRouter() *mux.Router {
	log.Println("<router NewRouter func>")
	r := mux.NewRouter().StrictSlash(true) //.StrictSlash(true)可以允許後面多一個"/"
	for _, route := range routes {
		r.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)
		log.Printf("Method:%v, Pattern:%v, Handler:%v", route.Method, route.Pattern, route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}

	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	log.Println("<router register func>")
	log.Print("<register>", method, pattern, handler, middleware)
	routes = append(routes, Route{method, pattern, handler, middleware})
}
