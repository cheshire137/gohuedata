import React from 'react';
import logo from './logo.svg';
import GoHueDataApi from './models/GoHueDataApi';
import useGetTemperatureReadings from './hooks/use-get-temperature-readings';
import './App.css';

function App() {
  const api = new GoHueDataApi(8080);
  const { temperatureReadings, fetching, error } = useGetTemperatureReadings(api);

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  if (!temperatureReadings) {
    return <p>No data</p>;
  }

  return <div className="App">
    {temperatureReadings.map(tempReading => <p key={tempReading.id()}>
      {tempReading.temperature}&deg; {tempReading.units} as of {tempReading.lastUpdated}
    </p>)}
  </div>;
}

export default App;
