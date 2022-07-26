const e = "\x1b[31m"; // error
const s = "\x1b[32m"; // success
const d = "\x1b[0m"; // default

module.exports = {
  error: (hook, msg) => {
    console.log(`${e}[${hook.toLowerCase()}] - ${msg}${d}`);
    process.exit(1);
  },

  success: (hook, msg) => {
    console.log(`${s}[${hook.toLowerCase()}] - ${msg}${d}`);
  },

  info: (hook, msg) => {
    console.log(`${d}[${hook.toLowerCase()}] - ${msg}${d}`);
  },
};
