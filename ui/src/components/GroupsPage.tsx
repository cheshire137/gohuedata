import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';
import GroupList from './GroupList';

const GroupsPage = () => {
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle('Groups'), [setPageTitle]);

  return <GroupList />;
};

export default GroupsPage;
