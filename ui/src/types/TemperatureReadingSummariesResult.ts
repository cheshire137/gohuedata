import TemperatureReadingSummary from '../models/TemperatureReadingSummary';

type TemperatureReadingSummariesResult = {
  temperatureReadingSummaries: TemperatureReadingSummary[];
  page: number;
  totalPages: number;
  totalCount: number;
}

export default TemperatureReadingSummariesResult;
