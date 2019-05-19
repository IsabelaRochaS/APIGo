package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/isabelarochas/restapi/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "urlteste"
)

type Url struct {
	model.Url
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("postgres",
		"host="+host+" user="+user+
			" dbname="+dbname+" sslmode=disable password="+
			password)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Url{})

	fmt.Println("Successfully connected!")

	r := mux.NewRouter()

	r.HandleFunc("/api/url", getUrls).Methods("GET")
	r.HandleFunc("/api/url/{alias}", getUrl).Methods("GET")
	r.HandleFunc("/api/url", createUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getUrls(w http.ResponseWriter, r *http.Request) {
	var url []Url
	db.Limit(10).Order("visit_num desc").Find(&url)
	json.NewEncoder(w).Encode(&url)
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var url Url
	db.Where("alias = ?", params["alias"]).First(&url)
	url.VisitNum = url.VisitNum + 1
	db.Save(&url)
	json.NewEncoder(w).Encode(&url)
	//exec.Command("xdg-open", url).Start()
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url.Url.Url).Start()
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	var url Url
	json.NewDecoder(r.Body).Decode(&url)

	if url.Alias == "" {
		//Entrar aqui com alias aleatorio
	}

	db.Create(&url)
	json.NewEncoder(w).Encode(&url)
}
