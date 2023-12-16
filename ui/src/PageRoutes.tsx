import React, { useContext } from 'react';
import { TemperatureSensorsContextProvider } from './contexts/TemperatureSensorsContext';
import { GroupsContextProvider } from './contexts/GroupsContext';
import { createHashRouter, RouterProvider, createRoutesFromElements, Route } from 'react-router-dom';
import TemperatureSensorsPage from './components/TemperatureSensorsPage';
import GroupsPage from './components/GroupsPage';
import TemperatureSensorPage from './components/TemperatureSensorPage';
import AppLayout from './components/AppLayout';
import ErrorPage from './components/ErrorPage';
import GoHueDataApi from './models/GoHueDataApi';
import { SettingsContext } from './contexts/SettingsContext';

const PageRoutes = () => {
  const { fahrenheit } = useContext(SettingsContext);
  const router = createHashRouter(createRoutesFromElements(
    <Route element={<AppLayout />}>
      <Route path="/" element={<TemperatureSensorsPage />} errorElement={<ErrorPage />} />
      <Route
        path="/sensor/:id"
        loader={async ({ params }) => {
          const sensorID = params.id!;
          return await GoHueDataApi.getTemperatureSensor(sensorID, fahrenheit);
        }}
        element={<TemperatureSensorPage />}
        errorElement={<ErrorPage />}
      />
      <Route path="/groups" element={<GroupsPage />} errorElement={<ErrorPage />} />
    </Route>
  ));

  return <TemperatureSensorsContextProvider fahrenheit={fahrenheit}>
    <GroupsContextProvider>
      <RouterProvider router={router} />
    </GroupsContextProvider>
  </TemperatureSensorsContextProvider>;
}

export default PageRoutes;
