const util = require("util");
const childProcess = require("child_process");
const exec = util.promisify(childProcess.exec);

const log = require("./log");
const Bench = require("./bench");

module.exports = async function (hook, script) {
  const bench = new Bench();
  log.info(hook, "Running...");

  try {
    const response = await exec(script);
    stdout = response.stdout;
  } catch (error) {
    if (error.code !== 0 && error.code !== 7) {
      log.error(hook, error.stderr);
    }
  }

  log.success(hook, `Done (took ${bench.end()} ms)`);
};
