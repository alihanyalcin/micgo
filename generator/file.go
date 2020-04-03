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

	p.cre("internal", "pkg")
	p.cre("internal", "service")
	p.cre("cmd", "service")
	// Create project directory
	/*err := os.Mkdir(p.name, 0755)
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

	// create /internal/services*
	p.createInternalServices()

	// create /cmd/services*
	p.createCmdServices()

	fmt.Println("micgo completed.")*/
}

func (p *project) createInternalServices() {
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
					if err != nil {
						return err
					}

					replacer := strings.NewReplacer("project", p.name, "servicename", service)
					dfile := replacer.Replace(string(sfile))

					err = ioutil.WriteFile(servicePath+"/"+fileName, []byte(dfile), 0644)
					if err != nil {
						return nil
					}
					fmt.Println(servicePath + "/" + fileName + " created.")
				}
				return nil
			})
		checkError(err)
	}
}

func (p *project) cre(key, dir string) {
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
		err := filepath.Walk(basePath+"/"+key+"/"+dir,
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

func (p *project) createFilesAndDirectories(fileName, path, service string, port string) error {
	if strings.Contains(fileName, ".go") ||
		strings.Contains(fileName, "Dockerfile") ||
		strings.Contains(fileName, ".toml") {
		source, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		replacer := strings.NewReplacer("project", p.name,
			"servicename", service,
			"portnumber", string(port))

		destination := replacer.Replace(string(source))

		err = ioutil.WriteFile(p.name+"/"+strings.Replace(fileName, "service", service, -1), []byte(destination), 0644)
		if err != nil {
			return err
		}
	} else {
		err := os.MkdirAll(p.name+"/"+strings.Replace(fileName, "service", service, -1), 0775)
		if err != nil {
			return err
		}
	}
	return nil
}
