import { createPlugin } from '@backstage/core';
import Bonedisease from './components/Bonedisease';
import BonediseaseTable from './components/BonediseaseTable';
import LoginPage from './components/LoginPage';
import WelcomePage from './components/WelcomePage';
import Dentalappointment from './components/Dentalappointment';
import Surgeryappointment from './components/Surgeryappointment';




export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Bonedisease', Bonedisease);
    router.registerRoute('/BonediseaseTable', BonediseaseTable);
    router.registerRoute('/Dentalappointment', Dentalappointment);
    router.registerRoute('/Surgeryappointment', Surgeryappointment);
  },
});
