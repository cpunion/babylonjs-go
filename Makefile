all: pkg gen build test

pkg:
	./scripts/version.sh
	cat scripts/pkgs.txt | sed -e 's/\/.*$$//' | xargs -I {} yarn add {}@$(BABYLONJS_VERSION)

gen:
	go run ./cmd/gen/ scripts/pkgs.txt

build:
	go build -v ./...

test:
	go test -v ./...

release:
	./scripts/version.sh
	if git rev-parse $(BABYLONJS_VERSION) >/dev/null 2>&1; then
		echo "Version $(BABYLONJS_VERSION) already exists"
		exit 0
	fi
	# Add files to git
	TAG=v$(BABYLONJS_VERSION)
	git add -A
	git commit -m "Release $TAG"
	git tag $TAG
	git push origin main
	git push origin $TAG
