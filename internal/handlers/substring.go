package handlers

import (
	"fmt"
	"go-test-softweather/internal/service"
	"io"
	"net/http"
)

func SubstringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {

			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("body is empty or ReadAll have error"))
			return
		}

		res := service.FindLongestSubstring(string(body))

		answer := fmt.Sprintf("Исходная строка - %s, самая длинная подстрока без повторяющихся символов - %s\n", string(body), res)
		_, _ = w.Write([]byte(answer))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
