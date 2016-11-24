package main

import (
	"net/http"
	"os"
	"fmt"
	//"github.com/lib/pq"
	"github.com/gorilla/mux"
	"log"
)

const (
	DB_DRIVER     = "postgres"
	DB_URI = "postgres://u735a0648be0147aca666d717ca207c9b:34a20b879af04ec7b2d85a4b2d82b655@10.72.6.143:5432/d8e9f989966c44f678a5e7fc8eddb8d57?sslmode=disable"
)

func main() {
	log.Println("Starting bjjdb...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/competitors", getCompetitors).Methods("GET")
	log.Println("Listening and serving...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	checkErr(err)


}


func getCompetitors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(req)
	//tenantUuid := vars["tenantUuid"]
	//
	//db, err := sql.Open(DB_DRIVER, DB_URI)
	//err = db.Ping()
	//checkErr(err)
	//
	//query := "SELECT * FROM usage_service.usage WHERE tenant_uuid = '" + tenantUuid +"'"
	//rows, err := db.Query(query)
	//
	//log.Printf(query)
	//
	//var usageArray UsageArray
	//
	//for rows.Next() {
	//	var id int
	//	var tenantUuid string
	//	var subscription_uuid string
	//	var usage_date time.Time
	//	var meter_name string
	//	var meter_value int
	//	var metadata string
	//
	//	err = rows.Scan(&id, &tenantUuid, &subscription_uuid, &usage_date, &meter_name, &meter_value, &metadata)
	//	checkErr(err)
	//
	//	u := Usage{id, tenantUuid, subscription_uuid, usage_date, meter_name, meter_value, metadata}
	//
	//	usageArray = append(usageArray, u)
	//}

	fmt.Fprintln(w, "BJJ Competitors")

	//rows.Close()
	//defer db.Close()
}


func checkErr(err error) {
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}
