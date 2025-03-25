package reporter

import (
	"encoding/json"
	"fmt"

	"barz-sc-audit-tool/analyzer"
)

func GenerateTextReport(vulns []analyzer.Vulnerability) {
	if len(vulns) == 0 {
		fmt.Println("No vulnerabilities found!")
		return
	}

	fmt.Println("Security Audit Report")
	fmt.Println("=====================")

	for _, v := range vulns {
		fmt.Printf("[%s] %s\n", v.Severity, v.Type)
		fmt.Printf("Contract: %s\n", v.Contract)
		fmt.Printf("Line: %d\n", v.Line)
		fmt.Printf("Description: %s\n\n", v.Description)
	}
}

func GenerateJSONReport(vulns []analyzer.Vulnerability) {
	jsonData, err := json.MarshalIndent(vulns, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON report:", err)
		return
	}
	fmt.Println(string(jsonData))
}
