import styled from 'styled-components';
import { SegmentedControl, themeGet } from '@primer/react';

const HeaderSegmentedControlButton = styled(SegmentedControl.Button).attrs({
})`
  color: ${props => props.selected ? themeGet('colors.fg.default') : themeGet('colors.header.text')};
`;

export default HeaderSegmentedControlButton;
