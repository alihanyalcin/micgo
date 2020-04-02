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
		// TODO: Write help
		fmt.Println("HELP!")
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
			// TODO: Write help
			fmt.Println("HELP!")
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
		services[s[0]] = port
	}
	p := project{
		name:     projectName,
		services: services,
	}
	p.create()
}

var names []string

func checkNameValid(name string) bool {
	validName, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)
	if !validName {
		fmt.Println("Invalid name:", name)
		// TODO: Write help
		fmt.Println("HELP!")
		return false
	}
	if len(names) < 1 {
		names = append(names, name)
	} else {
		for _, v := range names {
			if v == name {
				fmt.Println("Duplicate service names:", name)
				// TODO: Write help
				fmt.Println("HELP!")
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
		fmt.Println("Port number should be numeric:", port)
		// TODO: Write help
		fmt.Println("HELP!")
		return false, 0
	}
	if p <= 1023 || p >= 49151 {
		fmt.Println("Invalid port:", port)
		// TODO: Write help
		fmt.Println("HELP!")
		return false, 0
	}
	if len(ports) < 1 {
		ports = append(ports, p)
	} else {
		for _, v := range ports {
			if v == p {
				fmt.Println("Duplicate service ports:", port)
				// TODO: Write help
				fmt.Println("HELP!")
				return false, 0
			}
		}
		ports = append(ports, p)
	}
	return true, p
}
