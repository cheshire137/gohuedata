import TemperatureReading from "./TemperatureReading";

class GoHueDataApi {
  apiUrl: string;

  constructor(port: number) {
    this.apiUrl = `http://localhost:${port}/api`;
  }

  async getTemperatureReadings() {
    const result = await this.get('/temperature-readings');
    console.log(result);
    const tempReadings: TemperatureReading[] = result.temperatureReadings.map(
      (data: any) => new TemperatureReading(data)
    );
    return tempReadings;
  }

  async get(path: string) {
    const response = await fetch(`${this.apiUrl}${path}`);
    if (response.status >= 200 && response.status < 300) {
      const result = await response.json();
      return result
    }
    throw new Error(response.statusText);
  }
}

export default GoHueDataApi;
