package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectDigitalOcean() (is bool) {
	resp, err := util.Client().Get("http://169.254.169.254/metadata/v1.json")
	if err == nil && resp.StatusCode == http.StatusOK {
		is = true
		return
	}
	return
}
