import TemperatureReading from './TemperatureReading';
import TemperatureSensorExtended from './TemperatureSensorExtended';
import Group from './Group';
import GroupExtended from './GroupExtended';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';
import type TemperatureReadingsResult from '../types/TemperatureReadingsResult';
import type TemperatureSensorResult from '../types/TemperatureSensorResult';
import type GroupsResult from '../types/GroupsResult';

class GoHueDataApi {
  static apiUrl() {
    const env = import.meta.env;
    const port = env.VITE_BACKEND_PORT || 8080;
    return `http://localhost:${port}/api`;
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

  static async getGroup(id: string, bridgeName: string): Promise<GroupExtended> {
    const queryParams = new URLSearchParams();
    queryParams.append('bridge', bridgeName);
    queryParams.append('id', id);
    const data = await this.get(`/group?${queryParams.toString()}`);
    return new GroupExtended(data);
  }

  static async getLiveTemperatureSensors(fahrenheit?: boolean) {
    let path = '/live/temperature-sensors';
    if (typeof fahrenheit === 'boolean') {
      path += `?fahrenheit=${fahrenheit ? '1' : '0'}`;
    }
    const result = await this.get(path);
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
    if (typeof filter?.fahrenheit === 'boolean') {
      params.append('fahrenheit', filter.fahrenheit ? '1' : '0');
    }
    if (typeof filter?.updatedSince === 'string' && filter.updatedSince.length > 0) {
      params.append('updated_since', filter.updatedSince);
    }
    if (typeof filter?.updatedBefore === 'string' && filter.updatedBefore.length > 0) {
      params.append('updated_before', filter.updatedBefore);
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
