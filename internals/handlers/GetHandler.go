package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"KVStore/internals/utils"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	db := utils.GetShardDB(key)

	stmt, err := db.Prepare("SELECT value, ttl FROM KVStore WHERE `key` = ? AND ttl > ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var value string
	var ttl int64
	err = stmt.QueryRow(key, time.Now().Unix()).Scan(&value, &ttl)
	
	if err == sql.ErrNoRows {
		http.Error(w, "Key not found or TTL expired", http.StatusNotFound)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Value: %s", value)
}