import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Typography,
  Grid,
  List,
  ListItem,
  ListItemText,
  Link,
} from '@material-ui/core';
import {
  Content,
  InfoCard,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SupportButton,
} from '@backstage/core';
import Button from '@material-ui/core/Button';

const WelcomePage: FC<{}> = () => {
  const profile = { givenName: '' };

  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`ระบบบันทึกการนัดหมาย`}
      >
      </Header>
      <Content>
      <ContentHeader title="ตัวเลือกการนัดหมาย">  
       </ContentHeader>
       <tr>
         <td>
       <Link component={RouterLink} to="/Checksymptom">
         <Button variant="contained" color="primary">
             ตรวจอาการ
           </Button>
           </Link>
           </td>
      </tr>
      <tr>
       <Link component={RouterLink} to="/Dentalappointment">
         <Button variant="contained" color="primary">
             ทำฟัน
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Antenatal">
         <Button variant="contained" color="primary">
             ฝากครรภ์
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Createphysicaltherapyrecord">
         <Button variant="contained" color="primary">
             กายภาพบำบัด
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Bonedisease">
         <Button variant="contained" color="primary">
             โรคกระดูก
           </Button>
           </Link>
      </tr> 
      <tr>
       <Link component={RouterLink} to="/Surgeryappointment">
         <Button variant="contained" color="primary">
             ผ่าตัด
           </Button>
           </Link>
      </tr> 
  
      </Content>
    </Page>
  );
};

export default WelcomePage;
