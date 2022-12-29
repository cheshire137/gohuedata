import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';
import useGetGroups from '../hooks/use-get-groups';
import { Flash } from '@primer/react';
import GroupList from './GroupList';

const GroupsPage = () => {
  const { setPageTitle } = useContext(PageContext);
  const { groups, fetching, error } = useGetGroups();

  useEffect(() => setPageTitle('Groups'), [setPageTitle]);

  if (fetching) {
    return <p>Loading groups...</p>;
  }

  if (error) {
    return <Flash variant="danger">Error loading groups: {error}</Flash>;
  }

  return <GroupList groups={groups!} />;
};

export default GroupsPage;
