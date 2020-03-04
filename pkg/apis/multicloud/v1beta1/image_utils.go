// IBM Confidential
// OCO Source Materials
// (C) Copyright IBM Corporation 2019, 2020 All Rights Reserved
// The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.

package v1beta1

import (
	"fmt"

	"github.com/open-cluster-management/endpoint-operator/pkg/image"
)

var defaultComponentImageMap = map[string]string{
	"cert-manager-acmesolver":        "cert-manager-acmesolver",
	"cert-manager-controller":        "cert-manager-controller",
	"component-operator":             "endpoint-component-operator",
	"configmap-reload":               "donotexist",
	"connection-manager":             "multicloud-manager",
	"coredns":                        "donotexist",
	"curl":                           "donotexist",
	"deployable":                     "multicloud-operators-deployable",
	"policy-controller":              "mcm-compliance",
	"prometheus":                     "img-prometheus",
	"prometheus-config-reloader":     "donotexist",
	"prometheus-operator":            "img-prometheus-operator",
	"prometheus-operator-controller": "prometheus-operator-controller",
	"router":                         "management-ingress",
	"search-collector":               "search-collector",
	"service-registry":               "donotexist",
	"subscription":                   "multicloud-operators-subscription-release",
	"topology-collector":             "weave-collector",
	"work-manager":                   "multicloud-manager",
	"weave":                          "mcm-weavescope",
}

var defaultComponentTagMap = map[string]string{
	"cert-manager-acmesolver":        "0.10.1",
	"cert-manager-controller":        "0.10.1",
	"component-operator":             "1.0.0",
	"configmap-reload":               "0.0.0",
	"connection-manager":             "0.0.1",
	"coredns":                        "0.0.0",
	"curl":                           "0.0.0",
	"deployable":                     "3.5.0",
	"policy-controller":              "3.6.0",
	"prometheus":                     "absent",
	"prometheus-config-reloader":     "absent",
	"prometheus-operator":            "absent",
	"prometheus-operator-controller": "absent",
	"router":                         "1.0.0",
	"search-collector":               "3.5.0",
	"service-registry":               "0.0.0",
	"subscription":                   "3.5.0",
	"topology-collector":             "3.6.0",
	"work-manager":                   "0.0.1",
	"weave":                          "3.6.0",
}

// GetImage returns the image.Image for the specified component return error if information not found
func (instance Endpoint) GetImage(component string) (image.Image, error) {
	img := image.Image{
		PullPolicy: instance.Spec.ImagePullPolicy,
	}

	if instance.Spec.ImageRegistry != "" {
		img.Repository = instance.Spec.ImageRegistry + "/"
	}

	if imageName, ok := defaultComponentImageMap[component]; ok {
		img.Repository = img.Repository + imageName
	} else {
		return img, fmt.Errorf("unable to locate default image name for component %s", component)
	}

	if instance.Spec.ImageNamePostfix != "" {
		img.Repository = img.Repository + instance.Spec.ImageNamePostfix
	}

	if len(instance.Spec.ComponentsImagesTag) > 0 {
		if tag, ok := instance.Spec.ComponentsImagesTag[component]; ok {
			img.Tag = tag
		} // else {
		// TODO how to log - WARN("unable to locate tag for component %s", component)
		// don't want to add new dependencies to other projects importing this package
		//}
	}
	if img.Tag == "" {
		if tag, ok := defaultComponentTagMap[component]; ok {
			img.Tag = tag
		} else {
			return img, fmt.Errorf("unable to locate default tag for component %s", component)
		}
	}

	return img, nil
}
