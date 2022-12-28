import React, { useContext, useEffect } from 'react';
import { Box } from '@primer/react';
import type TemperatureSensorResult from '../types/TemperatureSensorResult';
import { useLoaderData } from 'react-router-dom';
import { PageContext } from '../contexts/PageContext';
import { TemperatureReadingsContextProvider } from '../contexts/TemperatureReadingsContext';
import { LiveTemperatureReadingContextProvider } from '../contexts/LiveTemperatureReadingContext';
import TemperatureReadingPagination from './TemperatureReadingPagination';
import TemperatureReadingGraph from './TemperatureReadingGraph';
import LiveTemperatureReadingBadge from './LiveTemperatureReadingBadge';
import TemperatureBadge from './TemperatureBadge';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const data = useLoaderData() as TemperatureSensorResult;
  const { temperatureSensor, minTemperature, maxTemperature, avgTemperature } = data;

  useEffect(() => setPageTitle(`${temperatureSensor.bridge.name} / ${temperatureSensor.name}`),
    [temperatureSensor.name, temperatureSensor.bridge.name, setPageTitle]);

  return <Box mb={2}>
    <LiveTemperatureReadingContextProvider sensorID={temperatureSensor.id}>
      <Box display="flex" alignItems="center" justifyContent="space-around" mb={3}>
        <LiveTemperatureReadingBadge />
        {minTemperature && <TemperatureBadge
          temperature={minTemperature}
          units="F"
        >Min</TemperatureBadge>}
        {avgTemperature && <TemperatureBadge
          temperature={avgTemperature}
          units="F"
        >Average</TemperatureBadge>}
        {maxTemperature && <TemperatureBadge
          temperature={maxTemperature}
          units="F"
        >Max</TemperatureBadge>}
      </Box>
      <TemperatureReadingsContextProvider filter={{ sensorID: temperatureSensor.id, perPage: 30 }}>
        <TemperatureReadingGraph />
        <TemperatureReadingPagination />
      </TemperatureReadingsContextProvider>
    </LiveTemperatureReadingContextProvider>
  </Box>;
};

export default TemperatureSensorPage;
