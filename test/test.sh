export LOG_DB_CONNECTION_STR="influenzanet-dev-pwvbz.mongodb.net/test?retryWrites=true&w=majority"
export LOG_DB_USERNAME="user-management-service"
export LOG_DB_PASSWORD="89kJAO43BRUyNSbr"
export LOG_DB_CONNECTION_PREFIX="+srv"

export DB_TIMEOUT=30
export DB_IDLE_CONN_TIMEOUT=45
export DB_MAX_POOL_SIZE=8
export DB_DB_NAME_PREFIX="INF_"

# use -v for verbose for the script
go test ./... $1