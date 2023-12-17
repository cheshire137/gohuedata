import React from 'react';
import { LightBulbIcon } from '@primer/octicons-react';
import { Box, Octicon } from '@primer/react';
import LightExtended from '../models/LightExtended';

interface Props {
  light: LightExtended;
}

const LightListItem = ({ light }: Props) => <Box as="li" mb={2}>
  {light.name}
  <Octicon icon={LightBulbIcon} sx={{ ml: 2 }} />
  <Box sx={{ display: 'inline-block', ml: 2, fontSize: 1, color: 'fg.muted' }}>{light.latestState.timestamp}</Box>
</Box>;

export default LightListItem;
