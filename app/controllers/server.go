package controllers

import (
	"fmt"
	"go-practice/app/models"
	"go-practice/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie_go")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}

	return session, err
}
