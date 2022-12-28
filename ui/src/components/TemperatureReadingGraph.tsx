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

  const dayDivider = {
    id: 'dayDivider',
    beforeDatasetsDraw: (chart: ChartJS) => {
      const { ctx, data: { labels }, chartArea: { top, height }, scales: { x } } = chart;
      if (!labels) return;

      const stringLabels = labels as string[];
      const dayLabels = stringLabels.map((label: string) => label.split(', ')[0]);
      if (dayLabels.length < 1) return;

      let dayIndex = 0;
      let day = dayLabels[dayIndex];
      for (let nextDayIndex=0; nextDayIndex < dayLabels.length; nextDayIndex++) {
        const nextDay = dayLabels[nextDayIndex];
        if (day === nextDay) continue;

        ctx.fillStyle = 'rgba(0, 0, 0, 0.1)';
        ctx.fillRect(x.getPixelForValue(nextDayIndex)-2, top, 4, height);

        day = nextDay;
      }
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
    }} plugins={[dayDivider]} options={{
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
