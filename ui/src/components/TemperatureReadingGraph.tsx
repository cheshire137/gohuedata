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

  const hoursToHighlight = ['8 PM', '9 PM', '10 PM', '11 PM', '12 AM', '1 AM', '2 AM', '3 AM', '4 AM', '5 AM', '6 AM', '7 AM'];
  const dayHighlighter = {
    id: 'dayHighlighter',
    beforeDatasetsDraw: (chart: ChartJS) => {
      const { ctx, data: { labels }, chartArea: { top, height }, scales: { x } } = chart;
      if (!labels) return;

      const stringLabels = labels as string[];
      const labelHours = stringLabels.map((label: string) => label.split(', ')[1]);
      const startLabel = stringLabels.find((_, i) => hoursToHighlight.includes(labelHours[i]));
      if (!startLabel) return;

      const startLabelIndex = stringLabels.indexOf(startLabel);
      const endLabel = stringLabels.find((_, i) => i > startLabelIndex && !hoursToHighlight.includes(labelHours[i]));
      if (!endLabel) return;

      const endLabelIndex = stringLabels.indexOf(endLabel);
      const startLabelX = x.getPixelForValue(startLabelIndex);
      const endLabelX = x.getPixelForValue(endLabelIndex);
      const highlightWidth = endLabelX - startLabelX;

      ctx.fillStyle = 'rgba(53, 162, 235, 0.1)';
      ctx.fillRect(startLabelX, top, highlightWidth, height);
    }
  }

  return <Box display="flex" justifyContent="center" mb={2} height="450px">
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
    }} plugins={[dayHighlighter]} options={{
      responsive: true,
      plugins: { legend: { display: false } },
      scales: {
        x: {
          ticks: { autoSkip: true },
          title: {
            display: true,
            text: 'Time',
          },
        },
        y: {
          title: {
            text: 'Temperature',
            display: true,
          },
          ticks: {
            callback: value => {
              if (!Number.isInteger(value)) return null;
              return `${value}°${units}`;
            }
          }
        }
      },
    }} />
  </Box>;
};

export default TemperatureReadingGraph;
