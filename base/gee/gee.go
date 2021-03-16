package gee

import (
	"fmt"
	"net/http"
)

/* 1：提供路由映射表 */
/* 2：提供用户注册静态路由办法 */
/* 3：重写了ServeHTTP方法 用于监听响应 */

type httpHandler func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]httpHandler //静态路由表
}

//New一张新的路由表
func New() *Engine {
	return &Engine{router: make(map[string]httpHandler)}
}

//路由表条目统一添加接口
func (e *Engine) addRoute(method string, pattern string, httpFunc httpHandler) {
	key := method + "-" + pattern
	e.router[key] = httpFunc
}

func (e *Engine) GET(pattern string, httpFunc httpHandler) {
	e.addRoute("GET", pattern, httpFunc)
}

func (e *Engine) POST(pattern string, httpFunc httpHandler) {
	e.addRoute("POST", pattern, httpFunc)
}

//RUN方法
func (e *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

//基于Engine实现ServeHTTP方法 实现自定义的响应操作
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if  httpFunc, ok := e.router[key]; ok {
		httpFunc(w, req);
	} else {
		fmt.Fprintf(w, "404 not found:%s\n", req.URL)
	}
}


