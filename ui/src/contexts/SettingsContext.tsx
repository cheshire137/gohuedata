import React, { useState, useEffect } from 'react';

export type SettingsContextProps = {
  setFahrenheit: (fahrenheit: boolean) => void;
  units: string;
  fahrenheit: boolean;
};

export const SettingsContext = React.createContext<SettingsContextProps>({
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

  useEffect(() => {
    if (typeof fahrenheit === 'boolean') {
      setUnits(fahrenheit ? 'F' : 'C');
    } else {
      setUnits('F');
    }
  }, [fahrenheit]);

  return <SettingsContext.Provider value={{
    units,
    fahrenheit: units === 'F',
    setFahrenheit: (fahrenheit) => setUnits(fahrenheit ? 'F' : 'C'),
  }}>{children}</SettingsContext.Provider>;
};
