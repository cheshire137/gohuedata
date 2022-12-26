import React, { useContext, useEffect } from 'react';
import { Box } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import TemperatureReadingDisplay from './TemperatureReadingDisplay';
import { useLoaderData } from 'react-router-dom';
import { PageContext } from '../contexts/PageContext';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const sensor = useLoaderData() as TemperatureSensorExtended;

  useEffect(() => setPageTitle(sensor.name), [sensor.name, setPageTitle]);

  return <Box as="li" mb={2}>
    {sensor.name}: <TemperatureReadingDisplay reading={sensor.latestReading} />
    <Box fontSize={1} color="fg.muted">
      {sensor.bridge.name}
    </Box>
  </Box>;
};

export default TemperatureSensorPage;
