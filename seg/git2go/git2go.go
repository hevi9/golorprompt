// TODO: mv file1 file2 does not show indicators, no wt delete, no wt new file

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"gopkg.in/libgit2/git2go.v24"
)

// Color levels
//
// 4 - Conflicted
// 3 - Wt New (Untracked)
// 2 - Wt Modified, Deleted, Renamed, TypeChange
// 1 - Index New, Modified, Deleted, Renamed, TypeChange
// 0 - OK

const (
	ColorLevelOk = iota
	ColorLevelIgnored
	ColorLevelCurrent

	ColorLevelIndexDeleted
	ColorLevelIndexRenamed
	ColorLevelIndexModified
	ColorLevelIndexTypeChange
	ColorLevelIndexNew

	ColorLevelWtDeleted
	ColorLevelWtRenamed
	ColorLevelWtModified
	ColorLevelWtTypeChange

	ColorLevelWtNew // Untracked

	ColorLevelConflicted
)

var colorLevelColor = map[int]colorful.Color{
	ColorLevelOk:      colorful.Hsv(90.0, sys.Config.FgSaturationLow, sys.Config.FgValue),
	ColorLevelIgnored: colorful.Hsv(90.0+10.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelCurrent: colorful.Hsv(90.0+20.0, sys.Config.FgSaturation, sys.Config.FgValue),

	ColorLevelIndexDeleted:    colorful.Hsv(55.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelIndexRenamed:    colorful.Hsv(55.0+5.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelIndexModified:   colorful.Hsv(55.0+10.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelIndexTypeChange: colorful.Hsv(55.0+15.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelIndexNew:        colorful.Hsv(55.0+20.0, sys.Config.FgSaturation, sys.Config.FgValue),

	ColorLevelWtDeleted:    colorful.Hsv(25.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelWtRenamed:    colorful.Hsv(25.0+5.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelWtModified:   colorful.Hsv(25.0+10.0, sys.Config.FgSaturation, sys.Config.FgValue),
	ColorLevelWtTypeChange: colorful.Hsv(25.0+15.0, sys.Config.FgSaturation, sys.Config.FgValue),

	ColorLevelWtNew: colorful.Hsv(330.0, sys.Config.FgSaturation, sys.Config.FgValue),

	ColorLevelConflicted: colorful.Hsv(0.0, sys.Config.FgSaturation, sys.Config.FgValue),
}

type GitRepoStatus struct {
	Current         int
	IndexNew        int
	IndexModified   int
	IndexDeleted    int
	IndexRenamed    int
	IndexTypeChange int
	WtNew           int
	WtModified      int
	WtDeleted       int
	WtRenamed       int
	WtTypeChange    int
	Ignored         int
	Conflicted      int
	Total           int
}

func getRepoStatus(repo *git.Repository) (*GitRepoStatus, error) {
	value := GitRepoStatus{}
	statusOptions := &git.StatusOptions{
		Show: git.StatusShowIndexAndWorkdir,
		Flags: git.StatusOptIncludeUntracked |
			git.StatusOptExcludeSubmodules |
			git.StatusOptRecurseUntrackedDirs |
			git.StatusOptRenamesHeadToIndex |
			git.StatusOptRenamesIndexToWorkdir |
			git.StatusOptRenamesFromRewrites |
			git.StatusOptNoRefresh,
		// Pathspec=nil is all files
	}
	statusList, err := repo.StatusList(statusOptions)
	if err != nil {
		return nil, err
	} else {
		defer statusList.Free()
		n, err := statusList.EntryCount()
		if err != nil {
			return nil, err
		}
		for i := 0; i < n; i++ {
			entry, err := statusList.ByIndex(i)
			if err != nil {
				log.Panicf("may not go out of index: statusList.ByIndex(%d): %s", i, err)
			}
			if (git.StatusCurrent & entry.Status) != 0 {
				value.Current += 1
			}
			if (git.StatusIndexNew & entry.Status) != 0 {
				value.IndexNew += 1
				value.Total += 1
			}
			if (git.StatusIndexModified & entry.Status) != 0 {
				value.IndexModified += 1
				value.Total += 1
			}
			if (git.StatusIndexDeleted & entry.Status) != 0 {
				value.IndexDeleted += 1
				value.Total += 1
			}
			if (git.StatusIndexRenamed & entry.Status) != 0 {
				value.IndexRenamed += 1
				value.Total += 1
			}
			if (git.StatusIndexTypeChange & entry.Status) != 0 {
				value.IndexTypeChange += 1
				value.Total += 1
			}
			if (git.StatusWtNew & entry.Status) != 0 {
				value.WtNew += 1
				value.Total += 1
			}
			if (git.StatusWtModified & entry.Status) != 0 {
				value.WtModified += 1
				value.Total += 1
			}
			if (git.StatusWtDeleted & entry.Status) != 0 {
				value.WtDeleted += 1
				value.Total += 1
			}
			if (git.StatusWtTypeChange & entry.Status) != 0 {
				value.WtTypeChange += 1
				value.Total += 1
			}
			if (git.StatusWtRenamed & entry.Status) != 0 {
				value.WtRenamed += 1
				value.Total += 1
			}
			if (git.StatusIgnored & entry.Status) != 0 {
				value.Ignored += 1
			}
			if (git.StatusConflicted & entry.Status) != 0 {
				value.Conflicted += 1
				value.Total += 1
			}
		}
	}
	return &value, nil
}

type Git struct{}

func main() {}

func (*Git) Render() []sys.Chunk {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Cannot get CWD: %s", err)
		return nil
	}
	repoPath, err := git.Discover(cwd, true, nil)
	if err != nil {
		//log.Printf("Cannot find git repo: %s", err)
		return nil
	}
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Printf("Cannot open git repository %s: %s\n", repoPath, err)
		return nil
	}
	if repo.IsBare() {
		return []sys.Chunk{sys.Chunk{Text: "-bare-", Fg: colorful.HappyColor()}}
	}
	detached, err := repo.IsHeadDetached()
	if err != nil {
		log.Printf("Cannot get detached info: %s", err)
		return nil
	}
	if detached {
		return []sys.Chunk{sys.Chunk{Text: "-detached-", Fg: colorful.HappyColor()}}
	}

	chunks := make([]sys.Chunk, 0)

	// Show head branch
	var headBranch *git.Branch
	var headBranchName string
	headRef, err := repo.Head()
	if err != nil {
		log.Printf("Cannot get head ref: %s\n", err)
	} else {
		headBranch = headRef.Branch()
		headBranchName, err = headBranch.Name()
		if err != nil {
			log.Printf("headBranch.Name(): %s", err)
			return nil
		}
	}
	headChunk := sys.Chunk{Text: headBranchName, Fg: colorful.HappyColor()}

	statusColorLevel := ColorLevelOk

	// Show git repository status
	status, err := getRepoStatus(repo)
	// TODO: handle git status err
	if status.Total > 0 {
		chunks = append(chunks, sys.Chunk{Text: " "})
	}

	if status.Current > 0 {
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.Current, 'c'),
			Fg:   colorLevelColor[ColorLevelCurrent],
		})
	}
	if status.IndexNew > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIndexNew)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.IndexNew, 'n'),
			Fg:   colorLevelColor[ColorLevelIndexNew],
		})
	}
	if status.IndexModified > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIndexModified)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.IndexModified, 'm'),
			Fg:   colorLevelColor[ColorLevelIndexModified],
		})
	}
	if status.IndexDeleted > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIndexDeleted)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.IndexDeleted, 'd'),
			Fg:   colorLevelColor[ColorLevelIndexDeleted],
		})
	}
	if status.IndexRenamed > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIndexRenamed)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.IndexRenamed, 'r'),
			Fg:   colorLevelColor[ColorLevelIndexRenamed],
		})
	}
	if status.IndexTypeChange > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIndexTypeChange)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.IndexTypeChange, 't'),
			Fg:   colorLevelColor[ColorLevelIndexTypeChange],
		})
	}
	if status.WtNew > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelWtNew)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.WtNew, 'N'),
			Fg:   colorLevelColor[ColorLevelWtNew],
		})
	}
	if status.WtModified > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelWtModified)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.WtModified, 'M'),
			Fg:   colorLevelColor[ColorLevelWtModified],
		})
	}
	if status.WtDeleted > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelWtDeleted)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.WtDeleted, 'D'),
			Fg:   colorLevelColor[ColorLevelWtModified],
		})
	}
	if status.WtTypeChange > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelWtTypeChange)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.WtTypeChange, 'T'),
			Fg:   colorLevelColor[ColorLevelWtTypeChange],
		})
	}
	if status.Ignored > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelIgnored)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.Ignored, 'I'),
			Fg:   colorLevelColor[ColorLevelIgnored],
		})
	}
	if status.Conflicted > 0 {
		statusColorLevel = sys.MaxInt(statusColorLevel, ColorLevelConflicted)
		chunks = append(chunks, sys.Chunk{
			Text: fmt.Sprintf("%d%c", status.Conflicted, 'C'),
			Fg:   colorLevelColor[ColorLevelConflicted],
		})
	}

	// TODO: Add space

	// Show ahead behind
	upstreamRef, err := headBranch.Upstream()
	if err != nil {
		// No upstream give a sign
		//log.Printf("headBranch.Upstream(): %s", err)
		chunks = append(chunks, sys.Chunk{
			Text: " noup",
			Fg:   sys.Config.FgWarning,
		})
	} else {
		ahead, behind, err := repo.AheadBehind(headRef.Target(), upstreamRef.Target())
		if err != nil {
			log.Printf("repo.AheadBehind(headRef.Target(), upstreamRef.Target()): %s", err)
		}
		if ahead > 0 || behind > 0 {
			chunks = append(chunks, sys.Chunk{Text: " "})
		}
		if ahead > 0 {
			chunks = append(chunks, sys.Chunk{
				Text: fmt.Sprintf("%d%s", ahead, sys.Sign.Ahead),
				Fg:   colorful.HappyColor(), // TODO: Set ahead color
			})
		}
		if behind > 0 {
			chunks = append(chunks, sys.Chunk{
				Text: fmt.Sprintf("%d%s", behind, sys.Sign.Behind),
				Fg:   colorful.HappyColor(), // TODO: Set behind color
			})
		}
	}

	// Show repository state
	state := repo.State()
	if state != git.RepositoryStateNone {
		name := stateName(state)
		hue := 360.0 * sys.HashToFloat64([]byte(name))
		chunks = append(chunks, sys.Chunk{
			Text: name,
			Fg:   colorful.Hsv(hue, sys.Config.FgSaturation, sys.Config.FgValue),
		})
	}

	// add head name chuck to front
	headChunk.Fg = colorLevelColor[statusColorLevel]
	chunks = append([]sys.Chunk{headChunk}, chunks...)

	// TODO: Show stash, when git2go used from version v24+

	return chunks
}

func stateName(state git.RepositoryState) string {
	switch state {
	case git.RepositoryStateNone:
		return ""
	case git.RepositoryStateMerge:
		return "merge"
	case git.RepositoryStateRevert:
		return "revert"
	case git.RepositoryStateCherrypick:
		return "cherrypick"
	case git.RepositoryStateBisect:
		return "bisect"
	case git.RepositoryStateRebase:
		return "rebase"
	case git.RepositoryStateRebaseInteractive:
		return "rebase-interactive"
	case git.RepositoryStateRebaseMerge:
		return "rebase-merge"
	case git.RepositoryStateApplyMailbox:
		return "apply-mailbox"
	case git.RepositoryStateApplyMailboxOrRebase:
		return "apply-mailbox-or-rebase"
	default:
		return "fault"
	}
}
