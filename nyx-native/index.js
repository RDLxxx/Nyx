const path = require('path');
const binaryPath = path.join(__dirname, 'build', 'Release', 'nyx-native.node');
module.exports = require(binaryPath)