import React, { useContext } from 'react';
import { Pagination } from '@primer/react';
import { TemperatureReadingsContext } from '../contexts/TemperatureReadingsContext';

const TemperatureReadingPagination = () => {
  const { page, totalPages, setPage } = useContext(TemperatureReadingsContext);
  if (!totalPages || !page || totalPages <= 1) return null;
  return <Pagination
    pageCount={totalPages}
    currentPage={page}
    onPageChange={(e, newPage) => {
      e.preventDefault();
      setPage(newPage);
    }}
  />;
};

export default TemperatureReadingPagination;
