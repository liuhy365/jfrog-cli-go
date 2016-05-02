package commands

import (
	"github.com/jfrogdev/jfrog-cli-go/artifactory/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
)

// Copies the artifacts using the specified move pattern.
func Copy(sourcePattern, destPath string, flags *utils.Flags) {
	utils.PreCommandSetup(flags)
	resultItems := utils.AqlSearch(sourcePattern, flags)
	regexpPath := cliutils.PathToRegExp(sourcePattern)

	utils.MoveFiles(regexpPath, resultItems, destPath, flags, utils.COPY)
}