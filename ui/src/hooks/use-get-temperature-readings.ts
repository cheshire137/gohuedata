import { useState, useEffect } from "react";
import GoHueDataApi from '../models/GoHueDataApi';
import TemperatureReading from "../models/TemperatureReading";

interface Results {
  temperatureReadings?: TemperatureReading[];
  fetching: boolean;
  error?: string;
}

interface Props {
  page?: number;
  perPage?: number;
}

function useGetTemperatureReadings({ page, perPage }: Props): Results {
  const [results, setResults] = useState<Results>({ fetching: true });

  useEffect(() => {
    async function fetchTemperatureReadings() {
      try {
        const temperatureReadings = await GoHueDataApi.getTemperatureReadings(page, perPage);
        setResults({ temperatureReadings, fetching: false })
      } catch (err: any) {
        console.error('failed to fetch temperature readings', err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchTemperatureReadings()
  }, [page, perPage]);

  return results;
}

export default useGetTemperatureReadings;
