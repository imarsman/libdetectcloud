package libdetectcloud

import (
	"net/http"
	"time"

	"github.com/imarsman/libdetectcloud/clouds"
)

const (
	aws                 = "Amazon Web Services"
	azure               = "Microsoft Azure"
	digitalOcean        = "Digital Ocean"
	googleComputeEngine = "Google Compute Engine"
	openStackFoundation = "OpenStack Foundation"
	softLayer           = "SoftLayer"
	vultr               = "Vultr"
	container           = "Container"
)

var hc = &http.Client{Timeout: 300 * time.Millisecond}

// cloudTypes type
// type cloudTypes struct {
// 	aws       string
// 	azure     string
// 	do        string
// 	gce       string
// 	ost       string
// 	sl        string
// 	vr        string
// 	container string
// }

func IsOnAWS() bool {
	return clouds.DetectAWS() == aws
}

func IsOnAzure() bool {
	return clouds.DetectAzure() == azure
}

func IsOnDigitalOcean() bool {
	return clouds.DetectDigitalOcean() == digitalOcean
}

func IsOnGce() bool {
	return clouds.DetectGCE() == googleComputeEngine
}

func IsOnKube() bool {
	return clouds.DetectContainer() == container
}

func IsOnOpenstack() bool {
	return clouds.DetectOpenStack() == openStackFoundation
}

func IsOnSoftlayer() bool {
	return clouds.DetectSoftlayer() == softLayer
}

func IsOnVultr() bool {
	return clouds.DetectVultr() == vultr
}

// // Detect function
// func Detect() string {
// 	// if runtime.GOOS != "darwin" {
// 	var c cloudTypes
// 	var wg sync.WaitGroup
// 	wg.Add(8)
// 	go func() {
// 		defer wg.Done()
// 		c.aws = clouds.DetectAWS()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.azure = clouds.DetectAzure()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.do = clouds.DetectDigitalOcean()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.gce = clouds.DetectGCE()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.ost = clouds.DetectOpenStack()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.sl = clouds.DetectSoftlayer()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.vr = clouds.DetectVultr()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		c.container = clouds.DetectContainer()
// 	}()
// 	wg.Wait()

// 	if c.aws != "" {
// 		return c.aws
// 	}
// 	if c.azure != "" {
// 		return c.azure
// 	}
// 	if c.do != "" {
// 		return c.do
// 	}
// 	if c.gce != "" {
// 		return c.gce
// 	}
// 	if c.ost != "" {
// 		return c.ost
// 	}
// 	if c.sl != "" {
// 		return c.sl
// 	}
// 	if c.vr != "" {
// 		return c.vr
// 	}
// 	if c.container != "" {
// 		return c.container
// 	}
// 	// }
// 	return ""
// }
