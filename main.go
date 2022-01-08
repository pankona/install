package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	installAsdf(ctx)
	installToolsViaAsdf(ctx)
	installVim(ctx)
	installVimrc(ctx)
	installAg(ctx)
	installBuildEssential(ctx)
	installBashrc(ctx)
	installGit(ctx)
	installGitConfig(ctx)
	installGitHubCLI(ctx)
	installDocker(ctx)
	installDockerCompose(ctx)
	installGouse(ctx)
	installGo(ctx)
	installGoImports(ctx)
	installPrettier(ctx)
}
