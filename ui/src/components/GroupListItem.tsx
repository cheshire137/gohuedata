import React from 'react';
import { Box } from '@primer/react';
import Group from '../models/Group';

interface Props {
  group: Group;
}

const GroupListItem = ({ group }: Props) => {
  return <Box as="li" mb={2}>
    {group.name}
    <Box fontSize={1} color="fg.muted">
      {group.bridge.name}
    </Box>
  </Box>;
};

export default GroupListItem;
