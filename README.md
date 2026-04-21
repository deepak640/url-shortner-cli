# 🔗 ZipLink CLI (ziplink-cli)

[![npm version](https://img.shields.io/npm/v/ziplink-cli.svg)](https://www.npmjs.com/package/ziplink-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fast, lightweight, and cross-platform command-line tool to shorten URLs instantly. Built with **Go** for high performance and distributed via **npm** for maximum convenience.

---

## 🚀 Quick Install

Install it globally using npm:

```bash
npm install -g ziplink-cli
```

---

## 🛠️ Usage

Once installed, you can use the `ziplink` command from anywhere in your terminal.

### 1. Shorten a URL
Generate a tiny link for any `https://` address.
```bash
ziplink shorten --url https://www.google.com
```

### 2. List your URLs
See every URL you've shortened in a beautiful table, including timestamps.
```bash
ziplink list
```

### 3. Remove a URL
Delete a shortened link using its unique short code.
```bash
ziplink remove R4anSi
```

---

## 🛠️ Command Options

### `shorten`
- `--url`: The URL to shorten (Required)
- `--custom`: Custom short code (Optional)
- `--expiry`: Expiry in hours (Optional, e.g., 24 for 1 day)

Example with custom code and expiry:
```bash
ziplink shorten --url https://www.google.com --custom google --expiry 24h
```

---

## ✨ Features

- **Blazing Fast**: Compiled Go binaries for every platform (Mac ARM/Intel, Linux, Windows).
- **Beautiful UI**: Colorful terminal output with clean data tables.
- **Zero Config**: Connects to a pre-configured Vercel backend by default.
- **Privacy Conscious**: Uses your hardware's unique MachineID to keep your links private to your device.
- **Customizable**: Pro users can override the backend API using environment variables.

---

## 💻 Local Development (For Developers)

If you want to run the source code directly or use a custom backend:

### Prerequisites
- [Go](https://golang.org/dl/) (version 1.20+)
- A backend server running the [URL Shortener API](https://url-shortner-rosy-omega.vercel.app/)

### Setup
1. Clone the repo:
   ```bash
   git clone https://github.com/deepak640/url-shortner-cli.git
   cd url-shortner-cli
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. (Optional) Custom Backend:
   Create a `.env` file in the root:
   ```env
   SERVER=http://localhost:3000/
   ```
4. Run:
   ```bash
   go run main.go list
   ```

---


## 🏗️ Architecture
This project uses a "dual-stack" approach:
- **Core Logic**: Written in Go for raw speed and easy cross-compilation.
- **Distribution**: Wrapped in a thin Node.js layer to allow simple installation via `npm install -g`.

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---

**Developed with ❤️ by [Deepak Negi](https://github.com/deepak640)**
