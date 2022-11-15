package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectGCE() (is bool) {
	r, err := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/tags", nil)
	if err != nil {
		return
	}
	r.Header.Add("Metadata-Flavor", "Google")
	resp, err := util.Client().Do(r)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		is = true
		return
	}
	return
}
