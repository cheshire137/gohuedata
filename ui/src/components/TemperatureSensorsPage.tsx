import React, { useContext, useEffect } from 'react';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';
import { PageContext } from '../contexts/PageContext';
import TemperatureSensorListItem from './TemperatureSensorListItem';

const TemperatureSensorsPage = () => {
  const { temperatureSensors, fetching, error } = useGetTemperatureSensors();
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle('Temperature sensors'), [setPageTitle]);

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return <ul>
    {temperatureSensors && temperatureSensors.map(tempSensor => <TemperatureSensorListItem
      key={tempSensor.id}
      sensor={tempSensor}
    />)}
  </ul>;
};

export default TemperatureSensorsPage;
