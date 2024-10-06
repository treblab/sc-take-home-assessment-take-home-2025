package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	// 1 - Filter folders by the given orgID
	folders := f.GetFoldersByOrgID(orgID)

	// 2 - Filter the root by the given name
	var rootFolder Folder

	for _, folder := range folders {
		if folder.Name == name {
			rootFolder = folder
			break
		}
	}

	// If no folder found, return empty list
	if rootFolder.Name == "" {
		return []Folder{}
	}

	// 3 - Filter the children of the root folder
	var children []Folder
	for _, folder := range folders {
		// Check if folder.Paths starts with rootFolder.Paths + "."
		if folder.Paths != rootFolder.Paths && strings.HasPrefix(folder.Paths, rootFolder.Paths+".") {
			children = append(children, folder)
		}
	}
	return children
}
