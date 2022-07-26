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
