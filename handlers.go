package main

import (
	"encoding/json"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {

}
