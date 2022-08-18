/*
* @description Benchmarking tool for git hooks. It will run the hook and measure
* the time it takes to run.
*
* @author Christian Gama e Silva
*/
module.exports = class {
  constructor() {
    this.start = Date.now();
  }

  end() {
    const end = Date.now();
    const duration = end - this.start;
    return duration;
  }
};
