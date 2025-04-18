# Barz Smart Contract Audit Tool 🔍

![Active Development](https://img.shields.io/badge/status-active%20development-yellowgreen) 
[![Go Report Card](https://goreportcard.com/badge/github.com/HarryBarz/barz-sc-audit-tool)](https://goreportcard.com/report/github.com/HarryBarz/barz-sc-audit-tool)

> **Project Status Notice** 📢  
> This project is currently under active development. Core features are being implemented and APIs may change without notice.  
> **Contributions and issue reports are welcome!**  

## Development Status

🚧 **Work in Progress** 🚧  
I'm actively working on:
- Implementing core static analysis features
- Expanding vulnerability detection patterns
- Improving AST parsing capabilities
- Enhancing reporting formats

**Current Stability Level**: Alpha (Not production-ready)  

## Contributing

🤝 **We welcome contributions!**   
⚠️ **Note**: This project is in evolution phase - expect frequent changes to code structure and interfaces.


### Install Dependencies
```bash
# Install Solidity compiler
sudo apt-get install solc

# Install Go dependencies
go mod tidy
```

### Build & Run
```bash
# Build the tool
go build -o barz-sc-audit cmd/main.go

# Analyze contracts (text report)
./barz-sc-audit --path testdata/contracts --output text

# Analyze contracts (JSON report)
./barz-sc-audit --path testdata/contracts --output json
```
