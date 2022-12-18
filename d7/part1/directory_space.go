package main

import (
	"bufio"
	"fmt"
	"os"
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

func getDirectoriesLte100000(node *Node, sum int) int {
	for _, childNode := range node.children {
		if !childNode.isDirectory {
			continue
		}

		if childNode.fileSize <= 100000 {
			sum += getDirectoriesLte100000(childNode, childNode.fileSize)
		} else if childNode.isDirectory {
			sum += getDirectoriesLte100000(childNode, 0)
		}
	}

	return sum
}

func main() {
	rootNode := parseAndBuildDirectoryTree()
	fillDirectorySizes(&rootNode)
	answer := getDirectoriesLte100000(&rootNode, 0)
	fmt.Println(answer)
}