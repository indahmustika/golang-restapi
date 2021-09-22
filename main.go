package main
import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/select", selectUser).Methods("GET")
	router.HandleFunc("/insert", insertUser).Methods("POST")
	router.HandleFunc("/update", updateUser).Methods("PUT")
	router.HandleFunc("/delete", deleteUser).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}