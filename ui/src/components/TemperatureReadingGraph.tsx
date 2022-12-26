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
  const units = temperatureReadings.length > 0 ? temperatureReadings[0].units : 'F';
  const thermScale = units === 'F' ? 'Fahrenheit' : 'Celsius';
  const labels = useMemo(() => temperatureReadings.map(tempReading => tempReading.timestamp), [temperatureReadings]);
  const data = {
    labels,
    datasets: [
      {
        label: `Temperature in ${thermScale}`,
        data: temperatureReadings.map(tempReading => tempReading.temperature),
        borderColor: 'rgb(53, 162, 235)',
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
      }
    ]
  };

  return <Box mb={2}>
    <Line data={data} options={chartOptions} />
  </Box>;
};

export default TemperatureReadingGraph;
