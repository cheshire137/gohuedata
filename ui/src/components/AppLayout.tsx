import React from 'react';
import { PageLayout } from '@primer/react';
import { useHref, Outlet } from 'react-router-dom';
import PageHeader from './PageHeader';

const AppLayout = () => <PageLayout>
  <PageHeader homeUrl={useHref('/')} />
  <PageLayout.Content sx={{ fontSize: 2 }}>
    <Outlet />
  </PageLayout.Content>
</PageLayout>;

export default AppLayout;
