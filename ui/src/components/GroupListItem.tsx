import React from 'react';
import { Box, Link } from '@primer/react';
import Group from '../models/Group';
import { useHref } from 'react-router-dom';

interface Props {
  group: Group;
}

const GroupListItem = ({ group }: Props) => {
  return <Box as="li" mb={2}>
    <Link href={useHref(`/group/${group.id}/?bridge=${encodeURIComponent(group.bridge.name)}`)}>
      {group.name} &mdash; {group.totalLights} {group.totalLights === 1 ? 'light' : 'lights'}
    </Link>
    <Box fontSize={1} display="inline-block" color="fg.muted">
      {group.bridge.name}
    </Box>
  </Box>;
};

export default GroupListItem;
