package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":4321", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.Contains(path, "favicon.ico") {
		return
	}
	if runtime.GOOS == "windows" {
		go updateGitFolder(path[1:len(path)])
	} else {
		go updateGitFolder(path)
	}
	fmt.Fprint(w, "ok")
}

func updateGitFolder(path string) {
	fmt.Println()
	fmt.Println("Path: " + path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Not exists! Ignored")
		return
	}
	if err := os.Chdir(path); err != nil {
		fmt.Println("Access denied")
		return
	}
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("Not a GIT folder! Ignored")
		return
	}
	fmt.Println("Revert all changes...")
	runCommand(exec.Command("git", "reset", "--hard", "HEAD"))
	fmt.Println("Pulling...")
	runCommand(exec.Command("git", "pull"))

	if _, err := os.Stat("deploy.sh"); err == nil {
		fmt.Println("Runing external deploy.sh script...")
		runCommand(exec.Command("./deploy.sh"))
	}

	if _, err := os.Stat("build.sh"); err == nil {
		fmt.Println("Runing external build.sh script...")
		runCommand(exec.Command("./build.sh"))
	}

	fmt.Println("Done")
}

func runCommand(cmd *exec.Cmd) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		fmt.Println(out.String())
	} else {
		fmt.Println(err)
	}
}
