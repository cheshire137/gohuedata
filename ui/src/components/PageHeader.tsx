import React, { useContext } from 'react';
import { Header, Heading, PageLayout } from '@primer/react';
import { PageContext } from '../contexts/PageContext';

interface Props {
  homeUrl?: string;
}

const PageHeader = ({ homeUrl }: Props) => {
  const { pageTitle } = useContext(PageContext);

  return <PageLayout.Header>
    <Header>
      <Header.Item>
        <Heading as="h1">
          {homeUrl ? <Header.Link href={homeUrl}>gohuedata</Header.Link> : 'gohuedata'}
        </Heading>
        {pageTitle && pageTitle.length > 0 && <Heading
          as="h2"
          sx={{ fontWeight: 'normal', fontSize: 3, mx: 4 }}
        >{pageTitle}</Heading>}
      </Header.Item>
    </Header>
  </PageLayout.Header>;
};

export default PageHeader;
