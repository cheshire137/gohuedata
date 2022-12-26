import React from 'react';
import { CircleBadge, Text, RelativeTime } from '@primer/react';
import TemperatureReading from '../models/TemperatureReading';

interface Props {
  temperatureReading: TemperatureReading;
}

const TemperatureReadingBadge = ({ temperatureReading }: Props) => <CircleBadge variant="large" sx={{ flexDirection: 'column' }}>
  <Text
    fontWeight="bold"
    fontSize="5"
  >{Math.round(temperatureReading.temperature)}&deg;{temperatureReading.units}</Text>
  <Text
    fontSize={1}
    color="fg.muted"
  ><RelativeTime threshold="P1D" date={temperatureReading.timestampAsDate()} /></Text>
</CircleBadge>;

export default TemperatureReadingBadge;
