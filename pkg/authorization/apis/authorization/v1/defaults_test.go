package v1

import (
	"reflect"
	"testing"

	runtime "k8s.io/apimachinery/pkg/runtime"

	v1 "github.com/openshift/api/authorization/v1"
	authorizationapi "github.com/openshift/openshift-apiserver/pkg/authorization/apis/authorization"
)

var scheme = runtime.NewScheme()

func init() {
	Install(scheme)
	authorizationapi.Install(scheme)
}

func TestDefaults(t *testing.T) {
	obj := &v1.PolicyRule{
		APIGroups: nil,
		Verbs:     []string{authorizationapi.VerbAll},
		Resources: []string{authorizationapi.ResourceAll},
	}
	out := &authorizationapi.PolicyRule{}
	if err := scheme.Convert(obj, out, nil); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(out.APIGroups, []string{authorizationapi.APIGroupAll}) {
		t.Errorf("did not default api groups: %#v", out)
	}
}
