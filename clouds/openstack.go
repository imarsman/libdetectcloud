package clouds

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func DetectOpenStack() bool {
	if runtime.GOOS != "windows" {
		data, err := ioutil.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			return false
		}
		if strings.Contains(string(data), "OpenStack Foundation") {
			return true
		}
		return false
	}
	return false
}
