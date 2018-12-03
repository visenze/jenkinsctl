package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/emman27/jenkinsctl/pkg/builds"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// GetBuild retrieves a particular build of a job
func (c *JenkinsClient) GetBuild(jobName string, buildID int) (*builds.Build, error) {
	resp, err := c.Get(fmt.Sprintf("/job/%s/%d/api/json", jobName, buildID))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get job %d for %s", buildID, jobName)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read response body")
	}
	var build = new(builds.Build)
	json.Unmarshal(content, build)
	glog.Infof("Parsed response %v", build)
	return build, nil
}

// CreateBuild starts a build in Jenkins
func (c *JenkinsClient) CreateBuild(jobName string, params map[string]interface{}) {
	reader := bytes.NewReader([]byte{})
	c.Post(fmt.Sprintf("/job/%s/build", jobName), reader)
}
