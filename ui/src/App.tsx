import React from 'react';
import { BaseStyles, ThemeProvider, Header, Heading, PageLayout } from '@primer/react';
import { PageContextProvider } from './contexts/PageContext';
import TemperatureReadings from './components/TemperatureReadings';
import TemperatureSensors from './components/TemperatureSensors';

function App() {
  return <ThemeProvider>
    <BaseStyles>
      <PageContextProvider>
        <PageLayout>
          <PageLayout.Header>
            <Header>
              <Header.Item>
                <Heading as="h1">gohuedata</Heading>
              </Header.Item>
            </Header>
          </PageLayout.Header>
          <PageLayout.Content sx={{ fontSize: 2 }}>
            <TemperatureSensors />
            <TemperatureReadings />
          </PageLayout.Content>
        </PageLayout>
      </PageContextProvider>
    </BaseStyles>
  </ThemeProvider>;
}

export default App;
