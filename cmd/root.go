package cmd

import (
	"fmt"
	"github.com/fr3dch3n/commit/git"
	"github.com/fr3dch3n/commit/input"
	"github.com/fr3dch3n/commit/utils"
	log "github.com/sirupsen/logrus"
	"os"
)
import "github.com/spf13/cobra"

var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "commit",
	Short: "Easily build up a commit-message that conforms your team-conventions.",
	Run: func(cmd *cobra.Command, args []string) {
		commit()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const CommitConfigPath = "/.commit-config"

func commit() {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	if x, _ := utils.Exists(".git"); !x {
		log.Fatal("not in a git dir, aborting")
		os.Exit(1)
	}

	homedir := os.Getenv("HOME")
	configPath := homedir + CommitConfigPath

	commitConfig, err := input.ReadCommitConfig(configPath)
	utils.Check(err)
	log.Debug(commitConfig)

	if err := input.ContainsMinimalSet(commitConfig); err != nil {
		utils.Check(err)
	}

	teamMembers, err := input.ReadTeamMembersConfig(commitConfig.TeamMembersConfigPath)
	utils.Check(err)
	log.Debug(teamMembers)

	log.Debug("GithubUsername: " + commitConfig.GithubUsername)

	pair, err := git.GetPair(commitConfig, teamMembers)
	utils.Check(err)
	log.Debug("Pair: " + pair.String())

	story, err := input.GetInputOrElse(os.Stdin, "Story", commitConfig.CurrentStory)
	utils.Check(err)
	log.Debug("Story: " + story)

	err = input.WriteCommitConfig(configPath, pair, story, commitConfig)
	utils.Check(err)

	summary, err := input.GetNonEmptyInput(os.Stdin, "Summary of your commit")
	utils.Check(err)
	log.Debug("Summary: " + summary)
	reviewedSummary := git.ReviewSummary(summary)
	log.Debug("ReviewedSummary: " + reviewedSummary)

	explanation, err := input.GetMultiLineInput(os.Stdin, "Why did you choose to do that? ")
	utils.Check(err)
	log.Debug("Explanation: " + explanation)

	commitMsg := git.BuildCommitMsg(story, pair, reviewedSummary, explanation, commitConfig.Short)
	log.Debug("CommitMsg: " + commitMsg)

	git.Commit(commitMsg)
}
