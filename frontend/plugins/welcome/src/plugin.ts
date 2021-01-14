import { createPlugin } from '@backstage/core';
import Bonedisease from './components/Bonedisease';
import LoginPage from './components/LoginPage';
import WelcomePage from './components/WelcomePage';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Bonedisease', Bonedisease);
  },
});
