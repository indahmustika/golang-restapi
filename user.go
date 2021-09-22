package main
import(
	"encoding/json"
	"log"
	"net/http"
)

func selectUser(w http.ResponseWriter, r *http.Request){
	var users Users
	var userdata []Users
	var response Response

	db := connect()
	defer db.Close()
	
	rows,err := db.Query("SELECT * FROM account")
	if err != nil{
		log.Print(err)
	}

	for rows.Next(){
		if err := rows.Scan(&users.Username, &users.Password); err != nil{
			log.Fatal(err.Error())
		} else{
			userdata = append(userdata, users)
		}
	}

	response.Status  = 1
	response.Message = "Success"
	response.Data 	 = userdata

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUser(w http.ResponseWriter, r *http.Request){
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil{
		panic(err)}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err = db.Exec("INSERT INTO account (username, password) VALUES (?, ?)", username, password)
	
	if err != nil{
		log.Print(err)}

	var response Response
	response.Status  = 1
	response.Message = "Account Created"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request){
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil{
		panic(err)}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err = db.Exec("UPDATE account SET password = ? WHERE username = ?", password, username)
	
	if err != nil{
		log.Print(err)}

	var response Response
	response.Status  = 1
	response.Message = "Account Updated"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil{
		panic(err)}

	username := r.Form.Get("username")
	_, err = db.Exec("DELETE FROM account WHERE username = ?", username)

	if err != nil{
		log.Print(err)}

	var response Response
	response.Status  = 1
	response.Message = "Account Deleted"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}