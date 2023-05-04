const os = require('node:os');
const path = require('node:path');
const childProcess = require('node:child_process');

function getBinary() {
    const availableBinaries = {
        darwin: {
            arm64: 'ae_prettier_action_darwin_arm64',
            x64: 'ae_prettier_action_darwin_amd64',
        },
        linux: {
            arm64: 'ae_prettier_action_linux_arm64',
            x64: 'ae_prettier_action_linux_amd64',
        },
        win32: {
            arm64: 'ae_prettier_action_windows_arm64.exe',
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
const distPath = path.resolve(`${__dirname}${path.sep}dist`);

if (process.env.AE_ACTION_DEBUG === true) {
    console.log(`Running go_action at ${distPath} with ${binary}`);
    console.log();
}

const spawnRtrn = childProcess.spawnSync(`${distPath}${path.sep}${binary}`, {stdio: 'inherit'});