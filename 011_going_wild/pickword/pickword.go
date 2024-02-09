package pickword

import (
    "log"
	"math/rand"
	"os"
	"bufio"
)

// Hello returns a greeting for the named person.
func Pick() string {
	dictPath := "../dict"

	chosenFile, err := choseOneFile(dictPath)

	if err != nil {
        log.Fatal(err)
    }

	word, err := choseOneWord(chosenFile)

	if err != nil {
        log.Fatal(err)
    }

	return word
}

//chose one of the files from the given path
func choseOneFile(path string) (string, error){
	files, err := os.ReadDir(path)
    if err != nil {
        return "", err
    }
	chosenFile := path + "/" + files[rand.Intn(len(files))].Name()
	return chosenFile, nil
}

//chose one of the words from the given file
func choseOneWord(path string) (string, error) {
	readFile, err := os.Open(path)
  
    if err != nil {
        return "", err
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()

	return fileLines[rand.Intn(len(fileLines))], nil
}