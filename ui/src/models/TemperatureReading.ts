import TemperatureSensor from "./TemperatureSensor";

class TemperatureReading {
  timestamp: string;
  temperature: number;
  units: string;
  sensor: TemperatureSensor;
  id: string;

  constructor(data: any) {
    this.id = data.id;
    this.timestamp = data.timestamp;
    this.temperature = data.temperature;
    this.units = data.units;
    this.sensor = new TemperatureSensor(data.temperatureSensor);
  }

  timestampAsDate() {
    if (!this.timestamp || this.timestamp.length < 1) return null;
    return new Date(`${this.timestamp}Z`); // parse as UTC
  }
}

export default TemperatureReading;
