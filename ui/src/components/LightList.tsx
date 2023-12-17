import React from 'react';
import LightExtended from '../models/LightExtended';
import LightListItem from './LightListItem';

interface Props {
  lights: LightExtended[];
}

const LightList = ({ lights }: Props) => {
  if (lights.length < 1) return <p>No lights</p>
  return <ol>
    {lights.map(light => <LightListItem key={light.uniqueID} light={light} />)}
  </ol>;
};

export default LightList;
