import TemperatureReading from '../models/TemperatureReading';

type TemperatureReadingsResult = {
  temperatureReadings: TemperatureReading[];
  page: number;
  totalPages: number;
  totalCount: number;
}

export default TemperatureReadingsResult;
