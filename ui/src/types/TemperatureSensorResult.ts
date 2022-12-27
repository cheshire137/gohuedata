import TemperatureSensor from '../models/TemperatureSensor';

type TemperatureSensorResult = {
  temperatureSensor: TemperatureSensor;
  maxTemperature?: number;
  minTemperature?: number;
}

export default TemperatureSensorResult;
