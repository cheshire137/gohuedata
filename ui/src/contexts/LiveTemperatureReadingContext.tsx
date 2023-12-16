import React, { createContext, PropsWithChildren, useMemo, useState, useEffect, useContext } from 'react';
import TemperatureReading from '../models/TemperatureReading';
import { TemperatureSensorsContext } from '../contexts/TemperatureSensorsContext';

export type LiveTemperatureReadingContextProps = {
  liveTemperatureReading: TemperatureReading | null;
};

export const LiveTemperatureReadingContext = createContext<LiveTemperatureReadingContextProps>({
  liveTemperatureReading: null,
});

interface Props extends PropsWithChildren {
  sensorID: string;
}

export const LiveTemperatureReadingContextProvider = ({ sensorID, children }: Props) => {
  const { temperatureSensors: liveTempSensors } = useContext(TemperatureSensorsContext);
  const [liveReading, setLiveReading] = useState<null | TemperatureReading>(null);
  const contextProps = useMemo(() => ({
    liveTemperatureReading: liveReading,
  } satisfies LiveTemperatureReadingContextProps), [liveReading]);

  useEffect(() => {
    const liveTempSensor = liveTempSensors.find(tempSensor => tempSensor.id === sensorID);
    if (liveTempSensor) {
      setLiveReading(liveTempSensor.latestReading);
    } else {
      setLiveReading(null);
    }
  }, [liveTempSensors, sensorID]);

  return <LiveTemperatureReadingContext.Provider value={contextProps}>
    {children}
  </LiveTemperatureReadingContext.Provider>;
};
