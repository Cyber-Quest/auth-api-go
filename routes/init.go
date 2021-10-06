package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Cyber-Quest/go-rest-api/help/array/interfaces"
	"github.com/Cyber-Quest/go-rest-api/help/hash"
	"github.com/Cyber-Quest/go-rest-api/help/token"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type user struct {
	ID       string `json:"ID"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Role     string `json:"Role"`
}
type authentication struct {
	User  user
	Token string
}

type allUsers []user
type allAuth []authentication

var auths = allAuth{}

func signIn(w http.ResponseWriter, r *http.Request) {
	var newAuth user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insira os dados  corretamente")
	}
	json.Unmarshal(reqBody, &newAuth)
	w.WriteHeader(http.StatusCreated)

	var index = interfaces.Find(auths, func(value interface{}) bool {
		return value.(authentication).User.Email == newAuth.Email
	})
	var valided = hash.CheckPasswordHash(newAuth.Password, auths[index].User.Password)
	var password = ""
	if valided == true {
		password = auths[index].User.Password
	}
	var contains = token.Compare(newAuth.Email, password, auths[index].User.Role, auths[index].User.ID, auths[index].Token)
	if contains == true {
		json.NewEncoder(w).Encode(auths[index])
	} else {
		json.NewEncoder(w).Encode("Usuário ou senha estão incorretos!")
	}

}

func signUp(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insira os dados  corretamente")
	}
	json.Unmarshal(reqBody, &newUser)
	password, err := hash.GeneratehashPassword(newUser.Password)
	if err != nil {
		fmt.Fprintf(w, "Erro ao gerar o token")
	}
	var index = interfaces.Find(auths, func(value interface{}) bool {
		return value.(authentication).User.Email == newUser.Email
	})
	if index == -1 {
		var id = uuid.New().String()
		var newToken = token.Create(newUser.Email, password, "customer", id)
		var newAuth authentication
		newAuth.User.ID = id
		newAuth.User.Email = newUser.Email
		newAuth.User.Password = password
		newAuth.User.Role = "customer"
		newAuth.Token = newToken

		auths = append(auths, newAuth)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(auths)
	} else {
		fmt.Fprintf(w, "Usuário já foi registrado no servidor!")
	}
}

// Public visibility
func InitializeServer(message string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/signin", signIn).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("POST")
	fmt.Printf(message)
	log.Fatal(http.ListenAndServe(":8080", router))
}
