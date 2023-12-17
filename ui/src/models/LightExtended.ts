import Light from './Light';
import LightState from './LightState';

class LightExtended extends Light {
  latestState: LightState;

  constructor(data: any) {
    super(data);
    this.latestState = new LightState(data.latestState);
  }
}

export default LightExtended;
