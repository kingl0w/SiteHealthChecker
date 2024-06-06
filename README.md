## HealthChecker

HealthChecker is a simple command-line tool for checking the health status of a service, including its availability and SSL certificate details. I started with a tutorial of how to check a websites health with Go in order to learn more about the language and then added some additional functionality to include the SSL information.

### Usage

```bash
./HealthChecker --domain example.com 
```
Optional
```bash
./HealthChecker --domain example.com --ports "8080,8000,25"
```

#### Flags

- `--domain` (`-d`): The URL to check (required).
- `--ports` (`-p`): The ports to check, separated by commas (default: 80,443). This flag is optional for checking other ports, 80 and 443 are checked automaically

### How It Works

The tool sends HTTP requests to the specified domain and ports to check their availability. It also performs SSL certificate checks for HTTPS-enabled services.

### Installation

To use HealthChecker, make sure you have Go installed on your system. Then, you can clone the repository and build the executable:

```bash
git clone https://github.com/yourusername/HealthChecker.git
cd HealthChecker
go build
```

### Dependencies

- [urfave/cli/v2](https://github.com/urfave/cli/v2): A powerful CLI library for Go.
