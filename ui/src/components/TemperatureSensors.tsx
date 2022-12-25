import React from 'react';
import { Box, Heading } from '@primer/react';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';

const TemperatureSensors = () => {
  const { temperatureSensors, fetching, error } = useGetTemperatureSensors();

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return <Box mb={4}>
    <Heading as="h2">Temperature sensors</Heading>
    <ul>
      {temperatureSensors && temperatureSensors.map(tempSensor => <li key={tempSensor.id}>
        ({tempSensor.bridge.name}) {tempSensor.name}: {tempSensor.latestReading.temperature}&deg; {tempSensor.latestReading.units} as of {tempSensor.lastUpdated}
      </li>)}
    </ul>
  </Box>;
};

export default TemperatureSensors;
