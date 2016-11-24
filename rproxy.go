package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "strings"
)

func HandleIt(w http.ResponseWriter, r *http.Request) {
  if strings.Contains(r.Host, "wiki.rebirtharmitage.com"){
    http.Redirect(w, r, "http://wiki.rebirtharmitage.com:8083/dokuwiki/doku.php", 302)
  }else if strings.Contains(r.Host, "internetatlas.co"){
    http.Redirect(w, r, "http://www.internetatlas.co:8081", 302)
  }else if strings.Contains(r.Host, "janus"){
    http.Redirect(w, r, "http://www.gojanus.co:8083", 302)
  }else if strings.Contains(r.Host, "rebirtharmitage.com"){
    http.Redirect(w, r, "http://www.rebirtharmitage.com:8080", 302)
  }
}

var router = mux.NewRouter()

func main(){
  http.HandleFunc("/", HandleIt)
	http.ListenAndServe(":80", nil)
}