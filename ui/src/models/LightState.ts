import Light from './Light';

class LightState {
  timestamp: string;
  on: boolean;
  light: Light;

  constructor(data: any) {
    this.timestamp = data.timestamp;
    this.on = data.on;
    this.light = new Light(data.light);
  }
}

export default LightState;
