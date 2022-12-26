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

const TemperatureReadingGraph = () => {
  const { temperatureReadings } = useContext(TemperatureReadingsContext);
  const units = temperatureReadings.length > 0 ? temperatureReadings[0].units : 'F';
  const thermScale = units === 'F' ? 'Fahrenheit' : 'Celsius';
  const labels = useMemo(() => temperatureReadings.map(tempReading => {
    const date = tempReading.timestampAsDate();
    if (date) {
      return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', hour: 'numeric' });
    }
    return tempReading.timestamp;
  }), [temperatureReadings]);
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

  return <Box mb={2} height="400px">
    <Line data={data} options={{
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
      scales: {
        x: { ticks: { autoSkip: true } },
        y: { ticks: { callback: value => Number.isInteger(value) ? value : null } }
      },
    }} />
  </Box>;
};

export default TemperatureReadingGraph;
