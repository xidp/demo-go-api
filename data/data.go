package data

import (
	"bufio"
	"fmt"
	"os"
)

type vData struct {
	Version     string
	GitSHA      string
	GitRepo     string
	GitProvider string
}

func new(name string) *vData {
	var res = &vData{}

	switch name {
	case "app":
		res.Version = "v0.1.0"
		res.GitSHA = "1b908c71f2f8b611743c8019db3c70ae2fb95688"
		res.GitRepo = "gabtec/demo-go-api"
		res.GitProvider = "github.com"
	case "github":
		res.Version = "v0.1.0"
		res.GitSHA = "c4ab8b471cc71903757d19895e29b2bea1da3d38"
		res.GitRepo = "xgeekshq/sie-az-prototype"
		res.GitProvider = "github.com"
	case "gitlab":
		res.Version = "v1.2.3"
		res.GitSHA = "1fe5fbec3d7bb7c83e598240cac3d81731e59c65"
		res.GitRepo = "ki-group-pt/xgeekshq/external/sx-frontend"
		res.GitProvider = "gitlab.com"
	}

	return res
}

func WriteToData(fileName, version, sha string) error {
	f := "./data/" + fileName + ".txt"
	// Modify the content (replace with new content)
	contentString := sha + "\n" + version + "\n"
	newContent := []byte(contentString)

	// Write the updated content back to the file
	err := os.WriteFile(f, newContent, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("File '%s' has been updated with new content.\n", f)
	return nil
}

func ReadData(fileName string) *vData {
	f := "./data/" + fileName + ".txt"

	// defaults
	var res *vData
	res = new(fileName)

	// Open a file for reading
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return res
	}
	defer file.Close()

	// Create a bufio.Reader
	reader := bufio.NewReader(file)

	var lines []string
	// Read lines from the file using ReadLine
	for {
		line, isPrefix, err := reader.ReadLine()

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Finish reading.")
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}

		// Check if the line is too long to fit in the buffer
		if isPrefix {
			fmt.Println("Line is too long and was truncated.")
			continue
		}

		// Process the line
		fmt.Println("Line:", string(line))
		lines = append(lines, string(line))
	}

	res.GitSHA = lines[0]
	res.Version = lines[1]

	return res
}
