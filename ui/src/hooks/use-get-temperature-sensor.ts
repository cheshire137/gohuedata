import { useState, useEffect } from "react";
import GoHueDataApi from '../models/GoHueDataApi';
import TemperatureSensorExtended from "../models/TemperatureSensorExtended";

interface Results {
  temperatureSensor?: TemperatureSensorExtended;
  fetching: boolean;
  error?: string;
}

function useGetTemperatureSensor(id: string): Results {
  const [results, setResults] = useState<Results>({ fetching: true });

  useEffect(() => {
    async function fetchTemperatureSensors() {
      try {
        const temperatureSensor = await GoHueDataApi.getTemperatureSensor(id);
        setResults({ temperatureSensor, fetching: false })
      } catch (err: any) {
        console.error(`failed to fetch temperature sensor ${id}`, err);
        setResults({ fetching: false, error: err.message });
      }
    }

    fetchTemperatureSensors()
  }, [id]);

  return results;
}

export default useGetTemperatureSensor;
