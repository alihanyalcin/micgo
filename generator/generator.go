package generator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type project struct {
	name     string
	services map[string]int
}

func Run(args []string) {
	length := len(args)
	if length == 0 || args[0] != "generate" || length < 3 {
		helpCall()
		return
	}

	// check args are correct
	// check project name
	projectName := args[1]
	checkProjectName := checkNameValid(projectName)
	if !checkProjectName {
		return
	}
	// check service name and port
	var services = make(map[string]int)
	for _, service := range args[2:] {
		s := strings.Split(service, ":")
		if len(s) != 2 {
			helpCall()
			return
		}
		// check service name
		checkServiceName := checkNameValid(s[0])
		if !checkServiceName {
			return
		}
		// check port number
		checkPortNumber, port := checkPortValid(s[1])
		if !checkPortNumber {
			return
		}
		services[strings.ToLower(s[0])] = port
	}
	p := project{
		name:     strings.ToLower(projectName),
		services: services,
	}
	// Create project files
	p.create()
}

var names []string

func checkNameValid(name string) bool {
	validName, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)
	if !validName {
		fmt.Println("Invalid name:", name)
		helpCall()
		return false
	}
	if len(names) < 1 {
		names = append(names, name)
	} else {
		for _, v := range names {
			if v == name {
				fmt.Println("Duplicate service names:", name)
				helpCall()
				return false
			}
		}
		names = append(names, name)
	}

	return true
}

var ports []int

func checkPortValid(port string) (bool, int) {
	p, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Port number should be numeric")
		helpCall()
		return false, 0
	}
	if p <= 1023 || p >= 49151 {
		fmt.Println("Invalid port:", port)
		helpCall()
		return false, 0
	}
	if len(ports) < 1 {
		ports = append(ports, p)
	} else {
		for _, v := range ports {
			if v == p {
				fmt.Println("Duplicate service ports:", port)
				helpCall()
				return false, 0
			}
		}
		ports = append(ports, p)
	}
	return true, p
}

var help = `
Usage: go run github.com/alihanyalcin/micgo generate <project_name> <service_name1>:<service_port1> <service_name2>:<service_port2> ... <service_nameX>:<service_portX>
Example: go run github.com/alihanyalcin/micgo generate testproject service1:12300 service2:12301
`

func helpCall() {
	fmt.Print(help)
}
