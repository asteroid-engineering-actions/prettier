const prettier = require('prettier/cli.js');

if (process.argv[2] === 'ae:action:prettier:cli') {
  prettier.run(process.argv.slice(3));
}
