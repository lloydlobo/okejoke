package execcmd

import (
	"log"
	"os/exec"
)

// /////////////////////////////////////////////////////
// OS Execute Commands
// /////////////////////////////////////////////////////

// osExecPwd returns slice of byte of jokes.
//
// Convert return to `string `with string(osExecPwd())
func OsExecPwd() string {
	cmdPwd := "pwd"
	out, err := exec.Command(cmdPwd).Output()
	if err != nil {
		log.Fatal(err)
	}
	location := string(out) // Convert []byte to string.
	return location
}
