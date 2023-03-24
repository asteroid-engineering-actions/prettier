const os = require('node:os');
const path = require('node:path');
const childProcess = require('node:child_process');

function getBinary() {
    const availableBinaries = {
        darwin: {
            arm64: 'ae_prettier_action_darwin_arm64',
            amd64: 'ae_prettier_action_darwin_amd64',
        },
        linux: {
            arm64: 'ae_prettier_action_linux_arm64',
            amd64: 'ae_prettier_action_linux_amd64',
            x64: 'ae_prettier_action_linux_amd64',
        },
        win32: {
            arm64: 'ae_prettier_action_windows_arm64.exe',
            amd64: 'ae_prettier_action_windows_amd64.exe',
            x64: 'ae_prettier_action_windows_amd64.exe',
        },
    }

    const arch = os.arch();
    const platform = os.platform();

    if (!availableBinaries[platform]) {
        console.log(`${platform} is currently not supported`);
        os.exit(1);
    }

    if (!availableBinaries[platform][arch]) {
        console.log(`${arch} is currently not supported on ${platform}`);
        os.exit(1);
    }

    return availableBinaries[platform][arch]
}

const binary = getBinary();
const distPath = path.resolve(`.${path.sep}dist`);

const spawnRtrn = childProcess.spawnSync(`${distPath}${path.sep}${binary}`, {stdio: 'inherit'});