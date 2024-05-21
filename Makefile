
dry-run:
	goreleaser --skip=publish --clean

release:
	goreleaser --clean

install-tools:
	go install github.com/goreleaser/goreleaser@latest
