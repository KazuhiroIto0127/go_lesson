package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func makeDir(dirname string){
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}

	fmt.Printf("Directory '%s' created.\n", dirname)
}

func createTxtfile(txtFileName string) {
	file, err := os.Create(txtFileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	fmt.Printf("File '%s' created.\n", txtFileName)
}

func removeDir(dirname string) {
	err := os.RemoveAll(dirname)
	if err != nil {
		log.Fatal("Error removing directory:", err)
	}
	fmt.Printf("Directory '%s' and its contents removed.\n", dirname)
}

func printAllCurrentDirFiles(){
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir(){
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range files{
		fmt.Println(key, val)
	}
}

func main(){
	const dirName = "sampleTmp"
	makeDir(dirName)
	for i := 1; i<=10; i++ {
		filename := dirName + "/tmp" + strconv.Itoa(i) + ".txt"
		createTxtfile(filename)
	}
	printAllCurrentDirFiles()
	removeDir(dirName)
}
