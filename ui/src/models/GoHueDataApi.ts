import TemperatureReading from './TemperatureReading';
import TemperatureSensorExtended from './TemperatureSensorExtended';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';
import type TemperatureReadingsResult from '../types/TemperatureReadingsResult';

class GoHueDataApi {
  static apiUrl() {
    return `http://localhost:${process.env.REACT_APP_BACKEND_PORT}/api`;
  }

  static async getTemperatureSensor(id: string) {
    const data = await this.get(`/temperature-sensor?id=${encodeURIComponent(id)}`);
    return new TemperatureSensorExtended(data);
  }

  static async getTemperatureSensors() {
    const result = await this.get('/temperature-sensors/live');
    const tempSensors: TemperatureSensorExtended[] = result.temperatureSensors.map(
      (data: any) => new TemperatureSensorExtended(data)
    );
    return tempSensors;
  }

  static async getTemperatureReadings(filter?: TemperatureReadingFilter): Promise<TemperatureReadingsResult> {
    const params = new URLSearchParams();
    if (typeof filter?.page === 'number') {
      params.append('page', filter.page.toString());
    }
    if (typeof filter?.perPage === 'number') {
      params.append('per_page', filter.perPage.toString());
    }
    if (typeof filter?.perPage === 'number') {
      params.append('per_page', filter.perPage.toString());
    }
    if (typeof filter?.sensorID === 'string' && filter.sensorID.length > 0) {
      params.append('sensor_id', filter.sensorID);
    }
    if (typeof filter?.bridge === 'string' && filter.bridge.length > 0) {
      params.append('bridge', filter.bridge);
    }
    const queryString = params.toString();
    let path = '/temperature-readings';
    if (queryString.length > 0) {
      path += `?${queryString}`;
    }
    const result = await this.get(path);
    const { temperatureReadings, ...rest } = result;
    const tempReadings: TemperatureReading[] = temperatureReadings.map(
      (data: any) => new TemperatureReading(data)
    );
    return { temperatureReadings: tempReadings, ...rest };
  }

  static async get(path: string) {
    const response = await fetch(`${this.apiUrl()}${path}`);
    if (response.status >= 200 && response.status < 300) {
      return await response.json();
    }
    throw new Error(response.statusText);
  }
}

export default GoHueDataApi;
