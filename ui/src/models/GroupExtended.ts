import LightExtended from './LightExtended';
import HueBridge from './HueBridge';

class GroupExtended {
  id: string;
  name: string;
  type: string;
  bridge: HueBridge;
  totalLights: number;
  totalSensors: number;
  groupClass: string;
  lights: LightExtended[];

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.bridge = new HueBridge(data.bridge);
    this.type = data.type;
    this.totalLights = data.totalLights;
    this.totalSensors = data.totalSensors;
    this.groupClass = data.class;
    this.lights = data.lights.map((lightData: any) => new LightExtended(lightData));
  }
}

export default GroupExtended;
