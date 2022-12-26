import React, { useContext, useMemo } from 'react';
import { Box } from '@primer/react';
import { TemperatureReadingsContext } from '../contexts/TemperatureReadingsContext';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

const chartOptions = {
  responsive: true,
  plugins: {
    legend: {
      position: 'top' as const,
    },
    title: {
      display: true,
      text: 'Temperatures over time',
    },
  },
};

const TemperatureReadingGraph = () => {
  const { temperatureReadings } = useContext(TemperatureReadingsContext);
  const labels = useMemo(() => temperatureReadings.map(tempReading => tempReading.timestamp), [temperatureReadings]);
  const data = {
    labels,
    datasets: [
      {
        label: 'Temperature',
        data: temperatureReadings.map(tempReading => tempReading.temperature),
      }
    ]
  };

  return <Box mb={2}>
    <Line data={data} options={chartOptions} />
  </Box>;
};

export default TemperatureReadingGraph;
