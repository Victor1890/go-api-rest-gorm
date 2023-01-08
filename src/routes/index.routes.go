package routes

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var headers = w.Header()

	fmt.Println(headers)

	w.Write([]byte("Hello world 3"))

}
