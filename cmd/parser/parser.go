package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ContractAST struct {
	FilePath string
	AST      map[string]interface{}
}

func ParseSolidityFiles(root string) ([]ContractAST, error) {
	var contracts []ContractAST

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sol" {
			ast, err := GetSolidityAST(path)
			if err != nil {
				return err
			}
			contracts = append(contracts, ContractAST{
				FilePath: path,
				AST:      ast.AST,
			})
		}
		return nil
	})

	return contracts, err
}

func GetSolidityAST(filePath string) (ContractAST, error) {
	cmd := exec.Command("solc", "--ast-json", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ContractAST{}, fmt.Errorf("solc error: %v\n%s", err, output)
	}

	var astData map[string]interface{}
	if err := json.Unmarshal(output, &astData); err != nil {
		return ContractAST{}, err
	}

	return ContractAST{
		FilePath: filePath,
		AST:      astData,
	}, nil
}
