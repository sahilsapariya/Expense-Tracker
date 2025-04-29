package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// Transaction model
type Transaction struct {
    ID          uint    `json:"id"`
    Amount      float64 `json:"amount"`
    Category    string  `json:"category"`
    Description string  `json:"description"`
    Date        string  `json:"date"`
    Type        string  `json:"type"`
}

// Initialize the database
func initDB() {
    db, err = gorm.Open("sqlite3", "./transactions.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    db.AutoMigrate(&Transaction{})
}

// Create a new transaction
func createTransaction(w http.ResponseWriter, r *http.Request) {
    var transaction Transaction
    json.NewDecoder(r.Body).Decode(&transaction)

    db.Create(&transaction)

    json.NewEncoder(w).Encode(transaction)
}

// Get all transactions
func getTransactions(w http.ResponseWriter, r *http.Request) {
    var transactions []Transaction
    db.Find(&transactions)

    json.NewEncoder(w).Encode(transactions)
}

// Setup routes
func handleRequests() {
    r := mux.NewRouter()
    r.HandleFunc("/transactions", createTransaction).Methods("POST")
    r.HandleFunc("/transactions", getTransactions).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
    initDB()
    handleRequests()
}
