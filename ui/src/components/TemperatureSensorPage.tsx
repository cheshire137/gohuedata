import React, { useContext, useEffect } from 'react';
import { Box } from '@primer/react';
import type TemperatureSensorResult from '../types/TemperatureSensorResult';
import { useLoaderData } from 'react-router-dom';
import { PageContext } from '../contexts/PageContext';
import { TemperatureReadingsContextProvider } from '../contexts/TemperatureReadingsContext';
import { LiveTemperatureReadingContextProvider } from '../contexts/LiveTemperatureReadingContext';
import TemperatureReadingList from './TemperatureReadingList';
import TemperatureReadingGraph from './TemperatureReadingGraph';
import LiveTemperatureReadingBadge from './LiveTemperatureReadingBadge';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const data = useLoaderData() as TemperatureSensorResult;
  const sensor = data.temperatureSensor;

  useEffect(() => setPageTitle(`${sensor.bridge.name} / ${sensor.name}`),
    [sensor.name, sensor.bridge.name, setPageTitle]);

  return <Box mb={2}>
    <LiveTemperatureReadingContextProvider sensorID={sensor.id}>
      <LiveTemperatureReadingBadge />
      <TemperatureReadingsContextProvider filter={{ sensorID: sensor.id, perPage: 30 }}>
        <TemperatureReadingGraph />
        <TemperatureReadingList />
      </TemperatureReadingsContextProvider>
    </LiveTemperatureReadingContextProvider>
  </Box>;
};

export default TemperatureSensorPage;
