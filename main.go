package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	installVim(ctx)
	installVimrc(ctx)
	installAg(ctx)
	installAsdf(ctx)
	installBuildEssential(ctx)
	installToolsViaAsdf(ctx)
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
