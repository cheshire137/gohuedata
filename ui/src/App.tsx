import React from 'react';
import { BaseStyles, ThemeProvider } from '@primer/react';
import { PageContextProvider } from './contexts/PageContext';
import { SettingsContextProvider } from './contexts/SettingsContext';
import PageRoutes from './PageRoutes';

function App() {
  return <ThemeProvider>
    <BaseStyles>
      <PageContextProvider>
        <SettingsContextProvider>
          <PageRoutes />
        </SettingsContextProvider>
      </PageContextProvider>
    </BaseStyles>
  </ThemeProvider>;
}

export default App;
