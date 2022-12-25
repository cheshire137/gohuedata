import React, { useContext } from 'react';
import { Header, Heading, PageLayout } from '@primer/react';
import { PageContext } from '../contexts/PageContext';
import { useHref, Outlet } from 'react-router-dom';

const AppLayout = () => {
  const { pageTitle } = useContext(PageContext);

  return <PageLayout>
    <PageLayout.Header>
      <Header>
        <Header.Item>
          <Heading as="h1">
            <Header.Link href={useHref('/')}>gohuedata</Header.Link>
          </Heading>
          {pageTitle && pageTitle.length > 0 && <Heading
            as="h2"
            sx={{ fontWeight: 'normal', fontSize: 3, mx: 4 }}
          >{pageTitle}</Heading>}
        </Header.Item>
      </Header>
    </PageLayout.Header>
    <PageLayout.Content sx={{ fontSize: 2 }}>
      <Outlet />
    </PageLayout.Content>
    <PageLayout.Footer divider="line"></PageLayout.Footer>
  </PageLayout>;
};

export default AppLayout;
