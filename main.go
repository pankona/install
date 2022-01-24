package main

import (
	"context"
	"flag"
)

func main() {
	ctx := context.Background()
	var (
		asdf     = flag.Bool("asdf", false, "install asdf")
		bashrc   = flag.Bool("bashrc", false, "install bashrc")
		ruby     = flag.Bool("ruby", false, "install ruby via asdf")
		nodejs   = flag.Bool("nodejs", false, "install nodejs via asdf")
		yarn     = flag.Bool("yarn", false, "install yarn via asdf")
		ghq      = flag.Bool("ghq", false, "install ghq via asdf")
		git      = flag.Bool("git", false, "install git")
		golang   = flag.Bool("golang", false, "install golang")
		vim      = flag.Bool("vim", false, "install vim")
		gh       = flag.Bool("gh", false, "install gh")
		docker   = flag.Bool("docker", false, "install docker")
		prettier = flag.Bool("prettier", false, "install prettier")
		kubectl  = flag.Bool("kubectl", false, "install kubectl")
	)
	flag.Parse()

	switch {
	case *asdf:
		installAsdf(ctx)
	case *bashrc:
		installBashrc(ctx)
	case *ruby:
		installBuildEssential(ctx)
		installRuby(ctx)
	case *nodejs:
		installNodejs(ctx)
	case *yarn:
		installYarn(ctx)
	case *ghq:
		installGHQ(ctx)
	case *git:
		installGitConfig(ctx)
		installGit(ctx)
	case *golang:
		installGouse(ctx)
		installGo(ctx)
		installGoImports(ctx)
	case *vim:
		installVim(ctx) // apt
		installVimrc(ctx)
		installAg(ctx) // apt
	case *gh:
		installGitHubCLI(ctx) // apt
	case *docker:
		installDocker(ctx) // apt
		installDockerCompose(ctx)
	case *prettier:
		installPrettier(ctx)
	case *kubectl:
		installKubectl(ctx)
		installArgoRollouts(ctx)
	}
}
