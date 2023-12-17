import React, { useContext } from 'react';
import { Header, Heading, PageLayout } from '@primer/react';
import { PageContext } from '../contexts/PageContext';
import { SettingsContext } from '../contexts/SettingsContext';
import HeaderSegmentedControlButton from './HeaderSegmentedControlButton';
import HeaderSegmentedControl from './HeaderSegmentedControl';

interface Props {
  homeUrl?: string;
}

const PageHeader = ({ homeUrl }: Props) => {
  const { pageTitle } = useContext(PageContext);
  const { setFahrenheit, fahrenheit } = useContext(SettingsContext);

  return <PageLayout.Header>
    <Header>
      <Header.Item full>
        <Heading as="h1">
          {homeUrl ? <Header.Link href={homeUrl}>gohuedata</Header.Link> : 'gohuedata'}
        </Heading>
        {pageTitle && pageTitle.length > 0 && <Heading
          as="h2"
          sx={{ fontWeight: 'normal', fontSize: 3, mx: 4 }}
        >{pageTitle}</Heading>}
      </Header.Item>
      <Header.Item>
        <HeaderSegmentedControl
          aria-label="Temperature display"
          onChange={selectedIndex => setFahrenheit(selectedIndex === 0)}
          sx={{ backgroundColor: 'transparent' }}
        >
          <HeaderSegmentedControlButton selected={fahrenheit}>Fahrenheit</HeaderSegmentedControlButton>
          <HeaderSegmentedControlButton selected={!fahrenheit}>Celsius</HeaderSegmentedControlButton>
        </HeaderSegmentedControl>
      </Header.Item>
    </Header>
  </PageLayout.Header>;
};

export default PageHeader;
