import HueBridge from "./HueBridge";

class TemperatureSensor {
  id: string;
  name: string;
  bridge: HueBridge;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
  }
}

export default TemperatureSensor;
