package main

/* 基于Engine实现了ServeHTTP方法 */
/* 利用ResponseWriter指针可以构造对该请求的响应 */
/* 第一步：将所有请求转向了自己的处理逻辑 */

/*
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  switch req.URL.Path {
    case "/":
      fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    case "/Hello":
      for k,v := range req.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
      }
    default:
      fmt.Fprintf(w, "404 not found:%s\n", req.URL)
  }
}

func main() {
  engine := new(Engine)
  log.Fatal(http.ListenAndServe(":9999", engine))
}
*/

func main() {
  r := gee.New()
  r.GET("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
  })
  
  r.GET("/Hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
  })
  
  r.RUN(":9999")
}
