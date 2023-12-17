import styled from 'styled-components';
import { SegmentedControl, themeGet } from '@primer/react';

const HeaderSegmentedControl = styled(SegmentedControl).attrs({
})`
  background-color: transparent;
  border-color: ${themeGet('colors.headerSearch.border')};
`;

export default HeaderSegmentedControl;
