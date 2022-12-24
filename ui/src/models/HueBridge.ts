class HueBridge {
  name: string;

  constructor(data: any) {
    this.name = data.name;
  }

  id() {
    return `bridge-${this.name}`;
  }
}

export default HueBridge;
