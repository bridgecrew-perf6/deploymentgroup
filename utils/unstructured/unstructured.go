package unstructured

import (
	"regexp"

	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/ssanders1449/deploymentgroup/pkg/apis/deploymentgroup/v1alpha1"
)

func StrToUnstructuredUnsafe(jsonStr string) *unstructured.Unstructured {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{Object: obj}
}

func StrToUnstructured(jsonStr string) (*unstructured.Unstructured, error) {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: obj}, nil
}

func ObjectToDeploymentGroup(obj interface{}) *v1alpha1.DeploymentGroup {
	un, ok := obj.(*unstructured.Unstructured)
	if ok {
		var ro v1alpha1.DeploymentGroup
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(un.Object, &ro)
		if err != nil {
			log.Warnf("Failed to convert DeploymentGroup from Unstructured object: %v", err)
			return nil
		}
		return &ro
	}
	ro, ok := obj.(*v1alpha1.DeploymentGroup)
	if !ok {
		log.Warnf("Object is neither a rollout or unstructured: %v", obj)
	}
	return ro
}
