package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tung"))
}

func ServeVideo(w http.ResponseWriter, r *http.Request) {

	maso := r.URL.Query().Get("id") //
	theloai := r.URL.Query().Get("category")
	w.Header().Add("Content-Type", "video/mp4")         //khai bao kieu video
	path := fmt.Sprintf("vmm/%s/%s.mp4", theloai, maso) //%d khia bao so, %s khai bao chu, khai bao sap xep link

	if _, err := os.Stat(path); os.IsNotExist(err) { // tim file co ton tai hay khong
		w.Header().Add("content-type", "text/html; charset=UTF-8")
		http.ServeFile(w, r, "./456.html") //neu khong tim thay thi hien thi file nay
		return
	}
	http.ServeFile(w, r, path) // khai bao duong dan, tao sever
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/abc", home)

	mux.HandleFunc("/video", ServeVideo)

	log.Println("Server dang khoi dong...")
	err := http.ListenAndServe(":4000", mux) //khai bao cong
	log.Fatal(err)                           //xuat chu mau do

}
