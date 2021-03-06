package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
	rbacv1conversions "k8s.io/kubernetes/pkg/apis/rbac/v1"

	v1 "github.com/openshift/api/authorization/v1"
	"github.com/openshift/openshift-apiserver/pkg/authorization/apis/authorization"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(
		authorization.Install,
		v1.Install,
		rbacv1conversions.AddToScheme,
		corev1conversions.AddToScheme,
		AddFieldSelectorKeyConversions,
		RegisterDefaults,
	)
	Install = localSchemeBuilder.AddToScheme
)
