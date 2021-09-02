package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// global var
var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)
// struct def
type User struct {
	Username      string
	Authenticated bool
}

// init function
func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	//required
	gob.Register(User{})
}

// authentication
func authenticate(user string, pass string) bool {

	if user == "admin" && pass == "password" {
		return true
	} else {
		return false
	}
}

// create session
func session(w http.ResponseWriter, r *http.Request, user *User) {
	// store.Get returns (*session, err)
	// returns new session is sessions doesn't exist
	session, _ := store.Get(r, "session-token")
	session.Values["user"] = user
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// permanent redirect to frontendURI/dashboard
	log.Println("jsalkdjalks")
	http.Redirect(w, r, "/secret", http.StatusMovedPermanently)
}

// post request
func formHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./static/form.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		username := r.FormValue("User")
		pass := r.FormValue("Password")

		var authenticated bool = authenticate(username, pass)
		if authenticated {
			user := &User{
				Username:      username,
				Authenticated: true,
			}
			session(w, r, user)
		} else {
			http.ServeFile(w, r, "./static/form.html")
		}

	default:
		fmt.Fprintf(w, "only GET and POST are supported.")
	}

}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user := getUser(session)
	log.Print(user.Username)
}

func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	// destructuring??
	// no idea what this means?
	// maybe its dereferencing because User is a pointer?
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

func main() {
	router := mux.NewRouter()

	/*
		// error handler
		http.HandleFunc("/test", testHandler);
		http.HandleFunc("/login", formHandler);
	*/
	router.HandleFunc("/login", formHandler)
	router.HandleFunc("/secret", secretHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
