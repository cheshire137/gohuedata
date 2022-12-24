import React from 'react';
import useGetTemperatureReadings from './hooks/use-get-temperature-readings';
import './App.css';

function App() {
  const { temperatureReadings, fetching, error } = useGetTemperatureReadings({
    page: 1,
    perPage: 100,
  });

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  if (!temperatureReadings) {
    return <p>No data</p>;
  }

  return <ul>
    {temperatureReadings.map(tempReading => <li key={tempReading.id()}>
      ({tempReading.sensor.bridge.name}) {tempReading.sensor.name}: {tempReading.temperature}&deg; {tempReading.units} as of {tempReading.lastUpdated}
    </li>)}
  </ul>;
}

export default App;
