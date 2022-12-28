import React, { useContext, useState, useEffect } from 'react';
import { Box } from '@primer/react';
import { TemperatureReadingsContext } from '../contexts/TemperatureReadingsContext';
import { SettingsContext } from '../contexts/SettingsContext';
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

const thermScaleFor = (units: string) => units === 'F' ? 'Fahrenheit' : 'Celsius';

const TemperatureReadingGraph = () => {
  const { temperatureReadings: rawTempReadings } = useContext(TemperatureReadingsContext);
  const { units } = useContext(SettingsContext);
  const [thermScale, setThermScale] = useState(thermScaleFor(units));
  const [normalizedTempReadings, setNormalizedTempReadings] = useState(rawTempReadings);
  const [labels, setLabels] = useState(rawTempReadings.map(tempReading => tempReading.timestamp));

  useEffect(() => setThermScale(thermScaleFor(units)), [units, setThermScale]);

  useEffect(() => {
    const newNormalizedReadings = [...rawTempReadings].filter(reading => reading.timestampAsDate())
      .sort((a, b) => a.timestamp.localeCompare(b.timestamp));
    setNormalizedTempReadings(newNormalizedReadings);
  }, [rawTempReadings]);

  useEffect(() => {
    if (normalizedTempReadings.length < 1) return;

    let dayIndex = 0;
    let day = normalizedTempReadings[dayIndex].timestampAsDate()!.toDateString();
    const newLabels = normalizedTempReadings.map((tempReading, i) => {
      const date = tempReading.timestampAsDate()!;
      if (date.toDateString() === day && dayIndex !== i) {
        return date.toLocaleTimeString('en-US', { hour: 'numeric' });
      }
      dayIndex = i;
      day = date.toDateString();
      return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', hour: 'numeric' });
    });
    setLabels(newLabels);
  }, [normalizedTempReadings, setLabels]);

  const dayDivider = {
    id: 'dayDivider',
    beforeDatasetsDraw: (chart: ChartJS) => {
      const { ctx, data: { labels }, chartArea: { top, height }, scales: { x } } = chart;
      if (!labels) return;

      (labels as string[]).forEach((label, i) => {
        if (!label.includes(', ')) return; // 'Dec 28, 11 AM' (start of a new day) vs '11 AM' (continuing same day)
        if (i === 0) return; // Don't draw a line on the left edge of the graph
        ctx.fillStyle = 'rgba(0, 0, 0, 0.1)';
        ctx.fillRect(x.getPixelForValue(i) - 2, top, 4, height);
      });
    }
  }

  return <Box display="flex" justifyContent="center" mb={2} height="450px">
    <Line data={{
      labels,
      datasets: [
        {
          label: `Temperature in ${thermScale}`,
          data: normalizedTempReadings.map(tempReading => tempReading.temperature),
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
