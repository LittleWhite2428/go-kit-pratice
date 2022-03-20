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
	router := mymux.NewRouter()
	{
		router.Handle(`/user/{uid:\d+}`,serviceHandler)
		router.Methods("GET").Path("/heath").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Contend-type","application")
			writer.Write([]byte(`{"status:OK"}`))
		})
	}

	http.ListenAndServe(":8080", router)

}
