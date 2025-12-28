const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path')

function createWindow() {
  const win = new BrowserWindow({
    icon: path.join(__dirname, "icon.ico"),
    width: 800,
    height: 600,
    transparent: true,
    frame: false,

    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
      enableRemoteModule: true
    }
  });

  win.loadFile('src/ui/index.html');
  win.webContents.openDevTools();
  ipcMain.on('close-window-h12uhd', () => {
    const window = BrowserWindow.getFocusedWindow();
    if (window) {
        window.close();
    }
  });
  ipcMain.on('minimize-window-11jsh2', () => {
    const window = BrowserWindow.getFocusedWindow();
    if (window) {
        window.minimize();
    }
  });
}

app.whenReady().then(createWindow);