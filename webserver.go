package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("p", 3001, "port number")
	flag.Parse()
	portNum := ":" + strconv.Itoa(*port)
	fmt.Printf("server start %s - %s \n", portNum, time.Now())
	served := "/home/justin/go/src/github.com/aegric/template"
	fs := http.FileServer(http.Dir(served))
	http.Handle("/", fs)
	http.ListenAndServe(portNum, nil)

}
