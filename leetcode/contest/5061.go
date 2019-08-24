package contest

import "strings"

type FileSystem struct {
	HeadNode *TreeNode22
}

type TreeNode22 struct {
	Value int
	Next  map[string]*TreeNode22
}

func Constructor() FileSystem {
	return FileSystem{
		HeadNode: &TreeNode22{
			Value: 0,
			Next:  make(map[string]*TreeNode22),
		},
	}
}

func (this *FileSystem) Create(path string, value int) bool {
	split := strings.Split(path, "/")
	if len(split) == 1 && len(split[0]) == 0 {
		return false
	}
	split = split[1:]

	nowNode := this.HeadNode
	for key, folderName := range split {
		node, ok := nowNode.Next[folderName]
		if !ok {
			if key != len(split)-1 {
				return false
			} else {
				nowNode.Next[folderName] = &TreeNode22{value, map[string]*TreeNode22{}}
				nowNode = nowNode.Next[folderName]
			}
		} else {
			nowNode = node
		}

	}
	return true
}

func (this *FileSystem) Get(path string) int {
	split := strings.Split(path, "/")
	if len(split) == 1 && len(split[0]) == 0 {
		return -1
	}
	split = split[1:]

	nowNode := this.HeadNode
	for key, folderName := range split {
		node, ok := nowNode.Next[folderName]
		if !ok {
			return -1
		}

		if key == len(split)-1 {
			return nowNode.Next[folderName].Value
		} else {
			nowNode = node

		}
	}
	return -1
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Create(path,value);
 * param_2 := obj.Get(path);
 */
