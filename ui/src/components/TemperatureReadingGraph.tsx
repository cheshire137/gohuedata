import React, { useContext, useMemo } from 'react';
import { Box } from '@primer/react';
import { TemperatureReadingsContext } from '../contexts/TemperatureReadingsContext';
import { AxisOptions, Chart, ChartOptions } from 'react-charts';

type HistoricalTemp = {
  date: Date | string;
  temperature: number;
}

type Series = {
  label: string;
  data: HistoricalTemp[];
};

const TemperatureReadingGraph = () => {
  const { temperatureReadings } = useContext(TemperatureReadingsContext);
  const data: Series[] = [{
    label: 'Temperatures',
    data: temperatureReadings.map(tempReading => {
      const date = tempReading.timestampAsDate() || tempReading.timestamp;
      const historicalTemp: HistoricalTemp = { date, temperature: tempReading.temperature };
      return historicalTemp;
    }),
  }];
  const primaryAxis = useMemo((): AxisOptions<HistoricalTemp> => ({ getValue: datum => datum.date }), []);
  const secondaryAxes = useMemo((): AxisOptions<HistoricalTemp>[] => [{
    getValue: datum => datum.temperature,
    elementType: 'line',
  }], []);
  const chartOptions: ChartOptions<HistoricalTemp> = { data, primaryAxis, secondaryAxes, initialHeight: 300 };

  return <Box mb={2}>
    <Chart options={chartOptions} />
  </Box>;
};

export default TemperatureReadingGraph;
