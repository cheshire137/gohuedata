import React from 'react';
import TemperatureReading from '../models/TemperatureReading';
import TemperatureReadingDisplay from './TemperatureReadingDisplay';

interface Props {
  reading: TemperatureReading;
}

const TemperatureReadingListItem = ({ reading }: Props) => {
  return <li>
    <TemperatureReadingDisplay reading={reading} />
  </li>;
};

export default TemperatureReadingListItem;
