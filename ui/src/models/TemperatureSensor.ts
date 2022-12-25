import HueBridge from "./HueBridge";

class TemperatureSensor {
  id: string;
  name: string;
  bridge: HueBridge;
  lastUpdated: string;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
    this.lastUpdated = data.lastUpdated;
  }

  lastUpdatedAt() {
    if (!this.lastUpdated || this.lastUpdated.length < 1) return null;
    return new Date(this.lastUpdated);
  }
}

export default TemperatureSensor;
