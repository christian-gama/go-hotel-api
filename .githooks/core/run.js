/*
* @description Core module for git hooks. It will run the hook and measure the time it took.
* It will also display a log with the current status of the hook.
*
* @author Christian Gama e Silva
*/

const util = require("util");
const childProcess = require("child_process");
const exec = util.promisify(childProcess.exec);

const log = require("./log");
const Benchmark = require("./benchmark");

module.exports = async function (hook, script) {
  const bench = new Benchmark();
  log.info(hook, "Running...");

  try {
    const response = await exec(script);
    stdout = response.stdout;
  } catch (error) {
    if (error.code !== 0 && error.code !== 7) {
      log.error(hook, error.stdout);
    }
  }

  log.success(hook, `Done (took ${bench.end()} ms)`);
};
