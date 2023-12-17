import React, { useContext, useEffect } from 'react';
import { Box, Heading } from '@primer/react';
import { PageContext } from '../contexts/PageContext';
import { useLoaderData } from 'react-router-dom';
import GroupExtended from '../models/GroupExtended';
import BridgeDisplay from './BridgeDisplay';
import LightList from './LightList';

const GroupPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const group = useLoaderData() as GroupExtended;

  useEffect(() => setPageTitle(`Group: ${group.name}`), [setPageTitle, group.name]);

  const onLights = group.lights.filter(light => light.latestState.on);
  const offLights = group.lights.filter(light => !light.latestState.on);

  return <div>
    <div>
      Type: <Box fontSize={1} display="inline-block" color="fg.muted">{group.type}</Box> / Class: <Box fontSize={1} display="inline-block" color="fg.muted">{group.groupClass}</Box> / Bridge: <BridgeDisplay bridge={group.bridge} />
    </div>
    <dl>
      <dt># lights</dt>
      <dd>{group.totalLights}</dd>
      <dt># sensors</dt>
      <dd>{group.totalSensors}</dd>
    </dl>
    <Heading as="h2">On</Heading>
    <LightList lights={onLights} />
    <Heading as="h2">Off</Heading>
    <LightList lights={offLights} />
  </div>;
};

export default GroupPage;
