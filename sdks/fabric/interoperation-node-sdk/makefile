.PHONY: build-local
build-local: .npmrc
	# echo "Calling update-protos-local script..."
	# ./scripts/update-protos-local.sh
	npm install
	# npm run protos
	npm run build

.PHONY: build
build: .npmrc
	npm install
	npm run build
	
.PHONE: publish
publish: .npmrc
	npm publish
	

clean:
	rm -rf build node_modules package-lock.json .nyc_output coverage
