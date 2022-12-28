import React, { useContext, useMemo, useState, useEffect } from 'react';
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

const defaultUnits = 'F';
const thermScaleFor = (units: string) => units === 'F' ? 'Fahrenheit' : 'Celsius';

const TemperatureReadingGraph = () => {
  const { temperatureReadings } = useContext(TemperatureReadingsContext);
  const [units, setUnits] = useState(defaultUnits);
  const [thermScale, setThermScale] = useState(thermScaleFor(units));
  const [sortedReadings, setSortedReadings] = useState(temperatureReadings);
  const [labels, setLabels] = useState(temperatureReadings.map(tempReading => tempReading.timestamp));

  useEffect(() => {
    setUnits(temperatureReadings.length > 0 ? temperatureReadings[0].units : defaultUnits);
  }, [temperatureReadings, setUnits]);

  useEffect(() => {
    setThermScale(thermScaleFor(units));
  }, [units, setThermScale]);

  useEffect(() => {
    const newSortedReadings = [...temperatureReadings].sort((a, b) => a.timestamp.localeCompare(b.timestamp));
    setSortedReadings(newSortedReadings);
  }, [temperatureReadings]);

  useEffect(() => {
    const newLabels = sortedReadings.map(tempReading => {
      const date = tempReading.timestampAsDate();
      if (date) {
        return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', hour: 'numeric' });
      }
      return tempReading.timestamp;
    });
    setLabels(newLabels);
  }, [sortedReadings]);

  return <Box mb={2} height="400px">
    <Line data={{
      labels,
      datasets: [
        {
          label: `Temperature in ${thermScale}`,
          data: sortedReadings.map(tempReading => tempReading.temperature),
          borderColor: 'rgb(53, 162, 235)',
          backgroundColor: 'rgba(53, 162, 235, 0.5)',
          tension: 0.4,
        }
      ]
    }} options={{
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
