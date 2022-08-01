
dry-run:
	goreleaser --skip-publish --rm-dist

release:
	goreleaser --rm-dist

install-tools:
	go install github.com/goreleaser/goreleaser@latest
