package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



// get request
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LAKSJDLKDJ")
}




// post request
// handles form data
func formHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w,r, "./static/form.html")
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		} 
		user := r.FormValue("User")
		pass := r.FormValue("Password")
		//log.Print(r.Form["User"])
		print(user, pass)
		http.ServeFile(w,r,"./static/form.html")

	default:
		fmt.Fprintf(w, "only GET and POST are supported.")
	}

	


}

func main() {
	router := mux.NewRouter()

/*

	// error handler
	http.HandleFunc("/test", testHandler);
	http.HandleFunc("/login", formHandler);
*/
	router.HandleFunc("/login", formHandler)

	fmt.Printf("Starting server at port 8080")
	if err:= http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}