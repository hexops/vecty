// +build js

package vecty

import "syscall/js"

// os/exec does not work in GOOS=js GOARCH=wasm, so we implement the minor functionality that we
// need for the test suite using the NodeJS API on our own.

func commandOutput(command string, args ...string) (string, error) {
	argsi := make([]interface{}, len(args))
	for i, arg := range args {
		argsi[i] = arg
	}
	proc := childProcess.Call("spawnSync", command, argsi)
	return proc.Get("stdout").String(), nil
}

var childProcess = js.Global().Call("require", "child_process")
