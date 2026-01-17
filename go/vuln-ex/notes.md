# govulncheck

`govulncheck` is a command-line tool for Go that helps find known vulnerabilities in your Go dependencies.

**Key Features:**
- **Vulnerability Detection:** Scans your Go modules for known vulnerabilities listed in the Go vulnerability database.
- **Call Stack Information:** Shows you the call stack from your code to the vulnerable function, helping you understand if your code is actually affected.
- **Low Noise:** Focuses on vulnerabilities that are actually reachable by your code, reducing false positives.

**Usage:**
To run `govulncheck` on your module:
```bash
govulncheck ./...
```
This command will analyze your current module and its dependencies.

For more details, refer to the official Go documentation.