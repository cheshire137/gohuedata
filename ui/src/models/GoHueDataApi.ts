import TemperatureReading from "./TemperatureReading";

class GoHueDataApi {
  static apiUrl() {
    return `http://localhost:${process.env.REACT_APP_BACKEND_PORT}/api`;
  }

  static async getTemperatureReadings() {
    const result = await this.get('/temperature-readings');
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
