#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");

const platform = process.platform;
const arch = process.arch;

let binary = "";

if (platform === "darwin") {
  binary = arch === "arm64" ? "ziplink-darwin-arm64" : "ziplink-darwin-amd64";
} else if (platform === "linux") {
  binary = "ziplink-linux-amd64";
} else if (platform === "win32") {
  binary = "ziplink-windows-amd64.exe";
}

if (!binary) {
  console.error("Error: Unsupported platform.");
  process.exit(1);
}

const binaryPath = path.join(__dirname, "bin", binary);
const args = process.argv.slice(2);

const child = spawn(binaryPath, args, { stdio: "inherit" });

child.on("error", (err) => {
  console.error(`Failed to start binary: ${err.message}`);
  process.exit(1);
});

child.on("exit", (code) => {
  process.exit(code);
});
