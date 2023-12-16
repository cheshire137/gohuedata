import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';
import { useLoaderData } from 'react-router-dom';
import GroupExtended from '../models/GroupExtended';

const GroupPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const group = useLoaderData() as GroupExtended;

  useEffect(() => setPageTitle(`Group ${group.name}`), [setPageTitle, group.name]);

  return <p></p>;
};

export default GroupPage;
