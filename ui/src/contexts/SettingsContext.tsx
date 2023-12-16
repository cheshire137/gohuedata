import React, { createContext, useMemo, useState, useEffect } from 'react';

export type SettingsContextProps = {
  setFahrenheit: (fahrenheit: boolean) => void;
  units: string;
  fahrenheit: boolean;
};

export const SettingsContext = createContext<SettingsContextProps>({
  setFahrenheit: () => { },
  units: 'F',
  fahrenheit: true,
});

interface Props {
  fahrenheit?: boolean;
  children: React.ReactNode;
}

export const SettingsContextProvider = ({ fahrenheit, children }: Props) => {
  const [units, setUnits] = useState('F');
  const contextProps = useMemo(() => ({
    units,
    fahrenheit: units === 'F',
    setFahrenheit: (fahrenheit) => setUnits(fahrenheit ? 'F' : 'C'),
  } satisfies SettingsContextProps), [units]);

  useEffect(() => {
    if (typeof fahrenheit === 'boolean') {
      setUnits(fahrenheit ? 'F' : 'C');
    } else {
      setUnits('F');
    }
  }, [fahrenheit]);

  return <SettingsContext.Provider value={contextProps}>{children}</SettingsContext.Provider>;
};
