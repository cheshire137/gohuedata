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

const morningHours = ['6 AM', '7 AM', '8 AM', '9 AM', '10 AM', '11 AM'];
const afternoonHours = ['12 PM', '1 PM', '2 PM', '3 PM', '4 PM'];
const eveningHours = ['5 PM', '6 PM', '7 PM', '8 PM'];
const nightHours = ['9 PM', '10 PM', '11 PM', '12 AM', '1 AM', '2 AM', '3 AM', '4 AM', '5 AM'];

const morningColor = 'rgba(255,249,200)';
const afternoonColor = 'rgba(255,230,200)';
const eveningColor = 'rgba(234,200,255)';
const nightColor = 'rgba(217,217,217)';

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

  const dayHighlighter = {
    id: 'dayHighlighter',
    beforeDatasetsDraw: (chart: ChartJS) => {
      const { ctx, data: { labels }, chartArea: { top, height }, scales: { x } } = chart;
      if (!labels) return;

      const stringLabels = labels as string[];
      const labelHours = stringLabels.map((label: string) => label.split(', ')[1]);

      for (let startLabelIndex=0; startLabelIndex<labelHours.length; startLabelIndex++) {
        const startLabelHour = labelHours[startLabelIndex];
        if (morningHours.includes(startLabelHour)) {
          ctx.fillStyle = morningColor;
        } else if (afternoonHours.includes(startLabelHour)) {
          ctx.fillStyle = afternoonColor;
        } else if (eveningHours.includes(startLabelHour)) {
          ctx.fillStyle = eveningColor;
        } else if (nightHours.includes(startLabelHour)) {
          ctx.fillStyle = nightColor;
        } else {
          continue;
        }
        const startLabelX = x.getPixelForValue(startLabelIndex);
        const endLabelIndex = Math.min(startLabelIndex + 1, labelHours.length);
        const endLabelX = x.getPixelForValue(endLabelIndex);
        const highlightWidth = Math.min(endLabelX - startLabelX, x.width - startLabelX);
        ctx.fillRect(startLabelX, top, highlightWidth, height);
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
