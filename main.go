package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <dominio>")
		os.Exit(1)
	}

	host := os.Args[1]

	fmt.Printf("Analizando %s...\n", host)

	result, err := analyzeHost(host)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	printResults(result)
}

func printResults(resp *HostResponse) {
	fmt.Printf("\nResultados para %s\n", resp.Host)

	for _, ep := range resp.Endpoints {
		fmt.Println("----------------------------")
		fmt.Println("IP:", ep.IPAddress)
		fmt.Println("Estado:", ep.StatusMessage)
		fmt.Println("Grade:", ep.Grade)

		fmt.Println("Protocolos soportados:")
		for _, p := range ep.Details.Protocols {
			fmt.Printf("  - %s %s\n", p.Name, p.Version)
		}

		fmt.Println("Vulnerabilidades:")
		fmt.Println("  Heartbleed:", ep.Details.Heartbleed)
		fmt.Println("  Poodle:", ep.Details.Poodle)
		fmt.Println("  Freak:", ep.Details.Freak)
		fmt.Println("  Logjam:", ep.Details.Logjam)
	}
}
