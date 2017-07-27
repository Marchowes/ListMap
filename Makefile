


# Install glide dependencies
.PHONY: vendor
vendor:
	docker-compose run --rm vendor
