import React from 'react';
import { Box, Heading } from '@primer/react';
import useGetTemperatureReadings from '../hooks/use-get-temperature-readings';

const TemperatureReadings = () => {
  const { temperatureReadings, fetching, error } = useGetTemperatureReadings({
    page: 1,
    perPage: 100,
  });

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return <Box>
    <Heading as="h2">Latest temperatures</Heading>
    <ul>
      {temperatureReadings && temperatureReadings.map(tempReading => <li key={tempReading.id}>
        ({tempReading.sensor.bridge.name}) {tempReading.sensor.name}: {tempReading.temperature}&deg; {tempReading.units} as of {tempReading.timestamp}
      </li>)}
    </ul>
  </Box>;
};

export default TemperatureReadings;
