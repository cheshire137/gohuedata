import TemperatureReading from './TemperatureReading';
import TemperatureSensorExtended from './TemperatureSensorExtended';
import Group from './Group';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';
import type TemperatureReadingsResult from '../types/TemperatureReadingsResult';
import type TemperatureSensorResult from '../types/TemperatureSensorResult';
import type GroupsResult from '../types/GroupsResult';

class GoHueDataApi {
  static apiUrl() {
    return `http://localhost:${process.env.REACT_APP_BACKEND_PORT}/api`;
  }

  static async getTemperatureSensor(id: string, fahrenheit: boolean): Promise<TemperatureSensorResult> {
    const queryParams = new URLSearchParams();
    queryParams.append('fahrenheit', fahrenheit ? '1' : '0');
    queryParams.append('id', id);
    const data = await this.get(`/temperature-sensor?${queryParams.toString()}`);
    const { temperatureSensor: tempSensorData, ...rest } = data;
    const temperatureSensor = new TemperatureSensorExtended(tempSensorData);
    return { temperatureSensor, ...rest };
  }

  static async getLiveTemperatureSensors() {
    const result = await this.get('/live/temperature-sensors');
    const tempSensors: TemperatureSensorExtended[] = result.temperatureSensors.map(
      (data: any) => new TemperatureSensorExtended(data)
    );
    return tempSensors;
  }

  static async getLiveGroups(): Promise<GroupsResult> {
    const result = await this.get('/live/groups');
    const { groups: groupData, ...rest } = result;
    const groups: Group[] = groupData.map((data: any) => new Group(data));
    return { groups, ...rest };
  }

  static async getTemperatureReadings(filter?: TemperatureReadingFilter): Promise<TemperatureReadingsResult> {
    const params = new URLSearchParams();
    if (typeof filter?.page === 'number') {
      params.append('page', filter.page.toString());
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
    const json = await response.json();
    if (response.status >= 200 && response.status < 300) {
      return json;
    }
    let errorMessage = response.statusText
    if (json && json.error) {
      errorMessage += `: ${json.error}`;
    }
    throw new Error(errorMessage);
  }
}

export default GoHueDataApi;
