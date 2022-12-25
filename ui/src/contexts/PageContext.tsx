import React, { useState, useCallback, useEffect } from 'react';

export type PageContextProps = {
  pageTitle: string;
  setPageTitle: (pageTitle: string) => void;
};

export const PageContext = React.createContext<PageContextProps>({
  pageTitle: '',
  setPageTitle: () => { },
});

interface Props {
  children: React.ReactNode;
}

export const PageContextProvider = ({ children }: Props) => {
  const [pageTitle, _setPageTitle] = useState('');
  const setPageTitle = useCallback((title: string) => _setPageTitle(title), [_setPageTitle]);

  useEffect(() => {
    if (pageTitle.length > 0) {
      document.title = `gohuedata - ${pageTitle}`;
    } else {
      document.title = 'gohuedata';
    }
  }, [pageTitle])

  return <PageContext.Provider value={{ pageTitle, setPageTitle }}>
    {children}
  </PageContext.Provider>;
};
