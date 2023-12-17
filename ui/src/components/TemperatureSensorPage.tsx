import React, { useContext, useEffect, useMemo } from 'react';
import { Box } from '@primer/react';
import type TemperatureSensorResult from '../types/TemperatureSensorResult';
import { useLoaderData } from 'react-router-dom';
import { PageContext } from '../contexts/PageContext';
import { SettingsContext } from '../contexts/SettingsContext';
import { TemperatureReadingsContextProvider } from '../contexts/TemperatureReadingsContext';
import { LiveTemperatureReadingContextProvider } from '../contexts/LiveTemperatureReadingContext';
import TemperatureReadingPagination from './TemperatureReadingPagination';
import TemperatureReadingGraph from './TemperatureReadingGraph';
import LiveTemperatureReadingBadge from './LiveTemperatureReadingBadge';
import TemperatureBadge from './TemperatureBadge';

const TemperatureSensorPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const { units } = useContext(SettingsContext);
  const data = useLoaderData() as TemperatureSensorResult;
  const { temperatureSensor, minTemperature, maxTemperature, avgTemperature } = data;
  const filter = useMemo(() => ({ sensorID: temperatureSensor.id, perPage: 30 }), [temperatureSensor.id])

  useEffect(() => setPageTitle(`${temperatureSensor.bridge.name} / ${temperatureSensor.name}`),
    [temperatureSensor.name, temperatureSensor.bridge.name, setPageTitle]);

  return <Box mb={2}>
    <LiveTemperatureReadingContextProvider sensorID={temperatureSensor.id}>
      <Box display="flex" alignItems="center" justifyContent="space-around" mb={3}>
        <LiveTemperatureReadingBadge />
        {minTemperature && <TemperatureBadge
          temperature={minTemperature}
          units={units}
        >Min</TemperatureBadge>}
        {avgTemperature && <TemperatureBadge
          temperature={avgTemperature}
          units={units}
        >Average</TemperatureBadge>}
        {maxTemperature && <TemperatureBadge
          temperature={maxTemperature}
          units={units}
        >Max</TemperatureBadge>}
      </Box>
      <TemperatureReadingsContextProvider filter={filter}>
        <TemperatureReadingGraph />
        <TemperatureReadingPagination />
      </TemperatureReadingsContextProvider>
    </LiveTemperatureReadingContextProvider>
  </Box>;
};

export default TemperatureSensorPage;
