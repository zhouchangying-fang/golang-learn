package zhou_net

import (
	"fmt"
	"net/http"
)

func init() {

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("acbf")
		fmt.Printf("writer:%s,request:%s", writer, request)
	})
	http.ListenAndServe(":8080", nil)
}
