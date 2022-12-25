import React from 'react';
import { Box } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import TemperatureReadingDisplay from './TemperatureReadingDisplay';

interface Props {
  sensor: TemperatureSensorExtended;
}

const TemperatureSensorListItem = ({ sensor }: Props) => {
  return <Box as="li" mb={2}>
    {sensor.name}: <TemperatureReadingDisplay reading={sensor.latestReading} />
    <Box fontSize={1} color="fg.muted">
      {sensor.bridge.name}
    </Box>
  </Box>;
};

export default TemperatureSensorListItem;
