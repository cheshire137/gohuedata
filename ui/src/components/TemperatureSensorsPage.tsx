import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';
import { TemperatureSensorsContextProvider } from '../contexts/TemperatureSensorsContext';
import TemperatureSensorList from './TemperatureSensorList';

const TemperatureSensorsPage = () => {
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle('Temperature sensors'), [setPageTitle]);

  return <TemperatureSensorsContextProvider>
    <TemperatureSensorList />
  </TemperatureSensorsContextProvider>;
};

export default TemperatureSensorsPage;
