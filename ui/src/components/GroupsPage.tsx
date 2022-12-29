import React, { useContext, useEffect } from 'react';
import { PageContext } from '../contexts/PageContext';

const GroupsPage = () => {
  const { setPageTitle } = useContext(PageContext);

  useEffect(() => setPageTitle('Groups'), [setPageTitle]);

  return <p></p>;
};

export default GroupsPage;
