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
	org1 := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")) // Using first orgID from sample data
	org2 := uuid.Must(uuid.FromString("b1556e17-b7c0-45a3-a6ae-9546248fb17b")) // Incorrect orgID

	tests := [...]struct {
		name   string
		orgID  uuid.UUID
		folder string // changed to string - was originally []folder.Folder
		want   []folder.Folder
	}{
		{
			// TODO: your tests here
			name:   "Root folder with children - 1",
			orgID:  org1,
			folder: "enabled-professor-monster",
			want: []folder.Folder{
				{Name: "glowing-elongated", OrgId: org1, Paths: "stunning-horridus.sacred-moonstar.nearby-maestro.enabled-professor-monster.glowing-elongated"},
				{Name: "equipped-hypno-hustler", OrgId: org1, Paths: "stunning-horridus.sacred-moonstar.nearby-maestro.enabled-professor-monster.equipped-hypno-hustler"},
			},
		},
		{
			name:   "Root folder with no children - 2",
			orgID:  org1,
			folder: "equipped-hypno-hustler",
			want:   []folder.Folder(nil),
		},
		{
			name:   "Null folder - 3",
			orgID:  org1,
			folder: "null",
			want:   []folder.Folder(nil),
		},
		{
			name:   "Folder with incorrect orgID - 4",
			orgID:  org2,
			folder: "enabled-professor-monster",
			want:   []folder.Folder(nil),
		},
		{
			name:   "No folder name provided - 5",
			orgID:  org1,
			folder: "",
			want:   []folder.Folder(nil),
		},
		{
			name:   "Case sensitivity check - 6",
			orgID:  org1,
			folder: "Enabled-Professor-Monster",
			want:   []folder.Folder(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := driver.GetAllChildFolders(tt.orgID, tt.folder)
			assert.Equal(t, tt.want, got)
		})
	}
}
