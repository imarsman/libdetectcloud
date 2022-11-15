package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectAWS() string {
	resp, err := util.Client().Get("http://169.254.169.254/latest/")
	if err == nil && resp.StatusCode == http.StatusOK {
		return "Amazon Web Services"
	}
	return ""
}
