import { createPlugin } from '@backstage/core';
import Bonedisease from './components/Bonedisease';
import LoginPage from './components/LoginPage';
import WelcomePage from './components/WelcomePage';
import Dentalappointment from './components/Dentalappointment';
import Surgeryappointment from './components/Surgeryappointment';
import Physicaltherapyrecord from './components/Physicaltherapyrecord';
import Checksymptom from './components/Checksymptoms';
import Antenatal from './components/Antenatal';



export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Bonedisease', Bonedisease);
    router.registerRoute('/Dentalappointment', Dentalappointment);
    router.registerRoute('/Surgeryappointment', Surgeryappointment);
    router.registerRoute('/Createphysicaltherapyrecord', Physicaltherapyrecord);
    router.registerRoute('/Checksymptom', Checksymptom);
    router.registerRoute('/Antenatal', Antenatal);
  },
});
