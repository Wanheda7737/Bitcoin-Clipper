# 🚀 Go-Bitcoin-Clipper (Educational Research Tool)

![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)
![License](https://img.shields.io/badge/license-Apache%202.0-green.svg)

## ⚠️ Disclaimer
**IMPORTANT:** This project is created strictly for **educational purposes**, **cybersecurity research**, and **malware analysis training**. The author is NOT responsible for any misuse or damage caused by this software. Unauthorized use of this tool against systems without prior consent is illegal.

---

## 📌 Overview
`Go-Bitcoin-Clipper` is a high-performance, lightweight implementation of a cryptocurrency address swapper, ported from Python to **Go**. It is designed to demonstrate how clipboard monitoring and regex-based pattern matching work in a low-latency environment.

### Key Features
- **High Performance:** Uses Go's native concurrency and efficient string handling (`strings.Builder`).
- **Smart Replacement:** 
    - Detects single Bitcoin addresses and replaces them instantly.
    - Detects Bitcoin addresses embedded within sentences and swaps only the address, preserving the surrounding context.
- **Low Footprint:** Minimal CPU and RAM usage due to Go's compiled nature.

## 🛠 Technical Stack
- **Language:** [Go (Golang)](https://golang.org/)
- **Clipboard Library:** `golang.design/x/clipboard`
- **Pattern Matching:** RE2 Regex Engine

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher installed on your system.
- C compiler (for `x/clipboard` dependen on Windows/Linux).

### Installation & Running
1. **Clone the repository:**
```bash
   git clone https://github.com/yourusername/go-bitcoin-clipper.git
   cd go-bitcoin-clipper
```
2. **Initialize modules:**
```bash
    go mod tidy
```
3. **Run the tool:**
```bash
   go run main.go
```
*🔍 How It Works*

1.*Monitoring:*The tool enters an infinite loop with a 1-second sleep interval (configurable).
2.*Detection:*It reads the system clipboard and applies a complex Regex pattern to identify Bitcoin (Legacy, SegWit, and Bech32) addresses.
3.*Logic Branching:*If the clipboard contains only a valid address
```bash → ``` Replace it.
If the clipboard contains a sentence with an address
```bash → ``` Perform an in-place replacement of the address while keeping the sentence intact.

*📜 License*

This project is licensed under the Apache License 2.0. See the LICENSE file for details.


