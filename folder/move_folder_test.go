package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var org1 = uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))
var org2 = uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a8"))

// Sample custom data for testing
func getCustomSampleData() []folder.Folder {

	return []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: org1},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2},
	}
}

func Test_folder_MoveFolder(t *testing.T) {
	driver := folder.NewDriver(getCustomSampleData())

	tests := []struct {
		name          string
		src           string
		dst           string
		expectedError string
		orgID         uuid.UUID
		want          []folder.Folder
	}{
		{
			name: "Move folder with children within same org",
			src:  "bravo",
			dst:  "delta",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.delta.bravo"},           // Moved folder
				{Name: "charlie", OrgId: org1, Paths: "alpha.delta.bravo.charlie"}, // Moved child folder
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
		},
		// {
		// 	name:          "Move folder to its own child",
		// 	src:           "alpha.bravo",
		// 	dst:           "alpha.bravo.charlie",
		// 	expectedError: "cannot move a folder to its own subtree",
		// },
		// {
		// 	name:          "Move folder to itself",
		// 	src:           "alpha.bravo",
		// 	dst:           "alpha.bravo",
		// 	expectedError: "cannot move a folder to itself",
		// },
		// {
		// 	name:          "Move folder across different orgID",
		// 	src:           "alpha.bravo",
		// 	dst:           "foxtrot",
		// 	expectedError: "cannot move a folder to a different organization",
		// },
		// {
		// 	name:          "Invalid source folder",
		// 	src:           "invalid_folder",
		// 	dst:           "alpha.delta",
		// 	expectedError: "folder not found",
		// },
		// {
		// 	name:          "Invalid destination folder",
		// 	src:           "alpha.bravo",
		// 	dst:           "invalid_folder",
		// 	expectedError: "destination folder not found",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := driver.MoveFolder(tt.src, tt.dst)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
