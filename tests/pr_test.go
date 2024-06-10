package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
	"testing"
)

func TestProjectsFullTest(t *testing.T) {
	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing: t,
		Prefix:  "alm-stack",
		StackConfigurationOrder: []string{
			"1 - Infrastructure for Application Lifecycle Management",
			"2 - Application Lifecycle Management"
		},
	})

	options.StackInputs = map[string]interface{}{
		"ibmcloud_api_key":            options.RequiredEnvironmentVars["TF_VAR_ibmcloud_api_key"],
		"prefix":                      options.Prefix,
		"region":                      options.Region
		
	}

	err := options.RunProjectsTest()
	if assert.NoError(t, err) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}

