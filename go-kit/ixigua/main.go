package main

import (
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang-learn/go-kit/ixigua/service"
	"golang-learn/go-kit/ixigua/uitl"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	svc := service.UserServiceImpl{}
	end := service.MakeUserEndpoint(svc)
	hand := kithttp.NewServer(end, service.DecodeUserServiceRequest, service.EncodeUserServiceResponse)
	r := mux.NewRouter()
	r.Methods("GET").Path("/user/{uid:[0-9]+}").Handler(hand)
	uitl.Register()

	errChan := make(chan error)
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			errChan <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-sig)
	}()
	err := <-errChan
	uitl.DeRegister()
	fmt.Println(err)
}
