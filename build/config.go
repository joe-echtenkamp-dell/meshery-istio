package build

import (
	"os"
	"path/filepath"

	"github.com/layer5io/meshery-adapter-library/adapter"

	"github.com/layer5io/meshkit/utils"
	"github.com/layer5io/meshkit/utils/manifests"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

var DefaultGenerationMethod string
var DefaultGenerationURL string
var LatestVersion string
var WorkloadPath string
var AllVersions []string

const Component = "Istio"

//NewConfig creates the configuration for creating components
func NewConfig(version string) manifests.Config {
	return manifests.Config{
		Name:        smp.ServiceMesh_Type_name[int32(smp.ServiceMesh_ISTIO)],
		Type:        Component,
		MeshVersion: version,
		Filter: manifests.CrdFilter{
			RootFilter:    []string{"$[?(@.kind==\"CustomResourceDefinition\")]"},
			NameFilter:    []string{"$..[\"spec\"][\"names\"][\"kind\"]"},
			VersionFilter: []string{"$[0]..spec.versions[0]"},
			GroupFilter:   []string{"$[0]..spec"},
			SpecFilter:    []string{"$[0]..openAPIV3Schema.properties.spec"},
			ItrFilter:     []string{"$[?(@.spec.names.kind"},
			ItrSpecFilter: []string{"$[?(@.spec.names.kind"},
			VField:        "name",
			GField:        "group",
		},
	}
}
func init() {
	wd, _ := os.Getwd()
	WorkloadPath = filepath.Join(wd, "templates", "oam", "workloads")
	AllVersions, _ = utils.GetLatestReleaseTagsSorted("istio", "istio")
	if len(AllVersions) == 0 {
		return
	}
	LatestVersion = AllVersions[len(AllVersions)-1]
	DefaultGenerationMethod = adapter.Manifests
	DefaultGenerationURL = "https://raw.githubusercontent.com/istio/istio/" + LatestVersion + "/manifests/charts/base/crds/crd-all.gen.yaml"
}
