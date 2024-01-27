package consts

const (
	DATABASE_CREATION_QUERY = "CREATE DATABASE IF NOT EXISTS %s"
	DATABASE_CREDS          = "test_user:12345678@/"
	DATABASE_ROOT_NAME      = "kvstore"
	SHARD_SEED_VALUE        = 42
	TABLE_CREATION_QUERY    = "CREATE TABLE IF NOT EXISTS KVStore (" +
		"`key` VARCHAR(255) PRIMARY KEY, " +
		"value VARCHAR(255), " +
		"ttl BIGINT, " +
		"INDEX (ttl)" +
		")"
)
