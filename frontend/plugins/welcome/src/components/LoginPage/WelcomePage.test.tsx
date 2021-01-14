import React from 'react';
import { render } from '@testing-library/react';
import LoginPage from './LoginPage';
import { ThemeProvider } from '@material-ui/core';
import { lightTheme } from '@backstage/theme';

describe('LoginPage', () => {
  it('should render', () => {
    const rendered = render(
      <ThemeProvider theme={lightTheme}>
        <LoginPage />
      </ThemeProvider>,
    );
    expect(rendered.baseElement).toBeInTheDocument();
  });
});
