import React from 'react';
import { BaseStyles, ThemeProvider } from '@primer/react';
import { PageContextProvider } from './contexts/PageContext';
import { createHashRouter, RouterProvider, createRoutesFromElements, Route } from 'react-router-dom';
import TemperatureSensorsPage from './components/TemperatureSensorsPage';
import AppLayout from './components/AppLayout';
import ErrorPage from './components/ErrorPage';

function App() {
  const router = createHashRouter(createRoutesFromElements(
    <Route element={<AppLayout />}>
      <Route path="/" element={<TemperatureSensorsPage />} errorElement={<ErrorPage />} />
    </Route>
  ));

  return <ThemeProvider>
    <BaseStyles>
      <PageContextProvider>
        <RouterProvider router={router} />
      </PageContextProvider>
    </BaseStyles>
  </ThemeProvider>;
}

export default App;
