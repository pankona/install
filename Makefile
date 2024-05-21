
dry-run:
	goreleaser --skip=publish --clean

release:
	goreleaser --rm-dist

install-tools:
	go install github.com/goreleaser/goreleaser@latest
