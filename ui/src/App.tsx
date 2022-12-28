import React from 'react';
import { BaseStyles, ThemeProvider } from '@primer/react';
import { PageContextProvider } from './contexts/PageContext';
import { TemperatureSensorsContextProvider } from './contexts/TemperatureSensorsContext';
import { createHashRouter, RouterProvider, createRoutesFromElements, Route } from 'react-router-dom';
import TemperatureSensorsPage from './components/TemperatureSensorsPage';
import TemperatureSensorPage from './components/TemperatureSensorPage';
import AppLayout from './components/AppLayout';
import ErrorPage from './components/ErrorPage';
import GoHueDataApi from './models/GoHueDataApi';
import { SettingsContextProvider } from './contexts/SettingsContext';

function App() {
  const router = createHashRouter(createRoutesFromElements(
    <Route element={<AppLayout />}>
      <Route path="/" element={<TemperatureSensorsPage />} errorElement={<ErrorPage />} />
      <Route
        path="/sensor/:id"
        loader={async ({ params }) => {
          const fahrenheit = true;
          const sensorID = params.id!;
          return await GoHueDataApi.getTemperatureSensor(sensorID, fahrenheit);
        }}
        element={<TemperatureSensorPage />}
        errorElement={<ErrorPage />}
      />
    </Route>
  ));

  return <ThemeProvider>
    <BaseStyles>
      <PageContextProvider>
        <SettingsContextProvider>
          <TemperatureSensorsContextProvider>
            <RouterProvider router={router} />
          </TemperatureSensorsContextProvider>
        </SettingsContextProvider>
      </PageContextProvider>
    </BaseStyles>
  </ThemeProvider>;
}

export default App;
