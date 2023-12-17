import React from 'react';
import { RelativeTime } from '@primer/react';
import TemperatureReading from '../models/TemperatureReading';
import TemperatureBadge from './TemperatureBadge';

interface Props {
  temperatureReading: TemperatureReading;
}

const TemperatureReadingBadge = ({ temperatureReading }: Props) => <TemperatureBadge
  temperature={temperatureReading.temperature}
  units={temperatureReading.units}
><RelativeTime threshold="P1D" date={temperatureReading.timestampAsDate()} /></TemperatureBadge>;

export default TemperatureReadingBadge;
