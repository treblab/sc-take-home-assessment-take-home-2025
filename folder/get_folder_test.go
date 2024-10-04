package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// Load data from sample.json using GetSampleData
	driver := folder.NewDriver(folder.GetSampleData())
	org1 := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")) // Using default orgID from sample data

	tests := [...]struct {
		name   string
		orgID  uuid.UUID
		folder string // changed to string - was originally []folder.Folder
		want   []folder.Folder
	}{
		{
			// TODO: your tests here
			// Type 1: Test for root folder with children
			name:  "Root folder with children",
			orgID: org1,
			want: []folder.Folder{
				{Name: "Bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "Charlie", OrgId: org1, Paths: "alpha.charlie"},
			},
		},
		{
			// Type 2: Test for root folder with no children
			name:   "Root folder with no children",
			orgID:  org1,
			folder: "delta",
			want:   []folder.Folder{},
		},
		{
			// Type 3: Null folder
			name:   "Null folder",
			orgID:  org1,
			folder: "null",
			want:   []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)
			got := driver.GetAllChildFolders(tt.orgID, tt.folder)
			assert.Equal(t, tt.want, got)
		})
	}
}
