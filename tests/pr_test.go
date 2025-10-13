package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
	"log"
	"testing"
	//"fmt"
)

const permanentResourcesYaml = "../common-permanent-resources.yaml"

var permanentResources map[string]interface{}

func TestProjectsFullTest(t *testing.T) {

	var err error
	permanentResources, err = common.LoadMapFromYaml(permanentResourcesYaml)
	if err != nil {
		log.Fatal(err)
	}

	const codeEngineStackDefPath = "stack_definition.json"
	//const kubernetesStackDefPath = "kubernetes/stack_definition.json"

	const ceFlavour = "alm-stack-ce"
	//const kubeFlavour = "alm-stack-kube"

	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing:                t,
		ResourceGroup:          "default",
		Prefix:                 "alm", // setting prefix here gets a random string appended to it
		StackConfigurationPath: codeEngineStackDefPath,
		CatalogProductName:     "deploy-arch-ibm-alm-stack",
		CatalogFlavorName:      ceFlavour,
	})

	//fmt.Print(permanentResources["secretsManagerCRN"])

	options.StackMemberInputs = map[string]map[string]interface{}{
		"5 - Secrets Manager": {
			"existing_resource_group_name":          "default",
			"existing_secrets_manager_instance_crn": permanentResources["secretsManagerCRN"],
		},
		"6 - DevSecOps Toolchains": {
			"autostart":               "false",
			"create_cos_api_key":      "false",
			"create_secret_group":     "false",
			"create_signing_key":      "false",
			"create_ibmcloud_api_key": "false",
		},
	}

	options.StackInputs = map[string]interface{}{
		"prefix":                          options.Prefix,
		"resource_group_name":             "stack-pr-rg",
		"sm_service_plan":                 "trial",
		"use_existing_resource_group":     "true",
		"create_cd_instance":              false,
		"region":                          "us-south",
		"enable_cos":                      "false",
		"existing_secrets_manager_crn":    permanentResources["secretsManagerCRN"],
		"ibmcloud_api_key":                options.RequiredEnvironmentVars["TF_VAR_ibmcloud_api_key"], // always required by the stack
		"app_repo_existing_url":           "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
		"evidence_repo_existing_url":      "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
		"issues_repo_existing_url":        "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
		"inventory_repo_existing_url":     "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
		"cd_deployment_repo_existing_url": "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
		"change_management_existing_url":  "https://us-south.git.cloud.ibm.com/padraic.edwards/hello-compliance-app-app-repo",
	}

	err1 := options.RunProjectsTest()
	if assert.NoError(t, err1) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}
