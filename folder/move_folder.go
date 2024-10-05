package folder

import (
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...

	var srcRoot Folder
	var subtree []Folder

	// 1. Locate the folder to be moved
	for _, folder := range f.folders {
		if folder.Name == name {
			srcRoot = folder
			break
		}
	}

	// If the folder is not found, return an error
	if srcRoot.Name == "" {
		return []Folder{}, fmt.Errorf("folder not found")
	}

	// 2. Find all child folders of the folder/subbtree to be moved
	srcPath := srcRoot.Paths
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, srcPath+".") {
			subtree = append(subtree, folder)
		}
	}

	// If no child folders found, return an error
	if len(subtree) == 0 {
		return []Folder{}, fmt.Errorf("no child folders found")
	}

	// 3. Validate the destination path
	// Cannot move a folder to its own subtree
	if strings.HasPrefix(srcPath, dst) {
		return []Folder{}, fmt.Errorf("cannot move a folder to its own subtree")
	}

	// 4. Move the subtree to the destination
	for i, folder := range subtree {
		newPath := strings.Replace(folder.Paths, srcPath, dst, 1)
		subtree[i].Paths = newPath
	}

	// 5. Return the updated folders
	return f.folders, nil
}
