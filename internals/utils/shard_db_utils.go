package utils

import (
	"KVStore/internals/consts"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/twmb/murmur3"
)

var shardDBs[10] *sql.DB

func init() {
    for i := 0; i < len(shardDBs); i++ {
        db, err := sql.Open("mysql", consts.DATABASE_CREDS)
        if err != nil {
            log.Fatal(err)
        }

        dbName := fmt.Sprintf("%s_%d", consts.DATABASE_ROOT_NAME, i+1)
        _, err = db.Exec(fmt.Sprintf(consts.DATABASE_CREATION_QUERY, dbName))

        if err != nil {
            log.Fatal(err)
        }

        db, err = sql.Open("mysql", fmt.Sprintf("%s%s", consts.DATABASE_CREDS, dbName))
        if err != nil {
            log.Fatal(err)
        }

        _, err = db.Exec(consts.TABLE_CREATION_QUERY)
		if err != nil {
			log.Fatal(err)
		}

		shardDBs[i] = db
    }
}

func CloseDBConnections() {
	for _, db := range shardDBs {
		if db == nil {
			continue
		}

		err := db.Close()
		if err != nil {
			log.Print("Error closing the database")
		}
	}
}

func GetShardDB(key string) *sql.DB {
	return shardDBs[getShardIndex(key)]
}

func getShardIndex(key string) int {
	hasher := murmur3.SeedNew32(consts.SHARD_SEED_VALUE)
	_, err := hasher.Write([]byte(key))

	if err != nil {
		log.Fatalf("Failed to write key to the hasher: %v", err)
	}

	return int(hasher.Sum32()) % len(shardDBs)
}