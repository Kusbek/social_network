read -p "provide name: " filename
migrate create -ext sql -dir pkg/db/migrations/sqlite $filename