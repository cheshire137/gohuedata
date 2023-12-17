import TemperatureSensor from "./TemperatureSensor";

class TemperatureReadingSummary {
  timestamp: string;
  minTemperature: number;
  maxTemperature: number;
  avgTemperature: number;
  units: string;
  sensor: TemperatureSensor;
  id: string;
  _date: Date | null;

  constructor(data: any) {
    this.id = data.id;
    this.timestamp = data.timestamp;
    this.minTemperature = data.minTemperature;
    this.maxTemperature = data.maxTemperature;
    this.avgTemperature = data.avgTemperature;
    this.units = data.units;
    this.sensor = new TemperatureSensor(data.temperatureSensor);
    this._date = null;
  }

  timestampAsDate() {
    if (this._date) return this._date;
    if (!this.timestamp || this.timestamp.length < 1) return null;
    this._date = new Date(`${this.timestamp}Z`); // parse as UTC
    return this._date;
  }
}

export default TemperatureReadingSummary;
