package clouds

import (
	"io/ioutil"
	"strings"
)

func DetectContainer() bool {
	b, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return false
	}

	fc := string(b)
	kube := strings.Contains(fc, "kube")
	container := strings.Contains(fc, "containerd")

	if kube {
		return true
	}

	if container {
		return true
	}

	return false
}
