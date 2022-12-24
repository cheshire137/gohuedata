import HueBridge from "./HueBridge";

class TemperatureSensor {
  name: string;
  bridge: HueBridge;

  constructor(data: any) {
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
  }

  id() {
    return `sensor-${this.name}-${this.bridge.id()}`;
  }
}

export default TemperatureSensor;
