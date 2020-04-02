package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	p.createServices()

	fmt.Println("micgo completed.")
}

func (p *project) createServices() {
	// create /internal/services*
	for service, _ := range p.services {
		// create directories
		servicePath := p.name + internal + "/" + service
		fullServicePath := servicePath + "/config"
		err := os.MkdirAll(fullServicePath, 0775)
		checkError(err)

		err = filepath.Walk(basePath+internalService,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.Contains(path, ".go") {
					name := strings.Split(path, "/")
					// get index of "service"
					index := func() int {
						for k, v := range name {
							if v == "service" {
								return k
							}
						}
						return 0
					}()

					fileName := strings.Join(name[index+1:], "/")
					// create config.go files
					sfile, err := ioutil.ReadFile(path)
					checkError(err)

					replacer := strings.NewReplacer("project", p.name, "servicename", service)
					dfile := replacer.Replace(string(sfile))

					err = ioutil.WriteFile(servicePath+"/"+fileName, []byte(dfile), 0644)
					checkError(err)
				}
				return nil
			})
		checkError(err)

		fmt.Println(service + " created.")
	}
}
