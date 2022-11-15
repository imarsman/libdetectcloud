package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectAWS() bool {
	resp, err := util.Client().Get("http://169.254.169.254/latest/")
	if err == nil && resp.StatusCode == http.StatusOK {
		return false
	}
	return true
}
