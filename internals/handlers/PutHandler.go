package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"time"
	"KVStore/internals/utils"
)

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   *int64 `json:"ttl"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Error Decoding the Payload", http.StatusBadRequest)
		return
	}

	if isPayloadValid, validationMessage := validatePayload(&p); !isPayloadValid {
		http.Error(w, validationMessage, http.StatusBadRequest)
		return
	}

	enrichTTL(&p)

	db := utils.GetShardDB(p.Key)

	stmt, err := db.Prepare("INSERT INTO KVStore (`key`, `value`, `ttl`) VALUES (?, ?, ?) " +
		"ON DUPLICATE KEY UPDATE `value` = VALUES(`value`), `ttl` = VALUES(`ttl`)")
	if err != nil {
		http.Error(w, "Error Preparing the Statement", http.StatusInternalServerError)
		log.Printf("Error Preparing the Statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Key, p.Value, *p.TTL)
	if err != nil {
		http.Error(w, "Error executing the statement", http.StatusInternalServerError)
		log.Printf("Error executing the statement: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Successfully Entered the Key!"))
	if err != nil {
		log.Printf("Error Writing the Response: %v", err)
		return
	}
}

func enrichTTL(p *Payload) {
	if p.TTL == nil {
		maxTTL := int64(math.MaxInt64)
		p.TTL = &maxTTL
	}
}

func validatePayload(p *Payload) (bool, string) {
	if p.Key == "" || p.Value == "" {
		return false, "Both key and value must be non-empty"
	}

	if p.TTL != nil && *p.TTL <= time.Now().Unix() {
		return false, "TTL must be a value greater than the current time"
	}

	return true, ""
}