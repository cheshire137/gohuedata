import React, { useState, useEffect } from 'react';
import { Box, Heading, Pagination } from '@primer/react';
import useGetTemperatureReadings from '../hooks/use-get-temperature-readings';
import type TemperatureReadingFilter from '../types/TemperatureReadingFilter';
import TemperatureReadingListItem from './TemperatureReadingListItem';

const TemperatureReadingList = (filter?: TemperatureReadingFilter) => {
  const [page, setPage] = useState(filter?.page || 1);
  const { temperatureReadings, totalPages, fetching, error } = useGetTemperatureReadings({page, ...filter});

  useEffect(() => setPage(filter?.page || 1), [filter?.page]);

  if (fetching) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return <Box>
    <Heading as="h2">Latest temperatures</Heading>
    <ul>
      {temperatureReadings && temperatureReadings.map(tempReading => <TemperatureReadingListItem
        reading={tempReading}
        key={tempReading.id}
      />)}
    </ul>
    {totalPages && page && totalPages > 1 && <Pagination
      pageCount={totalPages}
      currentPage={page}
      onPageChange={(e, newPage) => {
        e.preventDefault();
        setPage(newPage);
      }}
    />}
  </Box>;
};

export default TemperatureReadingList;
