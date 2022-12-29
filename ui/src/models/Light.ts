import HueBridge from './HueBridge';

class Light {
  uniqueID: string;
  id: string;
  name: string;
  bridge: HueBridge;

  constructor(data: any) {
    this.uniqueID = data.uniqueID;
    this.id = data.id;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
  }
}

export default Light;
