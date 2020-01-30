package integreatly_operator_test_harness

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"github.com/integr8ly/integreatly-operator-test-harness/pkg/metadata"

	_ "github.com/integr8ly/integreatly-operator-test-harness/pkg/tests"
)

const (
	testResultsDirectory = "/test-run-results"
	jUnitOutputFilename  = "junit-integreatly-operator.xml"
	addonMetadataName    = "addon-metadata.json"
)

func TestIntegreatlyOperatorTestHarness(t *testing.T) {
	RegisterFailHandler(Fail)
	jUnitReporter := reporters.NewJUnitReporter(filepath.Join(testResultsDirectory, jUnitOutputFilename))

	RunSpecsWithDefaultAndCustomReporters(t, "Integreatly Operator Test Harness", []Reporter{jUnitReporter})

	err := metadata.Instance.WriteToJSON(filepath.Join(testResultsDirectory, addonMetadataName))
	if err != nil {
		t.Errorf("error while writing metadata: %v", err)
	}
}
