import React from 'react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';
import PageHeader from '../components/PageHeader';
import { PageLayout } from '@primer/react';

export type TemperatureSensorsContextProps = {
  temperatureSensors: TemperatureSensorExtended[];
};

export const TemperatureSensorsContext = React.createContext<TemperatureSensorsContextProps>({
  temperatureSensors: [],
});

interface Props {
  fahrenheit?: boolean;
  children: React.ReactNode;
}

export const TemperatureSensorsContextProvider = ({ fahrenheit, children }: Props) => {
  const { temperatureSensors, fetching, error } = useGetTemperatureSensors(fahrenheit);

  if (fetching) {
    return <PageLayout>
      <PageHeader />
      <PageLayout.Content>
        <p>Loading temperature sensors...</p>
      </PageLayout.Content>
    </PageLayout>;
  }

  if (error) {
    return <PageLayout>
      <PageHeader />
      <PageLayout.Content>
        <p>Error loading temperature sensors: {error}</p>
      </PageLayout.Content>
    </PageLayout>;
  }

  return <TemperatureSensorsContext.Provider value={{
    temperatureSensors: temperatureSensors!,
  }}>{children}</TemperatureSensorsContext.Provider>;
};
