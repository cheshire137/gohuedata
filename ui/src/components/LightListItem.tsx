import React from 'react';
import { Box } from '@primer/react';
import LightExtended from '../models/LightExtended';

interface Props {
  light: LightExtended;
}

const LightListItem = ({ light }: Props) => {
  console.log('light', light)
  return <Box as="li" mb={2}>
    {light.name}
  </Box>;
};

export default LightListItem;
