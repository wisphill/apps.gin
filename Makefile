build:
	docker build -t apps.gin .
run:
	docker run --env-file .env -p 8080:8080 apps.gin
