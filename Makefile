IMAGE=tmdb-cli

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run --rm --env-file .env $(IMAGE) --type popular --limit 5
