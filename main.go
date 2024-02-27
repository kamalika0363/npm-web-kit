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

	fmt.Println("Which framework do you want to install?")
	fmt.Println("1. React")
	fmt.Println("2. Next.js")
	fmt.Println("3. Svelte")
	fmt.Println("4. Astro")

	framework, _ := reader.ReadString('\n')
	framework = strings.TrimSpace(framework)

	var cmd *exec.Cmd

	switch framework {
	case "1":
		cmd = exec.Command("npx", "create-react-app", "my-app")
	case "2":
		cmd = exec.Command("npx", "create-next-app", "my-app")
	case "3":
		cmd = exec.Command("npx", "degit", "sveltejs/template", "my-app")
	case "4":
		cmd = exec.Command("npm", "init", "astro")
	default:
		log.Fatal("Invalid option")
	}

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
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
