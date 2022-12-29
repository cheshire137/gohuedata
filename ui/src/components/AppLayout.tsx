import React from 'react';
import { PageLayout, NavList } from '@primer/react';
import { useHref, Outlet, useLocation } from 'react-router-dom';
import PageHeader from './PageHeader';

const AppLayout = () => {
  const { pathname } = useLocation();

  return <PageLayout>
    <PageHeader homeUrl={useHref('/')} />
    <PageLayout.Pane position="start">
      <NavList>
        <NavList.Item href={useHref('/')} aria-current={pathname === '/'}>Temperature sensors</NavList.Item>
      </NavList>
      <NavList>
        <NavList.Item href={useHref('/groups')} aria-current={pathname === '/groups'}>Groups</NavList.Item>
      </NavList>
    </PageLayout.Pane>
    <PageLayout.Content sx={{ fontSize: 2 }}>
      <Outlet />
    </PageLayout.Content>
  </PageLayout>;
};

export default AppLayout;
