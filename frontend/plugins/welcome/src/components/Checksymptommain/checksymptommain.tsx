import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './../TableCheck';
import Button from '@material-ui/core/Button';
//import Typography from '@material-ui/core/Typography';
 
import {

 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
const profile = { givenName: 'ระบบการนัดตรวจอาการ' };

 return (
   <Page theme={pageTheme.library}>
     <Header
     
     
       title={`ท่านกำลังใช้งาน ${profile.givenName || ':)'}`}
       subtitle="ชมรมคนชอบนัดตรวจอาการ"
       
     >
       <Link component={RouterLink} to="/WelcomePage">
           <Button variant="contained" color="secondary" >
             หน้า Welcome
           </Button>
         </Link>
   </Header>
     <Content>
       <ContentHeader title="ตารางการนัดตรวจอาการ">
       
       <Link component={RouterLink} to="/checksymptom">
           <Button variant="contained" color="primary" >
             เพิ่มบันทึก
           </Button>
         </Link>
         <Link component={RouterLink} to="/Checksymptomsearch">
     <Button variant="contained" color="primary" >
       ค้นหา
     </Button>
   </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
       <br></br>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;