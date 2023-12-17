type TemperatureReadingFilter = {
  perPage?: number;
  page?: number;
  sensorID?: string;
  bridge?: string;
  fahrenheit?: boolean;
  updatedSince?: string;
  updatedBefore?: string;
}

export default TemperatureReadingFilter;
