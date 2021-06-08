package router

import (
	"encoding/json"
	"net/http"
	"skeleton/internal/services/transport"
)

func (route *Router) getAuthToken(resp http.ResponseWriter, req *http.Request) {

	var (
		method = "get-auth"
		err    error
	)

	data, err := transport.Init(route.Conn, method, req)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data["token"],
	})
}

func (route *Router) registerUser(resp http.ResponseWriter, req *http.Request) {
	var (
		method = "register"
		err    error
	)
	data, err := transport.Init(route.Conn, method, req)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data["token"],
	})
}
