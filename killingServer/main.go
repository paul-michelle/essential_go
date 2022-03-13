package main

import (
	"fmt"
	"killingServer/fileUtils"
	"os"
)

func main() {
	fmt.Println("Entering main func. About to call GetServerPid")

	pid, err := fileUtils.GetServerPid("server.pid")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	if pid == 0 {
		os.Exit(1)
	}

	fmt.Println("Got server's PID. About to kill server")
	fmt.Printf("Killing server with PID=%d\n", pid)
}