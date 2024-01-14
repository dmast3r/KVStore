package consts

const (
	DATABASE_CREATION_QUERY = "CREATE DATABASE IF NOT EXISTS %s"
	DATABASE_CREDS = "root:2402@/"
	DATABASE_ROOT_NAME = "kvstore"
	SHARD_SEED_VALUE = 42
	TABLE_CREATION_QUERY = "CREATE TABLE IF NOT EXISTS KVStore (" +
	"`key` VARCHAR(255) PRIMARY KEY, " +
	"value VARCHAR(255), " +
	"ttl BIGINT, " +
	"INDEX (ttl)" +
")"
)