import TemperatureSensor from "./TemperatureSensor";
import TemperatureReading from "./TemperatureReading";

class TemperatureSensorExtended extends TemperatureSensor {
  latestReading: TemperatureReading;

  constructor(data: any) {
    super(data);
    this.latestReading = new TemperatureReading(data.latestReading);
  }
}

export default TemperatureSensorExtended;
