create-setup: drop-setup
	@docker run --name codeinquarentena \
		-e MONGO_INITDB_ROOT_USERNAME=root \
    -e MONGO_INITDB_ROOT_PASSWORD=root \
		-e MONGO_INITDB_DATABASE=madalena \
		-v ${PWD}/seed:/tmp/seed \
		-p 27017:27017 \
		-d mongo:latest


drop-setup:
	@docker rm -f codeinquarentena 2>/dev/null || exit 0

