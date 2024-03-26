run-api:
	cd app && swag init
	cp app/docs/swagger.json docs/swagger.json
	cd docs && sed -i '' '1r servers.json' swagger.json
	cd app && go run main.go

run-docs:
	cd docs && npm run start
