package libdetectcloud

import (
	"testing"
)

// func TestCloud(t *testing.T) {
// 	c := Detect()
// 	if c == Gce {
// 		t.Log("Gce")
// 	}
// 	t.Logf("client .%s.", c)

// }

func TestGCE(t *testing.T) {
	t.Log("IsOnGCE", IsOnGce())
}
