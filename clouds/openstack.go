package clouds

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func DetectOpenStack() (is bool) {
	if runtime.GOOS != "windows" {
		data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			return
		}
		if strings.Contains(string(data), "OpenStack Foundation") {
			is = true
			return
		}
		return
	}
	return
}
