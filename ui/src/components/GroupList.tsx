import React from 'react';
import Group from '../models/Group';
import GroupListItem from './GroupListItem';

interface Props {
  groups: Group[];
}

const GroupList = ({ groups }: Props) => <ul>
  {groups.map(group => <GroupListItem key={group.id} group={group} />)}
</ul>;

export default GroupList;
