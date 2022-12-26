import React, { useState, useEffect, useContext } from 'react';
import TemperatureReading from '../models/TemperatureReading';
import { TemperatureSensorsContext } from '../contexts/TemperatureSensorsContext';

export type LiveTemperatureReadingContextProps = {
  liveTemperatureReading: TemperatureReading | null;
};

export const LiveTemperatureReadingContext = React.createContext<LiveTemperatureReadingContextProps>({
  liveTemperatureReading: null,
});

interface Props {
  sensorID: string;
  children: React.ReactNode;
}

export const LiveTemperatureReadingContextProvider = ({ sensorID, children }: Props) => {
  const { temperatureSensors: liveTempSensors } = useContext(TemperatureSensorsContext);
  const [liveReading, setLiveReading] = useState<null | TemperatureReading>(null);

  useEffect(() => {
    const liveTempSensor = liveTempSensors.find(tempSensor => tempSensor.id === sensorID);
    if (liveTempSensor) {
      setLiveReading(liveTempSensor.latestReading);
    } else {
      setLiveReading(null);
    }
  }, [liveTempSensors, sensorID]);

  return <LiveTemperatureReadingContext.Provider value={{
    liveTemperatureReading: liveReading,
  }}>{children}</LiveTemperatureReadingContext.Provider>;
};
