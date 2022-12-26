import React from 'react';
import { RelativeTime, Text } from '@primer/react';
import TemperatureReading from '../models/TemperatureReading';

interface Props {
  reading: TemperatureReading;
}

const TemperatureReadingDisplay = ({ reading }: Props) => {
  return <span>
    <span>{Math.round(reading.temperature)}&deg;{reading.units} </span>
    <Text
      display="inline-block"
      fontSize={1}
      color="fg.muted"
    >as of <RelativeTime threshold="P1D" date={reading.timestampAsDate()} /></Text>
  </span>;
};

export default TemperatureReadingDisplay;
