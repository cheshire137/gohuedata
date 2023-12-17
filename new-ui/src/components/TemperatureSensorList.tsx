import React, { useContext } from 'react';
import { TemperatureSensorsContext } from '../contexts/TemperatureSensorsContext';
import TemperatureSensorListItem from './TemperatureSensorListItem';

const TemperatureSensorList = () => {
  const { temperatureSensors } = useContext(TemperatureSensorsContext);

  return <ul>
    {temperatureSensors.map(tempSensor => <TemperatureSensorListItem
      key={tempSensor.id}
      sensor={tempSensor}
    />)}
  </ul>;
};

export default TemperatureSensorList;
