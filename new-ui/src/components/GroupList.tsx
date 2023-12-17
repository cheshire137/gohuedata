import React, { useEffect, useContext } from 'react';
import { GroupsContext } from '../contexts/GroupsContext';
import { PageContext } from '../contexts/PageContext';
import GroupListItem from './GroupListItem';

const GroupList = () => {
  const { groups, totalCount } = useContext(GroupsContext);
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle(`Groups (${totalCount})`), [setPageTitle, totalCount]);

  return <ul>
    {groups.map(group => <GroupListItem key={group.uniqueID} group={group} />)}
  </ul>;
};

export default GroupList;
