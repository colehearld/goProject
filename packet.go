package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/docker/docker/client"
)

func main() {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		os.Exit(1)
	}

	// Create a new context
	ctx := context.Background()

	// Prompt the user to enter the name of a container
	var containerName string
	fmt.Print("Enter the name of a Docker container: ")
	fmt.Scanln(&containerName)

	// Get information about the container
	containerInfo, err := cli.ContainerInspect(ctx, containerName)
	if err != nil {
		fmt.Println("Failed to get information about container:", err)
		os.Exit(1)
	}

	// Store the container ID in a slice
	packet := make([]string, 0) 
	packet = append(packet, containerInfo.ID)

	// Print the container ID
	fmt.Println("Stored container ID:", containerInfo.ID)

	writePacketToTxt(containerInfo.ID)

}

func displayPacket(packet []string) { fmt.Println(packet) }

func getRunningPackets() {}

func writePacketToTxt(containerInfo string){
	f, err := os.Create("container_ids.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(containerInfo)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
}
