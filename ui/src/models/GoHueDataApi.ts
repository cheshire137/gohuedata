import TemperatureReading from "./TemperatureReading";
import TemperatureSensorExtended from "./TemperatureSensorExtended";

class GoHueDataApi {
  static apiUrl() {
    return `http://localhost:${process.env.REACT_APP_BACKEND_PORT}/api`;
  }

  static async getTemperatureSensors() {
    const result = await this.get('/temperature-sensors/live');
    const tempSensors: TemperatureSensorExtended[] = result.temperatureSensors.map(
      (data: any) => new TemperatureSensorExtended(data)
    );
    return tempSensors;
  }

  static async getTemperatureReadings(page?: number, perPage?: number) {
    const params = new URLSearchParams();
    if (typeof page === 'number') {
      params.append('page', page.toString());
    }
    if (typeof perPage === 'number') {
      params.append('per_page', perPage.toString());
    }
    const queryString = params.toString();
    let path = '/temperature-readings';
    if (queryString.length > 0) {
      path += `?${queryString}`;
    }
    const result = await this.get(path);
    const tempReadings: TemperatureReading[] = result.temperatureReadings.map(
      (data: any) => new TemperatureReading(data)
    );
    return tempReadings;
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
