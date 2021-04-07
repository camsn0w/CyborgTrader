package repl

import (
	"os"
	"os/exec"
)

func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func repl(){
	for  {
		printRepl()
		clear()


	}
}
