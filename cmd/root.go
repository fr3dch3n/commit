package cmd

import (
	"fmt"
	"os"

	"github.com/fr3dch3n/commit/git"
	"github.com/fr3dch3n/commit/input"
	"github.com/fr3dch3n/commit/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool
var GitAddP bool
var SkipStory bool
var SkipPair bool
var SkipExplanation bool
var SkipAbbreviations bool
var GodMode bool
var Message string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&GitAddP, "git-add", "a", false, "run git add -p beforehand")
	rootCmd.PersistentFlags().BoolVarP(&SkipStory, "skip-story", "s", false, "skip story integration")
	rootCmd.PersistentFlags().BoolVarP(&SkipPair, "skip-pair", "p", false, "skip pair integration")
	rootCmd.PersistentFlags().BoolVarP(&SkipExplanation, "skip-explanation", "e", false, "skip long explanation")
	rootCmd.PersistentFlags().BoolVarP(&SkipAbbreviations, "skip-abbreviations", "n", false, "skip listing abbreviations")
	rootCmd.PersistentFlags().BoolVarP(&GodMode, "add-all-with-defaults", "y", false, "git add all and use defaults from state")
	rootCmd.PersistentFlags().StringVarP(&Message, "message", "m", "", "provide the commit-message")
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

const CommitConfigPath = ".commit-config"
const StatePath = ".commit-config.state"

func commit() {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	if x, _ := utils.Exists(".git"); !x {
		panic("not in a git dir, aborting")
	}

	if !git.AreThereChanges() {
		fmt.Println("No changes.")
		os.Exit(0)
	}

	homedir := os.Getenv("HOME")
	configPath := homedir + "/" + CommitConfigPath

	commitConfig, err := input.InitCommitConfig(configPath)
	utils.Check(err)
	log.Debug(commitConfig)

	teamMembers, err := input.InitTeamMembersConfig(commitConfig.TeamMembersConfigPath)
	utils.Check(err)
	log.Debug(teamMembers)

	var me input.TeamMember
	me, err = git.GetTeamMemberByAbbreviation(teamMembers, commitConfig.Abbreviation)
	if err != nil && err.Error() == "not-found" {
		newMember, err := git.GetAndSaveNewTeamMember(commitConfig.TeamMembersConfigPath, commitConfig.Abbreviation, teamMembers)
		if err != nil {
			panic(err)
		}
		me = newMember
	} else if err != nil {
		panic(err)
	}

	state, err := input.ReadState(homedir + "/" + StatePath)
	utils.Check(err)
	log.Debug(state)

	if GodMode {
		git.Add(".")
	} else if GitAddP {
		git.Add("-p")
	}

	if !git.AnythingStage() {
		fmt.Println("There are no staged files.")
		os.Exit(0)
	}

	var pair input.TeamMember
	var story string
	if GodMode {
		pair, err = git.GetTeamMemberByAbbreviation(teamMembers, state.CurrentPair)
		utils.Check(err)
		story = state.CurrentStory
		fmt.Printf("Using pair \"%s\" with story \"%s\".\n", pair.GithubUserName, story)
	} else {
		if !SkipPair {
			pair, err = git.GetPair(commitConfig, state.CurrentPair, teamMembers)
			utils.Check(err)
		}
		if !SkipStory {
			story, err = input.GetWithDefault("Current story", state.CurrentStory)
			utils.Check(err)
		}
		err = input.WriteState(homedir+"/"+StatePath, pair, story)
		utils.Check(err)
	}
	log.Debug("Pair: " + pair.String())
	log.Debug("Story: " + story)

	var summary string
	if Message == "" {
		summary, err = input.GetNonEmpty("Summary of your commit")
		utils.Check(err)
	} else {
		summary = Message
	}
	log.Debug("Summary: " + summary)
	reviewedSummary := git.ReviewSummary(summary)
	log.Debug("ReviewedSummary: " + reviewedSummary)

	var explanation string
	if !SkipExplanation {
		explanation, err = input.GetMultiLineInputV2("Why did you choose to do that? ")
		utils.Check(err)
		log.Debug("Explanation: " + explanation)
	}

	commitMsg := git.BuildCommitMsg(story, pair, reviewedSummary, explanation, me, SkipAbbreviations)
	log.Debug("CommitMsg: " + commitMsg)

	git.Commit(commitMsg)
}
