DB_URL=postgres://${DB_USER_NAME}:${DB_PASSWORD}@${DB_HOST_NAME}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}

migrateup:
	migrate -path ./scripts/migrations/ -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path ./scripts/migrations/ -database "$(DB_URL)" -verbose up 1
migratedown:
	migrate -path ./scripts/migrations/ -database "$(DB_URL)" -verbose down
migratedown1:
	migrate -path ./scripts/migrations/ -database "$(DB_URL)" -verbose down 1

server:
	go run .

test:
	go test -v -cover ./...
	
git:
	git add . && git commit -m "$(ARG)" && git push
