package v1

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/diff"

	v1 "github.com/openshift/api/image/v1"
	newer "github.com/openshift/openshift-apiserver/pkg/image/apis/image"
	imagev1 "github.com/openshift/openshift-apiserver/pkg/image/apis/image/v1"

	_ "github.com/openshift/openshift-apiserver/pkg/image/apis/image/install"
)

func TestImageStreamStatusConversionPreservesTags(t *testing.T) {
	in := &newer.ImageStreamStatus{
		Tags: map[string]newer.TagEventList{
			"v3.5.0": {},
			"3.5.0":  {},
		},
	}
	expOutVersioned := &v1.ImageStreamStatus{
		Tags: []v1.NamedTagEventList{{Tag: "3.5.0"}, {Tag: "v3.5.0"}},
	}

	outVersioned := v1.ImageStreamStatus{Tags: []v1.NamedTagEventList{}}
	err := imagev1.Convert_image_ImageStreamStatus_To_v1_ImageStreamStatus(in, &outVersioned, nil)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if a, e := &outVersioned, expOutVersioned; !reflect.DeepEqual(a, e) {
		t.Fatalf("got unexpected output: %s", diff.ObjectDiff(a, e))
	}

	// convert back from v1 to internal scheme
	out := newer.ImageStreamStatus{}
	err = imagev1.Convert_v1_ImageStreamStatus_To_image_ImageStreamStatus(&outVersioned, &out, nil)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if a, e := &out, in; !reflect.DeepEqual(a, e) {
		t.Fatalf("got unexpected output: %s", diff.ObjectDiff(a, e))
	}
}
