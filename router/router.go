package router

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// ToDoService implement HTTP server
type Router struct {
	JWTKey string
	Conn   *gorm.DB
}

func (route *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		{
			route.getAuthToken(resp, req)
			return
		}
	case "/register":
		{
			switch req.Method {
			case http.MethodPost:
				route.registerUser(resp, req)
			}
		}
	}
}
