const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const platform = process.platform;
const arch = process.arch;

let binary = '';

if (platform === 'darwin') {
    binary = arch === 'arm64' ? 'ziplink-darwin-arm64' : 'ziplink-darwin-amd64';
} else if (platform === 'linux') {
    binary = 'ziplink-linux-amd64';
} else if (platform === 'win32') {
    binary = 'ziplink-windows-amd64.exe';
}

if (binary) {
    const binaryPath = path.join(__dirname, 'bin', binary);
    if (fs.existsSync(binaryPath) && platform !== 'win32') {
        try {
            fs.chmodSync(binaryPath, 0o755);
            console.log(`Permissions set for ${binary}`);
        } catch (err) {
            console.error(`Failed to set permissions: ${err.message}`);
        }
    }
} else {
    console.warn('Unsupported platform for ziplink-cli pre-built binaries.');
}
