package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the frontend setup tool!")
	fmt.Println("Name your project:")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	fmt.Println("Which framework do you want to install?")
	fmt.Println("1. React")
	fmt.Println("2. Next.js")
	fmt.Println("3. Svelte")
	fmt.Println("4. Astro")
	fmt.Println("5. Angular")
	
	var cmd *exec.Cmd

	frameworkInput, _ := reader.ReadString('\n')
	frameworkInput = strings.TrimSpace(frameworkInput)
	var flag bool
	flag = false
	fmt.Println("Want to Install MERN/MEAN Stack? (y/n)")
	stackk, _ := reader.ReadString('\n')
	stackk = strings.TrimSpace(stackk)
	if stackk == "y" {
		flag = true
	}

	switch frameworkInput {
	case "1":
		cmd = exec.Command("npx", "create-react-app@latest", projectName)
	case "2":
		cmd = exec.Command("npx", "create-next-app@latest", projectName)
	case "3":
		cmd = exec.Command("npm", "create", "svelte@latest", projectName)
	case "4":
		cmd = exec.Command("npm", "create", "astro@latest")
	case "5":
		cmd = exec.Command("npm", "install", "-g", "@angular/cli")
		cmd.Run()
		cmd = exec.Command("ng", "new", projectName)
		cmd.Dir = projectName
		cmd.Run()

	default:
		log.Fatal("Invalid option")
	}


	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	if flag {
		cmd = exec.Command("npm", "install", "-g", "express-generator", "nodemon", "mongodb", "mongoose")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Do you want to use Tailwind CSS? (yes/no)")
	tailwind, _ := reader.ReadString('\n')
	tailwind = strings.TrimSpace(tailwind)

	if tailwind == "yes" {
		cmd = exec.Command("npm", "install", "-D", "tailwindcss@latest", "postcss@latest", "autoprefixer@latest")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		cmd = exec.Command("npx", "tailwindcss", "init")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Setup complete!")
}
