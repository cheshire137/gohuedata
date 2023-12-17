import { useState, useEffect } from 'react';
import GoHueDataApi from '../models/GoHueDataApi';
import TemperatureReading from '../models/TemperatureReading';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';

interface Results {
  temperatureReadings?: TemperatureReading[];
  page?: number;
  perPage?: number;
  totalPages?: number;
  totalCount?: number;
  fetching: boolean;
  error?: string;
}

function useGetTemperatureReadings(filter?: TemperatureReadingFilter): Results {
  const [results, setResults] = useState<Results>({ fetching: true });
  const { page, perPage, sensorID, bridge, fahrenheit, updatedBefore, updatedSince } = filter || {};

  useEffect(() => {
    async function fetchTemperatureReadings() {
      try {
        const result = await GoHueDataApi.getTemperatureReadings({
          fahrenheit,
          page,
          perPage,
          sensorID,
          bridge,
          updatedBefore,
          updatedSince,
        });
        setResults({ ...result, fetching: false })
      } catch (err: any) {
        console.error('failed to fetch temperature readings', err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchTemperatureReadings()
  }, [page, perPage, sensorID, bridge, setResults, fahrenheit, updatedBefore, updatedSince]);

  return results;
}

export default useGetTemperatureReadings;
