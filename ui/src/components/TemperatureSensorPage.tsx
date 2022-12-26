import React, { useContext, useEffect, useState } from 'react';
import { Box, CircleBadge, Text, RelativeTime } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import TemperatureReading from '../models/TemperatureReading';
import { useLoaderData } from 'react-router-dom';
import { TemperatureSensorsContext } from '../contexts/TemperatureSensorsContext';
import { PageContext } from '../contexts/PageContext';
import { TemperatureReadingsContextProvider } from '../contexts/TemperatureReadingsContext';
import TemperatureReadingList from './TemperatureReadingList';
import TemperatureReadingGraph from './TemperatureReadingGraph';
import TemperatureReadingBadge from './TemperatureReadingBadge';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const { temperatureSensors: liveTempSensors } = useContext(TemperatureSensorsContext);
  const [liveReading, setLiveReading] = useState<null | TemperatureReading>(null);
  const sensor = useLoaderData() as TemperatureSensorExtended;

  useEffect(() => {
    const liveTempSensor = liveTempSensors.find(tempSensor => tempSensor.id === sensor.id);
    if (liveTempSensor) {
      setLiveReading(liveTempSensor.latestReading);
    }
  }, [liveTempSensors, sensor.id]);

  useEffect(() => setPageTitle(`${sensor.bridge.name} / ${sensor.name}`),
    [sensor.name, sensor.bridge.name, setPageTitle]);

  return <Box mb={2}>
    {liveReading && <TemperatureReadingBadge temperatureReading={liveReading} />}
    <TemperatureReadingsContextProvider filter={{ sensorID: sensor.id, perPage: 30 }}>
      <TemperatureReadingGraph />
      <TemperatureReadingList />
    </TemperatureReadingsContextProvider>
  </Box>;
};

export default TemperatureSensorPage;
