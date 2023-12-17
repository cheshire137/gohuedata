import React, { createContext, PropsWithChildren, useState, useCallback, useEffect, useMemo } from 'react';

export type PageContextProps = {
  pageTitle: string;
  setPageTitle: (pageTitle: string) => void;
};

export const PageContext = createContext<PageContextProps>({
  pageTitle: '',
  setPageTitle: () => { },
});

export const PageContextProvider = ({ children }: PropsWithChildren) => {
  const [pageTitle, _setPageTitle] = useState('');
  const setPageTitle = useCallback((title: string) => _setPageTitle(title), [_setPageTitle]);
  const contextProps = useMemo(() => ({ pageTitle, setPageTitle }), [pageTitle, setPageTitle])

  useEffect(() => {
    if (pageTitle.length > 0) {
      document.title = `${pageTitle} - gohuedata`;
    } else {
      document.title = 'gohuedata';
    }
  }, [pageTitle])

  return <PageContext.Provider value={contextProps}>{children}</PageContext.Provider>;
};
