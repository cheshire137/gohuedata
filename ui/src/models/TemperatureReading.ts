import TemperatureSensor from "./TemperatureSensor";

class TemperatureReading {
  lastUpdated: string;
  temperature: number;
  units: string;
  sensor: TemperatureSensor;

  constructor(data: any) {
    this.lastUpdated = data.lastUpdated;
    this.temperature = data.temperature;
    this.units = data.units;
    this.sensor = new TemperatureSensor(data.temperatureSensor);
  }

  id() {
    return `reading-${this.lastUpdated}-${this.sensor.id()}`;
  }
}

export default TemperatureReading;
