package libdetectcloud

import (
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
	k8SContainer        = "K8S Container"
)

func IsOnAWS() bool {
	return clouds.DetectAWS()
}

func IsOnAzure() bool {
	return clouds.DetectAzure()
}

func IsOnDigitalOcean() bool {
	return clouds.DetectDigitalOcean()
}

func IsOnGce() bool {
	return clouds.DetectGCE()
}

func IsOnKube() bool {
	return clouds.DetectContainer()
}

func IsOnOpenstack() bool {
	return clouds.DetectOpenStack()
}

func IsOnSoftlayer() bool {
	return clouds.DetectSoftlayer()
}

func IsOnVultr() bool {
	return clouds.DetectVultr()
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
