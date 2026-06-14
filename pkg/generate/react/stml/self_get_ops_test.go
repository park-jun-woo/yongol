//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what selfGetOps — Delete<Entity>→Get<Entity> 이름짝+path-param 매칭, 비delete/무param/미fetch/시블링 제외 검증

package stml

import (
	"reflect"
	"testing"
)

func TestSelfGetOps(t *testing.T) {
	pathTypes := map[string]map[string]string{
		"DeleteBuilding":         {"buildingId": "integer"},
		"GetBuilding":            {"buildingId": "string"}, // types ignored, names match
		"CheckBuildingDeletable": {"buildingId": "integer"},
		"DeleteRoom":             {}, // no path params
	}

	tests := []struct {
		name       string
		deleteOpID string
		fetchOps   []string
		want       []string
	}{
		{
			name:       "self GET matched by name and path params",
			deleteOpID: "DeleteBuilding",
			fetchOps:   []string{"GetBuilding", "CheckBuildingDeletable"},
			want:       []string{"GetBuilding"},
		},
		{
			name:       "self GET not fetched on page",
			deleteOpID: "DeleteBuilding",
			fetchOps:   []string{"CheckBuildingDeletable", "ListBuildingPhotos"},
			want:       nil,
		},
		{
			name:       "opID without Delete prefix",
			deleteOpID: "ArchiveBuilding",
			fetchOps:   []string{"GetBuilding"},
			want:       nil,
		},
		{
			name:       "delete with no path params",
			deleteOpID: "DeleteRoom",
			fetchOps:   []string{"GetRoom"},
			want:       nil,
		},
		{
			name:       "bare Delete with empty entity",
			deleteOpID: "Delete",
			fetchOps:   []string{"Get"},
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selfGetOps(tt.deleteOpID, tt.fetchOps, pathTypes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfGetOps() = %v, want %v", got, tt.want)
			}
		})
	}
}
