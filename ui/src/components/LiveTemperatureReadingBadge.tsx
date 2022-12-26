import React, { useContext } from 'react';
import { LiveTemperatureReadingContext } from '../contexts/LiveTemperatureReadingContext';
import TemperatureReadingBadge from './TemperatureReadingBadge';

const LiveTemperatureReadingBadge = () => {
  const { liveTemperatureReading } = useContext(LiveTemperatureReadingContext);
  return liveTemperatureReading ? <TemperatureReadingBadge temperatureReading={liveTemperatureReading} /> : null;
};

export default LiveTemperatureReadingBadge;
