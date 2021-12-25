package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tung"))
}

func ServeVideo(w http.ResponseWriter, r *http.Request) {

	log.Println("Bat dau video...")
	vars := mux.Vars(r)
	id := vars["maso"] //
	theloai := vars["theloai"]
	w.Header().Add("Content-Type", "video/mp4")       //khai bao kieu video
	path := fmt.Sprintf("vmm/%s/%s.mp4", theloai, id) //%d khia bao so, %s khai bao chu, khai bao sap xep link

	if _, err := os.Stat(path); os.IsNotExist(err) { // tim file co ton tai hay khong
		w.Header().Add("content-type", "text/html; charset=UTF-8")
		http.ServeFile(w, r, "./456.html") //neu khong tim thay thi hien thi file nay
		return
	}
	http.ServeFile(w, r, path) // khai bao duong dan, tao sever
}
