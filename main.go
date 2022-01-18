package main

import (
	"context"
	"flag"
)

func main() {
	ctx := context.Background()
	var (
		asdf      = flag.Bool("asdf", false, "install asdf")
		bashrc    = flag.Bool("bashrc", false, "install bashrc")
		asdfTools = flag.Bool("asdf-tools", false, "install tools via asdf")
		git       = flag.Bool("git", false, "install git")
		golang    = flag.Bool("golang", false, "install golang")
		vim       = flag.Bool("vim", false, "install vim")
		gh        = flag.Bool("gh", false, "install gh")
		docker    = flag.Bool("docker", false, "install docker")
		prettier  = flag.Bool("prettier", false, "install prettier")
		kubectl   = flag.Bool("kubectl", false, "install kubectl")
	)
	flag.Parse()

	switch {
	case *asdf:
		installAsdf(ctx)
	case *bashrc:
		installBashrc(ctx)
	case *asdfTools:
		installBuildEssential(ctx)
		installToolsViaAsdf(ctx)
	case *git:
		installGitConfig(ctx)
		installGit(ctx)
	case *golang:
		installGouse(ctx)
		installGo(ctx)
		installGoImports(ctx)
	case *vim:
		installVim(ctx)
		installVimrc(ctx)
		installAg(ctx)
	case *gh:
		installGitHubCLI(ctx)
	case *docker:
		installDocker(ctx)
		installDockerCompose(ctx)
	case *prettier:
		installPrettier(ctx)
	case *kubectl:
		installKubectl(ctx)
		installArgoRollouts(ctx)
	}
}
