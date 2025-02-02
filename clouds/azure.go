package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectAzure() (is bool) {
	resp, err := util.Client().Get("http://169.254.169.254/metadata/v1/InstanceInfo")
	if err == nil && resp.StatusCode == http.StatusOK {
		is = true
		return
	}
	return
}
