package analyzer

import (
	"github.com/HarryBarz/barz-sc-audit-tool/parser"
)

type Vulnerability struct {
	Contract    string `json:"contract"`
	Line        int    `json:"line"`
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

func AnalyzeContracts(contracts []parser.ContractAST) []Vulnerability {
	var vulnerabilities []Vulnerability

	for _, contract := range contracts {
		vulns := AnalyzeAST(contract)
		vulnerabilities = append(vulnerabilities, vulns...)
	}

	return vulnerabilities
}

func AnalyzeAST(contract parser.ContractAST) []Vulnerability {
	var vulns []Vulnerability

	// Example detection logic (simplified)
	if detectReentrancy(contract.AST) {
		vulns = append(vulns, Vulnerability{
			Contract:    contract.FilePath,
			Line:        42, // Should get from AST
			Type:        "Reentrancy",
			Severity:    "High",
			Description: "Potential reentrancy vulnerability",
		})
	}

	return vulns
}

func detectReentrancy(ast map[string]interface{}) bool {
	// Implement actual AST traversal here
	return true // Simplified for example
}
