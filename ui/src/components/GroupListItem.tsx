import React from 'react';
import { Box } from '@primer/react';
import Group from '../models/Group';
import BridgeDisplay from './BridgeDisplay';

interface Props {
  group: Group;
}

const GroupListItem = ({ group }: Props) => {
  return <Box as="li" mb={2}>
    {group.name}
    <BridgeDisplay bridge={group.bridge} />
  </Box>;
};

export default GroupListItem;
