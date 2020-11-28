package main

import "net/http"
import "github.com/chengchaos/go-web-demo/data"

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err) {
	cookie, err := r.Cookie("_cookie")

	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")	
		}
	}

	return
}