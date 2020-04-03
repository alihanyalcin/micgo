package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *project) create() {
	fmt.Println("Starting to generate project", p.name)

	// create directories and files
	p.walk("internal", "/pkg")
	p.walk("internal", "/service")
	p.walk("cmd", "/service")
	// create single files
	for file, path := range files {
		err := p.createFilesAndDirectories(file, path, "", "")
		checkError(err)
	}
	fmt.Println("Completed.")
}

func (p *project) walk(key, dir string) {
	service := make(map[string]int)
	if strings.Contains(dir, "service") {
		service = p.services
	} else {
		service = map[string]int{"": 0}
	}

	for service, port := range service {
		if service != "" {
			err := os.MkdirAll(p.name+"/"+key+"/"+service, 0775)
			checkError(err)
		}
		err := filepath.Walk(basePath+key+dir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				name := strings.Split(path, "/")
				// get index of key
				index := func() int {
					for k, v := range name {
						if v == key {
							return k
						}
					}
					return 0
				}()

				fileName := strings.Join(name[index:], "/")
				//fmt.Println(p.name+"/"+fileName)

				err = p.createFilesAndDirectories(fileName, path, service, strconv.Itoa(port))
				if err != nil {
					return err
				}

				return nil
			})
		checkError(err)
	}
}

func (p *project) createFilesAndDirectories(fileName, path, service, port string) error {
	if strings.Contains(fileName, ".go") ||
		strings.Contains(fileName, "Dockerfile") ||
		strings.Contains(fileName, ".toml") ||
		strings.Contains(fileName, ".md") ||
		strings.Contains(fileName, ".mod") ||
		strings.Contains(fileName, "VERSION") {
		source, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		replacer := strings.NewReplacer("project", p.name,
			"servicename", service,
			"portnumber", port)

		destination := replacer.Replace(string(source))
		destPath := p.name + "/" + strings.Replace(fileName, "service", service, -1)
		err = ioutil.WriteFile(destPath, []byte(destination), 0644)
		if err != nil {
			return err
		}
		fmt.Println(destPath, "created.")
	} else {
		err := os.MkdirAll(p.name+"/"+strings.Replace(fileName, "service", service, -1), 0775)
		if err != nil {
			return err
		}
	}
	return nil
}
