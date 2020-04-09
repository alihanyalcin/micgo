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

	p.walk()

	fmt.Print("Completed. ")
}

func (p *project) walk() {
	var path = os.Getenv("GOPATH") + "/src/github.com/alihanyalcin/micgo/base/"

	name := strings.Split(path, "/")
	// get index of "base" key word
	index := len(name) - 2

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			d := strings.Split(path, "/")
			destination := strings.Join(d[index+1:], "/")
			if info.IsDir() { // directories
				if strings.Contains(path, "service") {
					for service, _ := range p.services {
						destinationPath := strings.Replace(destination, "service", service, -1)
						err := p.createProjectDirectories(destinationPath)
						if err != nil {
							return err
						}
					}
				} else {
					err := p.createProjectDirectories(destination)
					if err != nil {
						return err
					}
				}
			} else { // files
				if strings.Contains(path, "service") {
					for service, port := range p.services {
						destinationPath := strings.Replace(destination, "service", service, -1)
						err := p.createProjectFiles(path, destinationPath, service, strconv.Itoa(port))
						if err != nil {
							return err
						}
					}
				} else {
					err := p.createProjectFiles(path, destination, "", "")
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
	checkError(err)

}

func (p *project) createProjectFiles(sourcePath, destinationPath, service, port string) error {
	s, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	var replacer *strings.Replacer
	var perm os.FileMode = 0664
	if strings.Contains(sourcePath, "Makefile") {
		replacer = p.createMakefile()
	} else if strings.Contains(sourcePath, ".sh") {
		replacer = p.createScript()
		perm = 0775
	} else if strings.Contains(sourcePath, ".yml") {
		replacer = p.createDockerCompose()
	} else {
		replacer = strings.NewReplacer("{project}", p.name,
			"{servicename}", service,
			"{portnumber}", port)
	}

	d := replacer.Replace(string(s))
	path := p.name + "/" + destinationPath
	err = ioutil.WriteFile(path, []byte(d), perm)
	if err != nil {
		return err
	}
	fmt.Println(path, "file created.")
	return nil
}

func (p *project) createProjectDirectories(destinationPath string) error {
	path := p.name + "/" + destinationPath
	err := os.MkdirAll(path, 0775)
	if err != nil {
		return err
	}
	fmt.Println(path, "directory created.")
	return nil
}

func (p *project) createMakefile() *strings.Replacer {
	var docker = "docker_servicename"
	var dockers string

	var microservice = "cmd/servicename/servicename"
	var microservices string

	var build = "cmd/servicename/servicename:\n\t$(GO) build $(GOFLAGS) -o $@ ./cmd/servicename"
	var builds string

	var dockerBuild = "docker_servicename:\n" +
		"\tdocker build \\\n" +
		"\t\t-f cmd/servicename/Dockerfile \\\n" +
		"\t\t-t project/servicename:$(VERSION) \\\n" +
		"\t\t."
	var dockerBuilds string

	for service, _ := range p.services {
		dockers += strings.Replace(docker, "servicename", service, -1) + " "
		microservices += strings.Replace(microservice, "servicename", service, -1) + " "
		builds += "\n" + strings.Replace(build, "servicename", service, -1) + "\n"
		dockerBuilds += "\n" + strings.NewReplacer("servicename", service, "project", p.name).Replace(dockerBuild) + "\n"
	}
	return strings.NewReplacer("{project}", p.name,
		"{microservices}", microservices,
		"{builds}", builds,
		"{dockers}", dockers,
		"{dockerbuilds}", dockerBuilds)
}

func (p *project) createScript() *strings.Replacer {
	var microservice = "# servicename Microservice\n" +
		"cd $CMD/servicename\n" +
		"exec -a project-servicename ./servicename &\n" +
		"cd $DIR\n"
	var services string

	for service, _ := range p.services {
		services += strings.NewReplacer("servicename", service, "project", p.name).Replace(microservice) + "\n"
	}

	return strings.NewReplacer("{project}", p.name,
		"{services}", services)
}

func (p *project) createDockerCompose() *strings.Replacer {
	var dockerservice = "  {servicename}:\n" +
		"    image: {project}/{servicename}:1.0\n" +
		"    container_name: {project}-{servicename}\n" +
		"    restart: always\n" +
		"    ports:\n" +
		"     - {portnumber}:{portnumber}\n"
	var dockerservices string

	for service, port := range p.services {
		dockerservices += strings.NewReplacer("{servicename}", service, "{project}", p.name, "{portnumber}", strconv.Itoa(port)).Replace(dockerservice)
	}
	return strings.NewReplacer("{project}", p.name,
		"{dockerservices}", dockerservices)
}
