import TemperatureSensor from "./TemperatureSensor";

class TemperatureReading {
  lastUpdated: string;
  temperature: number;
  units: string;
  sensor: TemperatureSensor;
  id: string;

  constructor(data: any) {
    this.id = data.id;
    this.lastUpdated = data.lastUpdated;
    this.temperature = data.temperature;
    this.units = data.units;
    this.sensor = new TemperatureSensor(data.temperatureSensor);
  }
}

export default TemperatureReading;
