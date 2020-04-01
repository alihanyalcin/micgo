package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *project) create() {
	fmt.Println("Starting to generate project", p.name)
	// Create project directory
	err := os.Mkdir(p.name, 0755)
	checkError(err)

	// create directories
	for _, path := range directories {
		err := os.MkdirAll(p.name+path, 0775)
		checkError(err)
	}
	fmt.Println("Directories created.")

	for _, f := range files {
		sfile, err := ioutil.ReadFile(basePath + f)
		checkError(err)

		dfile := strings.Replace(string(sfile), "project", p.name, -1)
		err = ioutil.WriteFile(p.name+f, []byte(dfile), 0644)
		checkError(err)
		fmt.Println(p.name + f + " created.")
	}
}
