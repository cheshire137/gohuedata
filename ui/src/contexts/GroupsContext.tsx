import React, { createContext, PropsWithChildren } from 'react';
import useGetGroups from '../hooks/use-get-groups';
import Group from '../models/Group';
import PageHeader from '../components/PageHeader';
import { Flash, PageLayout } from '@primer/react';

export type GroupsContextProps = {
  groups: Group[];
  totalCount: number;
};

export const GroupsContext = createContext<GroupsContextProps>({
  groups: [],
  totalCount: 0,
});

export const GroupsContextProvider = ({ children }: PropsWithChildren) => {
  const { groups, totalCount, fetching, error } = useGetGroups();

  if (fetching) {
    return <PageLayout>
      <PageHeader />
      <PageLayout.Content>
        <p>Loading groups...</p>
      </PageLayout.Content>
    </PageLayout>;
  }

  if (error) {
    return <PageLayout>
      <PageHeader />
      <PageLayout.Content>
        <Flash variant="danger">Error loading groups: {error}</Flash>
      </PageLayout.Content>
    </PageLayout>;
  }

  return <GroupsContext.Provider value={{
    groups: groups!,
    totalCount: totalCount!,
  }}>{children}</GroupsContext.Provider>;
};
