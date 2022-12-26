import React, { useContext } from 'react';
import { Box, Heading, Pagination } from '@primer/react';
import TemperatureReadingListItem from './TemperatureReadingListItem';
import { TemperatureReadingsContext } from '../contexts/TemperatureReadingsContext';

const TemperatureReadingList = () => {
  const { temperatureReadings, page, totalPages, setPage } = useContext(TemperatureReadingsContext);
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
