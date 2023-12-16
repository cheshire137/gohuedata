import HueBridge from './HueBridge';
import Light from './Light';

class Group {
  id: string;
  uniqueID: string;
  name: string;
  type: string;
  bridge: HueBridge;
  totalLights: number;
  totalSensors: number;
  groupClass: string;
  lights: Light[];

  constructor(data: any) {
    this.id = data.id;
    this.uniqueID = data.uniqueID;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
    this.type = data.type;
    this.totalLights = data.totalLights;
    this.totalSensors = data.totalSensors;
    this.groupClass = data.class;
    this.lights = data.lights.map((lightData: any) => new Light(lightData));
  }
}

export default Group;
