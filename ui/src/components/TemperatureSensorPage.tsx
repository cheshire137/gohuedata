import React, { useContext, useEffect } from 'react';
import { Box, CircleBadge, Text, RelativeTime } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import { useLoaderData } from 'react-router-dom';
import { PageContext } from '../contexts/PageContext';
import { TemperatureReadingsContextProvider } from '../contexts/TemperatureReadingsContext';
import TemperatureReadingList from './TemperatureReadingList';
import TemperatureReadingGraph from './TemperatureReadingGraph';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const sensor = useLoaderData() as TemperatureSensorExtended;
  const { latestReading } = sensor;

  useEffect(() => setPageTitle(`${sensor.bridge.name} / ${sensor.name}`),
    [sensor.name, sensor.bridge.name, setPageTitle]);

  return <Box mb={2}>
    <CircleBadge variant="large" sx={{ flexDirection: 'column' }}>
      <Text
        fontWeight="bold"
        fontSize="5"
      >{Math.round(latestReading.temperature)}&deg;{latestReading.units}</Text>
      <Text
        fontSize={1}
        color="fg.muted"
      ><RelativeTime threshold="P1D" date={latestReading.timestampAsDate()} /></Text>
    </CircleBadge>
    <TemperatureReadingsContextProvider filter={{ sensorID: sensor.id, perPage: 30 }}>
      <TemperatureReadingGraph />
      <TemperatureReadingList />
    </TemperatureReadingsContextProvider>
  </Box>;
};

export default TemperatureSensorPage;
