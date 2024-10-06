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

	// 1. Locate the folder to be moved, as well as the target folder
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

	// 2. Locate the destination folder
	for _, folder := range f.folders {
		if folder.Name == dst {
			dstRoot = folder
			break
		}
	}

	// 2. Find all child folders of the folder/subtree to be moved
	// subtree = f.GetAllChildFolders(srcRoot.OrgId, srcRoot.Name)
	subtree = append([]Folder{srcRoot}, f.GetAllChildFolders(srcRoot.OrgId, srcRoot.Name)...)

	// If no child folders found, return an error
	if len(subtree) == 0 {
		return []Folder{}, fmt.Errorf("no child folders found")
	}

	// 3. Validate the destination path
	// Cannot move a folder to its own subtree
	srcPath := srcRoot.Paths
	fmt.Println("Source path: ", srcPath)
	if strings.HasPrefix(srcPath, dst) {
		return []Folder{}, fmt.Errorf("cannot move a folder to its own subtree")
	}

	// 4. Move the subtree to the destination
	for i, folder := range subtree {
		newPath := strings.Replace(folder.Paths, srcPath, dstRoot.Paths+"."+srcRoot.Name, 1)
		fmt.Println("New path: ", newPath)
		subtree[i].Paths = newPath
	}

	// 5. Update the original folder list with moved folders
	for i, folder := range f.folders {
		for _, movedFolder := range subtree {
			if folder.Name == movedFolder.Name {
				f.folders[i].Paths = movedFolder.Paths
			}
		}
	}

	// 6. Return the updated folders
	return f.folders, nil
}
