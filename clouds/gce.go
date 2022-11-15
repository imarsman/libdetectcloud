package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectGCE() bool {
	r, err := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/tags", nil)
	if err != nil {
		return false
	}
	r.Header.Add("Metadata-Flavor", "Google")
	resp, err := util.Client().Do(r)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
