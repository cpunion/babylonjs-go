BABYLONJS_VERSION := v$(shell node -p "require('./package.json').version")

all: pkg gen build test

pkg:
	cat scripts/pkgs.txt | sed -e 's/\/.*$$//' | xargs -I {} yarn add {}@$(BABYLONJS_VERSION)

gen:
	go run ./cmd/gen/ scripts/pkgs.txt

build:
	go build -v ./...

test:
	go test -v ./...

release:
	if git rev-parse $(BABYLONJS_VERSION) >/dev/null 2>&1;\
	then\
		echo "Version "$(BABYLONJS_VERSION)" already exists!";\
	else\
		echo "Releasing version $(BABYLONJS_VERSION)";\
		git add -A;\
		git commit -m "Release $(BABYLONJS_VERSION)";\
		git tag $(BABYLONJS_VERSION);\
		git push origin main;\
		git push origin $(BABYLONJS_VERSION);\
	fi
