package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"go-kit-exercise/service"
	"net/http"
)

func main() {
	user := service.UserService{}
	endp := service.GenUserEndpoint(user)

	serviceHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)
	http.ListenAndServe(":8080",serviceHandler)

}
