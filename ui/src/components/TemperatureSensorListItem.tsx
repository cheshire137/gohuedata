import React from 'react';
import { RelativeTime } from '@primer/react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';

interface Props {
  sensor: TemperatureSensorExtended;
}

const TemperatureSensorListItem = ({ sensor }: Props) => {
  return <li>
    ({sensor.bridge.name}) {sensor.name}: {sensor.latestReading.temperature}&deg; {sensor.latestReading.units} as of <RelativeTime date={sensor.lastUpdatedAt()} />
  </li>;
};

export default TemperatureSensorListItem;
