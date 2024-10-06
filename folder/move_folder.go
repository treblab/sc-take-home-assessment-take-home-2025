package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	var srcRoot Folder
	var dstRoot Folder
	var subtree []Folder

	// 1. Locate the paths of folder to be moved (src) and the dest folder
	for _, folder := range f.folders {
		if folder.Name == name {
			srcRoot = folder
		}
		if folder.Name == dst {
			dstRoot = folder
		}
		// Break early if both folders are found
		if srcRoot.Name != "" && dstRoot.Name != "" {
			// fmt.Println(srcRoot.Name, dstRoot.Name)
			break
		}
	}

	// Base cases:

	// If the source folder is the same as the destination folder, return an error
	if srcRoot.Paths == dstRoot.Paths {
		// fmt.Println("Cannot move a folder to itself")
		return f.folders, fmt.Errorf("cannot move a folder to itself")
	}

	// If the source folder is not found, return an error
	if srcRoot.Name == "" {
		return []Folder{}, fmt.Errorf("source folder not found")
	}

	// If the destination folder is not found, return an error (if required)
	if dstRoot.Name == "" {
		return []Folder{}, fmt.Errorf("destination folder not found")
	}

	// 2. Find all child folders of the folder/subtree to be moved
	subtree = append([]Folder{srcRoot}, f.GetAllChildFolders(srcRoot.OrgId, srcRoot.Name)...)

	// If no child folders found, return a message
	if len(subtree) == 0 {
		fmt.Println("No child folders found")
	}

	// 3. Move the subtree to the destination
	for i, folder := range subtree {
		newPath := strings.Replace(folder.Paths, srcRoot.Paths, dstRoot.Paths+"."+srcRoot.Name, 1)
		subtree[i].Paths = newPath
	}

	// 4. Update the original folder list with moved folders
	for i, folder := range f.folders {
		for _, movedFolder := range subtree {
			if folder.Name == movedFolder.Name {
				f.folders[i].Paths = movedFolder.Paths
			}
		}
	}

	// 5. Return the updated folders
	return f.folders, nil
}
