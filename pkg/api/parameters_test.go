package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/visenze/jenkinsctl/pkg/builds"
)

func Test_getParamsErrorIfNoParams(t *testing.T) {
	build := builds.Build{
		Actions: []builds.BuildAction{},
	}
	_, err := getParams(&build)
	assert.NotNil(t, err)
}

func Test_getParams(t *testing.T) {
	build := builds.Build{
		Actions: []builds.BuildAction{
			{Class: parameterActionClass, Parameters: &[]builds.BuildParameter{}},
		},
	}
	params, err := getParams(&build)
	assert.Nil(t, err)
	assert.NotNil(t, params)
}
