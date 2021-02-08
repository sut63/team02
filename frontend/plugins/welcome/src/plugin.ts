import { createPlugin } from '@backstage/core';
import Bonedisease from './components/Bonedisease';
import BonediseaseSearch from './components/BonediseaseSearch';
import LoginPage from './components/LoginPage';
import WelcomePage from './components/WelcomePage';
import Dentalappointment from './components/Dentalappointment';
import Surgeryappointment from './components/Surgeryappointment';
import Physicaltherapyrecord from './components/Physicaltherapyrecord';
import Checksymptom from './components/Checksymptoms';
import Antenatal from './components/Antenatal';
import SearchDentalappointment from './components/SearchDentalappointment';
import SearchSurgeryappointment from './components/SearchSurgeryappointment';
import Searchtable from './components/SearchChecksymptom';
import checksymptommain from './components/Checksymptommain';
import ComponentsTable from './components/TableCheck';
import Searchphysicaltherapyrecord from './components/Searchphysicaltherapyrecord';
import SearchAntenatalinformation from './components/SearchAntenatalinformation';
 

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', LoginPage);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Bonedisease', Bonedisease);
    router.registerRoute('/BonediseaseSearch', BonediseaseSearch);
    router.registerRoute('/Dentalappointment', Dentalappointment);
    router.registerRoute('/Surgeryappointment', Surgeryappointment);
    router.registerRoute('/Createphysicaltherapyrecord', Physicaltherapyrecord);
    router.registerRoute('/Checksymptom', Checksymptom);
    router.registerRoute('/Checksymptomsearch',Searchtable);
    router.registerRoute('/Tables',ComponentsTable);
    router.registerRoute('/Checksymptommain',checksymptommain);
    router.registerRoute('/Antenatal', Antenatal);
    router.registerRoute('/Searchdentalappointment', SearchDentalappointment);
    router.registerRoute('/SearchSurgeryappointment', SearchSurgeryappointment);
    router.registerRoute('/Searchphysicaltherapyrecord', Searchphysicaltherapyrecord);
    router.registerRoute('/SearchAntenatalinformation', SearchAntenatalinformation);
  
  },
});
