import { useState, useEffect } from "react";
import GoHueDataApi from '../models/GoHueDataApi';
import TemperatureSensorExtended from "../models/TemperatureSensorExtended";

interface Results {
  temperatureSensors?: TemperatureSensorExtended[];
  fetching: boolean;
  error?: string;
}

function useGetTemperatureSensors(fahrenheit?: boolean): Results {
  const [results, setResults] = useState<Results>({ fetching: true });

  useEffect(() => {
    async function fetchTemperatureSensors() {
      try {
        const temperatureSensors = await GoHueDataApi.getLiveTemperatureSensors(fahrenheit);
        setResults({ temperatureSensors, fetching: false })
      } catch (err: any) {
        console.error('failed to fetch temperature sensors', err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchTemperatureSensors()
  }, [fahrenheit]);

  return results;
}

export default useGetTemperatureSensors;
