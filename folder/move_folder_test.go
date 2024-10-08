package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Set 2 different orgIDs for testing
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
	tests := []struct {
		name          string
		src           string
		dst           string
		expectedError string
		orgID         uuid.UUID
		want          []folder.Folder
	}{
		{
			name: "Move folder with children within same org - 1",
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
		{
			name: "Move folder to itself - 2",
			src:  "bravo",
			dst:  "bravo",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"}, // No move due to invalid operation
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "cannot move a folder to itself",
		},
		{
			name: "Move folder across different orgID - 3",
			src:  "bravo",
			dst:  "foxtrot",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"}, // No move due to cross-org operation
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "cannot move a folder to a different organisation",
		},
		{
			name: "Invalid source folder - 4",
			src:  "invalid_folder",
			dst:  "delta",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"}, // No move due to invalid source
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "source folder not found",
		},
		{
			name: "Invalid destination folder - 5",
			src:  "bravo",
			dst:  "invalid_folder",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"}, // No move due to invalid destination
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "destination folder not found",
		},
		{
			name: "Move folder into its own descendant - 6",
			src:  "bravo",
			dst:  "charlie",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"}, // No move due to invalid operation
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "cannot move a folder to its own subtree",
		},
		{
			name: "Move folder without children - 7",
			src:  "echo",
			dst:  "bravo",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.bravo.echo"}, // Moved echo to bravo
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "",
		},
		{
			name: "Move folder to its own parent (no-op) - 8",
			src:  "charlie",
			dst:  "bravo",
			want: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"}, // No move because it's already under bravo
				{Name: "delta", OrgId: org1, Paths: "alpha.delta"},
				{Name: "echo", OrgId: org1, Paths: "alpha.delta.echo"},
				{Name: "foxtrot", OrgId: org2, Paths: "foxtrot"},
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reinitialize the driver for each test case
			driver := folder.NewDriver(getCustomSampleData())

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
