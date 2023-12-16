import React, { createContext, PropsWithChildren, useMemo } from 'react';
import TemperatureSensorExtended from '../models/TemperatureSensorExtended';
import useGetTemperatureSensors from '../hooks/use-get-temperature-sensors';
import PageHeader from '../components/PageHeader';
import { Flash, PageLayout } from '@primer/react';

export type TemperatureSensorsContextProps = {
  temperatureSensors: TemperatureSensorExtended[];
};

export const TemperatureSensorsContext = createContext<TemperatureSensorsContextProps>({
  temperatureSensors: [],
});

interface Props extends PropsWithChildren {
  fahrenheit?: boolean;
}

export const TemperatureSensorsContextProvider = ({ fahrenheit, children }: Props) => {
  const { temperatureSensors, fetching, error } = useGetTemperatureSensors(fahrenheit);
  const contextProps = useMemo(() => ({
    temperatureSensors: temperatureSensors!
  } satisfies TemperatureSensorsContextProps), [temperatureSensors]);

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
        <Flash variant="danger">Error loading temperature sensors: {error}</Flash>
      </PageLayout.Content>
    </PageLayout>;
  }

  return <TemperatureSensorsContext.Provider value={contextProps}>{children}</TemperatureSensorsContext.Provider>;
};
