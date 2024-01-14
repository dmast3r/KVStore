package handlers

import (
	"KVStore/internals/utils"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	db := utils.GetShardDB(key)

	var ttl int64
	err := db.QueryRow("SELECT ttl FROM KVStore WHERE `key` = ? AND ttl > ?", key, time.Now().Unix()).Scan(&ttl)
	
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Printf("Error checking key existence: %v", err)
		return
	}

	_, err = db.Exec("UPDATE KVStore SET ttl = -1 WHERE `key` = ?", key)
	if err != nil {
		http.Error(w, "Error executing the soft delete", http.StatusInternalServerError)
		log.Printf("Error executing the soft delete: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Key soft-deleted successfully"))
	if err != nil {
		log.Printf("Error Writing the Response: %v", err)
		return
	}
}