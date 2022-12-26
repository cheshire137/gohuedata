import { useState, useEffect } from 'react';
import GoHueDataApi from '../models/GoHueDataApi';
import TemperatureReading from '../models/TemperatureReading';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';

interface Results {
  temperatureReadings?: TemperatureReading[];
  fetching: boolean;
  error?: string;
}

function useGetTemperatureReadings(filter?: TemperatureReadingFilter): Results {
  const [results, setResults] = useState<Results>({ fetching: true });

  useEffect(() => {
    async function fetchTemperatureReadings() {
      try {
        const temperatureReadings = await GoHueDataApi.getTemperatureReadings(filter);
        setResults({ temperatureReadings, fetching: false })
      } catch (err: any) {
        console.error('failed to fetch temperature readings', err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchTemperatureReadings()
  }, [filter, setResults]);

  return results;
}

export default useGetTemperatureReadings;
