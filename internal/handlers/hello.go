package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, _ = w.Write([]byte("чтобы найти самую длинную подстроку в строке, отправте ее POST запросом по url /api/substring"))
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
