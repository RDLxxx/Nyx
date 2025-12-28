const nn = require('nyx-native');
const { ipcRenderer } = require('electron');

function closeWND() {
    ipcRenderer.send("close-window-h12uhd");
}

function minimizeWND() {
    ipcRenderer.send("minimize-window-11jsh2");
}

function newFile() {
    const result = nn.addFile();

    const heading = document.getElementById('fs');
    heading.textContent = 'File: ' + result
}

function readFile() {
    const content = nn.readFile("nyxconf/filename.txt");

    const heading = document.getElementById('fs');
    heading.textContent = 'File: ' + content
}