package main

import (
	"flag"
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang-learn/go-kit/ixigua/service"
	"golang-learn/go-kit/ixigua/uitl"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	name := flag.String("name", "service", "服务名")
	port := flag.Int("p", 8080, "端口")
	flag.Parse()
	fmt.Println("--", *name, *port)
	limit := rate.NewLimiter(1, 3)
	uitl.SetNamePort(*name, *port)
	svc := service.UserServiceImpl{}
	end := service.RateLimit(limit)(service.MakeUserEndpoint(svc))
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(uitl.MyErrorEncoder),
	}
	hand := kithttp.NewServer(end, service.DecodeUserServiceRequest, service.EncodeUserServiceResponse, options...)
	r := mux.NewRouter()
	r.Methods("GET").Path("/user/{uid:[0-9]+}").Handler(hand)
	uitl.Register()
	errChan := make(chan error)
	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(*port), r); err != nil {
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
