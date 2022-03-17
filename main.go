package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"go-kit-exercise/service"
	mymux "github.com/gorilla/mux"
	"net/http"
)

func main() {
	user := service.UserService{}
	endp := service.GenUserEndpoint(user)

	serviceHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)
	r := mymux.NewRouter()
	r.Handle(`/user/{uid:\d+}`,serviceHandler)
	http.ListenAndServe(":8080",r)

}
