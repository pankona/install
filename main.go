package main

import (
	"context"
	"flag"
)

func main() {
	ctx := context.Background()
	var (
		asdf   = flag.Bool("asdf", false, "install asdf")
		bashrc = flag.Bool("bashrc", false, "install bashrc")
		all    = flag.Bool("all", false, "install all")
	)
	flag.Parse()

	switch {
	case *all:
		installAsdf(ctx)
		installBashrc(ctx)
		installToolsViaAsdf(ctx)
		installGouse(ctx)
		installGo(ctx)
		installGoImports(ctx)
		installGit(ctx)
		installGitConfig(ctx)
		installVim(ctx)
		installVimrc(ctx)
		installAg(ctx)
		installBuildEssential(ctx)
		installGitHubCLI(ctx)
		installDocker(ctx)
		installDockerCompose(ctx)
		installPrettier(ctx)
	case *asdf:
		installAsdf(ctx)
	case *bashrc:
		installBashrc(ctx)
	}
}
