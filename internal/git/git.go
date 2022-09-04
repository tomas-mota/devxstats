package git

import (
	"devxstats/internal/config"
	"devxstats/internal/git/bitbucket"
	"devxstats/internal/git/github"
	"devxstats/model"
)

type GitSyncer struct {
	sources []GitClient
}

type GitClient interface {
	GetCommits() ([]*model.Commit, error)
	GetOpenPullRequests() ([]*model.PullRequest, error)
}

func NewGitSyncer(c *config.GitConfig) *GitSyncer {
	syncer := &GitSyncer{}
	// Add sources based on configuration
	bc, err := bitbucket.NewBitbucketClient("https://dcgit.dac.local")
	if err != nil {
		panic(err)
	}
	syncer.sources = append(syncer.sources, bc)

	githubClient, err := github.NewClient(
		&github.GithubConfig{
			BaseUrl: c.Github.Url,
			Token:   c.Github.Token,
		})
	if err != nil {
		panic(err)
	}

	syncer.sources = append(syncer.sources, githubClient)

	return syncer
}

func (git *GitSyncer) Sync() error {
	for _, source := range git.sources {
		_, err := source.GetCommits()
		if err != nil {
			return err
		}

		// TODO: Persist Commits

		_, err = source.GetOpenPullRequests()
		if err != nil {
			return err
		}

		// TODO: Persist PullRequests
	}
	return nil
}
