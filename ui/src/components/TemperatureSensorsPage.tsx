import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';
import TemperatureSensorList from './TemperatureSensorList';

const TemperatureSensorsPage = () => {
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle('Temperature sensors'), [setPageTitle]);

  return <TemperatureSensorList />;
};

export default TemperatureSensorsPage;
