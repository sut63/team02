import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Grid,
  Link,
} from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import Button from '@material-ui/core/Button';


const WelcomePage: FC<{}> = () => {

  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`ระบบบันทึกการนัดหมาย`}
      >
      </Header>
      <Content>
      <ContentHeader title="ตัวเลือกการนัดหมาย">
      <Link component={RouterLink} to="/">
         <Button variant="contained" color="secondary" size="large" >
             ออกจากระบบ
           </Button>
           </Link>
       </ContentHeader>
      
       <tr>
         <td>
       <Link component={RouterLink} to="/Checksymptom">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
         ระบบการนัดตรวจอาการ
           </Button>
           </Link>
           </td>
      </tr>
      <Grid item xs={3}>
      <tr>
       <Link component={RouterLink} to="/Dentalappointment">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
             ระบบการนัดทำทันตกรรม
           </Button>
           </Link>
      </tr> 
      </Grid>
      <tr>
       <Link component={RouterLink} to="/Antenatal">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
         ระบบการนัดฝากครรภ์
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Createphysicaltherapyrecord">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
         ระบบการนัดกายภาพบำบัด
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Bonedisease">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
         ระบบการนัดโรคกระดูก
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Surgeryappointment">
         <Button variant="contained" color="primary" size="large" style={{ width: 300 , margin: 10}}>
         ระบบการนัดผ่าตัด
           </Button>
           </Link>
      </tr> 
  
      </Content>
    </Page>
  );
};

export default WelcomePage;
