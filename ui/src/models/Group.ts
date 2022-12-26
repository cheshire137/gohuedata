import HueBridge from './HueBridge';

class Group {
  id: string;
  name: string;
  type: string;
  bridge: HueBridge;
  totalLights: number;
  totalSensors: number;
  groupClass: string;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
    this.type = data.type;
    this.totalLights = data.totalLights;
    this.totalSensors = data.totalSensors;
    this.groupClass = data.class;
  }
}

export default Group;
