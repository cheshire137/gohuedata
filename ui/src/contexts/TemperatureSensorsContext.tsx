import React from 'react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';

export type TemperatureSensorsContextProps = {
  temperatureSensors: TemperatureSensorExtended[];
};

export const TemperatureSensorsContext = React.createContext<TemperatureSensorsContextProps>({
  temperatureSensors: [],
});

interface Props {
  children: React.ReactNode;
}

export const TemperatureSensorsContextProvider = ({ children }: Props) => {
  const { temperatureSensors, fetching, error } = useGetTemperatureSensors();

  if (fetching) {
    return <p>Loading temperature sensors...</p>;
  }

  if (error) {
    return <p>Error loading temperature sensors: {error}</p>;
  }

  return <TemperatureSensorsContext.Provider value={{
    temperatureSensors: temperatureSensors!,
  }}>{children}</TemperatureSensorsContext.Provider>;
};
