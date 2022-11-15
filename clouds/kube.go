package clouds

import (
	"io/ioutil"
	"strings"
)

func DetectContainer() (is bool) {
	b, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return false
	}

	fc := string(b)
	kube := strings.Contains(fc, "kube")
	container := strings.Contains(fc, "containerd")

	if kube {
		is = true
		return
	}

	if container {
		is = true
		return
	}

	return
}
