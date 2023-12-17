import React, { ReactNode } from 'react';
import { CircleBadge, Text } from '@primer/react';

interface Props {
  temperature: number;
  units: string;
  children: ReactNode;
}

const TemperatureBadge = ({ temperature, units, children }: Props) => <CircleBadge variant="large" sx={{ flexDirection: 'column' }}>
  <Text
    fontWeight="bold"
    fontSize="5"
  >{Math.round(temperature)}&deg;{units}</Text>
  <Text
    fontSize={1}
    color="fg.muted"
  >{children}</Text>
</CircleBadge>;

export default TemperatureBadge;
