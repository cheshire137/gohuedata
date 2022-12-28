import React, { useState, useEffect, useContext } from 'react';
import TemperatureReading from '../models/TemperatureReading';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';
import useGetTemperatureReadings from '../hooks/use-get-temperature-readings';
import { SettingsContext } from './SettingsContext';

export type TemperatureReadingsContextProps = {
  temperatureReadings: TemperatureReading[];
  page: number;
  perPage: number;
  totalPages: number;
  totalCount: number;
  setPage(page: number): void;
};

export const TemperatureReadingsContext = React.createContext<TemperatureReadingsContextProps>({
  temperatureReadings: [],
  page: 1,
  perPage: 10,
  totalPages: 1,
  totalCount: 0,
  setPage: () => { },
});

interface Props {
  filter?: TemperatureReadingFilter;
  children: React.ReactNode;
}

export const TemperatureReadingsContextProvider = ({ filter, children }: Props) => {
  const [page, setPage] = useState(filter?.page || 1);
  const { fahrenheit, setFahrenheit } = useContext(SettingsContext);
  const { temperatureReadings, totalPages, perPage, totalCount, fetching, error } = useGetTemperatureReadings({
    page, fahrenheit, ...filter
  });

  useEffect(() => setPage(filter?.page || 1), [filter?.page]);

  useEffect(() => {
    if (typeof filter?.fahrenheit === 'boolean') {
      setFahrenheit(filter.fahrenheit);
    }
  }, [filter?.fahrenheit, setFahrenheit]);

  if (fetching) {
    return <p>Loading temperature history...</p>;
  }

  if (error) {
    return <p>Error loading temperature history: {error}</p>;
  }

  return <TemperatureReadingsContext.Provider value={{
    temperatureReadings: temperatureReadings!,
    totalPages: totalPages!,
    page,
    perPage: perPage!,
    totalCount: totalCount!,
    setPage,
  }}>{children}</TemperatureReadingsContext.Provider>;
};
