import { createPlugin } from '@backstage/core';
import Bonedisease from './components/Bonedisease';
import LoginPage from './components/LoginPage';
import WelcomePage from './components/WelcomePage';
import Dentalappointment from './components/Dentalappointment';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Bonedisease', Bonedisease);
    router.registerRoute('/Dentalappointment', Dentalappointment);
  },
});
