package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/learn_go/controllers"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/abc", controllers.Home)

	mux.HandleFunc("/vmm/{theloai}/{maso}", controllers.ServeVideo)

	log.Println("Server dang khoi dong...")
	err := http.ListenAndServe(":4000", mux) //khai bao cong
	log.Fatal(err)                           //xuat chu mau do

}
