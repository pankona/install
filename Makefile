
dry-run:
	goreleaser --skip-publish --rm-dist

release:
	goreleaser --rm-dist

