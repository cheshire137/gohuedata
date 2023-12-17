import React from 'react';
import { Box, Link } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import TemperatureReadingDisplay from './TemperatureReadingDisplay';
import { useHref } from 'react-router-dom';

interface Props {
  sensor: TemperatureSensorExtended;
}

const TemperatureSensorListItem = ({ sensor }: Props) => {
  return <Box as="li" mb={2}>
    <Link href={useHref(`/sensor/${sensor.id}`)}>
      {sensor.name}: <TemperatureReadingDisplay reading={sensor.latestReading} />
    </Link>
    <Box fontSize={1} display="inline-block" color="fg.muted">
      {sensor.bridge.name}
    </Box>
  </Box>;
};

export default TemperatureSensorListItem;
