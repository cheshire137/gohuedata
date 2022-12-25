import React from 'react';
import { Box, Heading } from '@primer/react';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';
import TemperatureSensorListItem from './TemperatureSensorListItem';

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
      {temperatureSensors && temperatureSensors.map(tempSensor => <TemperatureSensorListItem
        key={tempSensor.id}
        sensor={tempSensor}
      />)}
    </ul>
  </Box>;
};

export default TemperatureSensors;
