package bundler_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitBundler(t *testing.T) {
	suite := spec.New("bundler", spec.Report(report.Terminal{}))
	suite("Build", testBuild)
	suite("BuildpackYMLParser", testBuildpackYMLParser)
	suite("Detect", testDetect)
	suite("GemfileLockParser", testGemfileLockParser)
	suite("LogEmitter", testLogEmitter)
	suite("PlanRefinery", testPlanRefinery)
	suite("VersionShimmer", testVersionShimmer)
	suite.Run(t)
}
