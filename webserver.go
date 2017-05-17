package main 

import ( "fmt"
         "net/http"
         "strconv"
         "flag"
         "time"
     )


     func main (){
         port := flag.Int("p",3001,"port number")
         flag.Parse()
         port_num := ":" + strconv.Itoa(*port)
         fmt.Printf("server start %s - %s \n",port_num,time.Now())
         served := "/home/justin/go/src/webserver/template"
         fs := http.FileServer(http.Dir(served))
         http.Handle("/",fs)
         http.ListenAndServe(port_num,nil)

     }

