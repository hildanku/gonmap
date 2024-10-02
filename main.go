package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// the format data value is <subdomain>,<ip>
	data := `subdomain1,ip1`

	resultFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Failed to create result file:", err)
		return
	}
	defer resultFile.Close()

	writer := bufio.NewWriter(resultFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Println("Invalid Format:", line)
			continue
		}
		subdomain := parts[0]
		ip := parts[1]

		fmt.Printf("Scanning %s (%s)...\n", subdomain, ip)

		cmd := exec.Command("nmap", ip)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Can't run nmap:", err)
			continue
		}

		writer.WriteString(fmt.Sprintf("Successfull scanning %s (%s):\n", subdomain, ip))
		writer.WriteString(string(output) + "\n")
		writer.WriteString("=========================================================\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error read data:", err)
	}
}
