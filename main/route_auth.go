package main

import (
	"net/http"

	"github.com/chengchaos/go-web-demo/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	uesr, _ := data.UserByEmail(r.PostFormValue("email"))

	if uesr.Password == data.Encrypt(r.PostFormValue("password")) {
		session := uesr.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
