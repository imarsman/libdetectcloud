package clouds

import (
	"net/http"

	"github.com/imarsman/libdetectcloud/internal/util"
)

func DetectSoftlayer() (is bool) {
	resp, err := util.Client().Get("https://api.service.softlayer.com/rest/v3/SoftLayer_Resource_Metadata/UserMetadata.txt")
	if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound) {
		is = true
		return
	}
	return
}
