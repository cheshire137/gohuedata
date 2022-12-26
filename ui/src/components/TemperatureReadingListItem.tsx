import React from 'react';
import { Box, Heading } from '@primer/react';
import TemperatureReading from '../models/TemperatureReading';

interface Props {
  reading: TemperatureReading;
}

const TemperatureReadingListItem = ({ reading }: Props) => {
  return <li>
    ({reading.sensor.bridge.name}) {reading.sensor.name}: {reading.temperature}&deg; {reading.units} as of
    {reading.timestampAsDate()?.toLocaleDateString()} {reading.timestampAsDate()?.toLocaleTimeString()}
  </li>;
};

export default TemperatureReadingListItem;
