import React, { useContext, useEffect } from 'react';
import { Box, CounterLabel, Pagehead } from '@primer/react';
import { PageContext } from '../contexts/PageContext';
import { useLoaderData } from 'react-router-dom';
import GroupExtended from '../models/GroupExtended';
import LightList from './LightList';

const GroupPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const group = useLoaderData() as GroupExtended;

  useEffect(() => setPageTitle(`Group: ${group.name}`), [setPageTitle, group.name]);

  const onLights = group.lights.filter(light => light.latestState.on);
  const offLights = group.lights.filter(light => !light.latestState.on);

  return <div>
    <div>
      Type: <Box fontSize={1} display="inline-block" color="fg.muted">{group.type}</Box> / Class: <Box fontSize={1} display="inline-block" color="fg.muted">{group.groupClass}</Box> / Bridge: <Box fontSize={1} display="inline-block" color="fg.muted">{group.bridge.name}</Box>
    </div>
    <dl>
      <dt># lights</dt>
      <dd>{group.totalLights}</dd>
      <dt># sensors</dt>
      <dd>{group.totalSensors}</dd>
    </dl>
    <Pagehead as="h2" sx={{ fontSize: 3, display: 'flex', alignItems: 'baseline' }}>
      On
      <CounterLabel sx={{ ml: 2, fontSize: 2 }}>{onLights.length}</CounterLabel>
    </Pagehead>
    <LightList lights={onLights} />
    <Pagehead as="h2" sx={{ fontSize: 3, display: 'flex', alignItems: 'baseline' }}>
      Off
      <CounterLabel sx={{ ml: 2, fontSize: 2 }}>{offLights.length}</CounterLabel>
    </Pagehead>
    <LightList lights={offLights} />
  </div>;
};

export default GroupPage;
