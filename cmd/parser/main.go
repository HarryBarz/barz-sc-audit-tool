package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HarryBarz/barz-sc-audit-tool/analyzer"
	"github.com/HarryBarz/barz-sc-audit-tool/parser"
	"github.com/HarryBarz/barz-sc-audit-tool/reporter"
)

func main() {
	path := flag.String("path", "testdata/contracts", "Path to analyze")
	outputFormat := flag.String("output", "text", "Report format (text/json)")
	flag.Parse()

	// Parse Solidity files
	contracts, err := parser.ParseSolidityFiles(*path)
	if err != nil {
		fmt.Printf("Error parsing contracts: %v\n", err)
		os.Exit(1)
	}

	// Analyze contracts
	results := analyzer.AnalyzeContracts(contracts)

	// Generate report
	switch *outputFormat {
	case "text":
		reporter.GenerateTextReport(results)
	case "json":
		reporter.GenerateJSONReport(results)
	default:
		fmt.Println("Unsupported output format")
		os.Exit(1)
	}
}
