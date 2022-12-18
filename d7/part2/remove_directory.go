package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	isDirectory bool
	name string
	fileSize int
	parent *Node
	children map[string]*Node
}

var directories []int

func parseAndBuildDirectoryTree() Node {
	fileName := os.Args[2]
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	var currentDirectory *Node

	rootNode := Node {
		isDirectory: true,
		name: "/",
		fileSize: 0,
		children: make(map[string]*Node),
	}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		isCommand := splitLine[0] == "$"

		if isCommand {
			isCdCommand := len(splitLine) >= 2 && splitLine[1] == "cd"
			if isCdCommand {
				target := splitLine[2]

				if target == "/" {
					currentDirectory = &rootNode
				} else if target == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					currentDirectory = currentDirectory.children[target]
				}
			}
		} else {
			fileSize, fileName := splitLine[0], splitLine[1]
			castedFileSize, _ := strconv.Atoi(fileSize)
			var node Node

			if splitLine[0] == "dir" {
				node = Node {
					isDirectory: true,
					name: fileName,
					fileSize: 0,
					parent: currentDirectory,
					children: make(map[string]*Node),
				}
			} else {
				node = Node {
					name: fileName,
					fileSize: castedFileSize,
					parent: currentDirectory,
					children: make(map[string]*Node),
				}
			}
			currentDirectory.children[fileName] = &node
		}
	}
	return rootNode
} 

func fillDirectorySizes(node *Node) int {
	for _, childNode := range node.children {
		node.fileSize += fillDirectorySizes(childNode)
	}

	return node.fileSize
}

func getDirectories(requiredSpace int, node *Node) []int {
	for _, childNode := range node.children {
		if !childNode.isDirectory {
			continue
		}

		getDirectories(requiredSpace, childNode)

		if childNode.fileSize >= requiredSpace {
			directories = append(directories, childNode.fileSize)
		}
	}
	return directories
}

func main() {
	TOTAL_SPACE := 70000000
	UPDATE_SIZE := 30000000

	// create directory tree
	rootNode := parseAndBuildDirectoryTree()
	fillDirectorySizes(&rootNode)
	
	// calculate necessary space
	requiredSpace := UPDATE_SIZE - (TOTAL_SPACE - rootNode.fileSize)

	// populate directories slice
	getDirectories(requiredSpace, &rootNode)

	// sort directories
	sort.Slice(directories, func(i, j int) bool {
		return directories[i] < directories[j]
	})

	fmt.Println(directories[0])
}