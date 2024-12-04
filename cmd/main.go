package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"vk-im/internal/echo"
	"vk-im/internal/router"
)

var addr = flag.String("addr", ":8080", "websocket service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	//flag.Parse()
	//server := echo.NewServer()
	//go server.Run()
	//http.HandleFunc("/", serveHome)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	echo.ServeWS(server, w, r)
	//})
	//err := http.ListenAndServe(*addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	newRouter := router.Routers()

	go echo.NewServer().Run()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
