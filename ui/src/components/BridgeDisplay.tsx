import React from 'react';
import { Box } from '@primer/react';
import HueBridge from '../models/HueBridge';

interface Props {
  bridge: HueBridge;
}

const BridgeDisplay = ({ bridge }: Props) => <Box fontSize={1} display="inline-block" color="fg.muted">
  {bridge.name}
</Box>;

export default BridgeDisplay;
