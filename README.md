# TRC20 Block Speed Monitor

This repository contains a Go-based tool designed to measure the speed at which new blocks are created on the TRON blockchain. It specifically targets TRC20 token transactions to assess the network's block generation performance.

## Features

- **Real-time Monitoring:** Continuously tracks the time interval between new blocks.
- **Simple and Efficient:** Minimalistic design for quick deployment and accurate measurement.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/devflex-pro/trc20-speed.git
   ```
2. Navigate to the project directory:
   ```bash
   cd trc20-speed
   ```
3. Install the required Go modules:
   ```bash
   go mod download
   ```

## Usage

To start measuring block speed, run the following command:

```bash
go run main.go
```

The tool will output the time intervals between consecutive blocks.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
