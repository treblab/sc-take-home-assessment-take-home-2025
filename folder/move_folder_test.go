package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	driver := folder.NewDriver(folder.GetSampleData())

	tests := [...]struct {
		name          string
		src           string
		dst           string
		expectedError string
		expectedPaths []string
	}{
		{
			name: "1 --- Move folder WITH children -------------------------------------------- 1",
			src:  "stunning-horridus",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "2 --- Move folder with NO children -------------------------------------------- 2",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "3 --- Move folder to itself/own subtree -------------------------------------------- 3",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "4 --- Invalid source folder -------------------------------------------- 4",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "5 --- Invalid destination folder -------------------------------------------- 5",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "6 --- Move folder across different orgID -------------------------------------------- 6",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "7 --- Move parent folder to its own child -------------------------------------------- 7",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
		{
			name: "8 --- Check if order of children has been maintained -------------------------------------------- 8",
			src:  "Invalid-folder",
			dst:  "new-destination",
			expectedPaths: []string{
				"new-destination",
				"new-destination.sacred-moonstar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// To ensure tt is correctly captured within the closure
			tt := tt

			got, err := driver.MoveFolder(tt.src, tt.dst)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPaths, got)
			}
		})
	}
}
