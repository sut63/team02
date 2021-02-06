import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import { Link, Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';

import moment from 'moment';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

import { EntBonedisease } from '../../api/models/EntBonedisease';
import { EntPersonnel } from '../../api/models/EntPersonnel';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },

    withoutLabel: {
      marginTop: theme.spacing(1),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 100,
    },
  }),
);

const check = {
  personnelcheck : true
}

export default function Searchtable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [bonedisease, setBonedisease] = useState<EntBonedisease[]>([]);
  const [loading, setLoading] = useState(true);
  const profile = { givenName: 'ระบบบันทึกนัดหมายการตรวจโรคกระดูก' };
  const [status, setStatus] = useState(false);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [personnel, setPersonnel] = useState<EntPersonnel[]>([]);
  const [bonediseasesearch, setBonediseaseSearch] = useState(String);

  

  useEffect(() => {
  
  }, [loading]);



  const SearchBonedisease = async () => {
    const res = await api.listBonedisease({ offset: 0 });
    const search = BonediseaseSearch(res);
    setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setBonedisease([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setBonedisease(search);
                }
            })
        }

        setStatus(true);
  }

  const BonediseaseSearch = (res: any) => {
    const data = res.filter((filter: EntBonedisease) => filter?.identificationCard?.includes(bonediseasesearch))
    if (data.length != 0 && bonediseasesearch != "") {
        return data;
    }
    else{
      return data;
        }
    }


  const handleSearchChange = (event: any) => {
    setBonediseaseSearch(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.website}>
     <Header
       title={`ท่านกำลังใช้งาน ${profile.givenName || ':)'}`}
     ><Link component={RouterLink} to="/Bonedisease">
     <Button variant="contained" color="primary" >
       ย้อนกลับ
     </Button>
   </Link>
   </Header>
     <Content>
        <ContentHeader title="ค้นหาบันทึกนัดหมายการนัดตรวจโรคกระดูก">
        {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
        </ContentHeader>
  
        <div className={classes.paper}><strong>ค้นหาผู้ป่วยที่เคยได้รับการตรวจโรคกระดูก (ด้วยเลข ปชช)</strong></div>
        <TextField className={classes.textField}
          style={{ width: 400, marginLeft: 20, marginRight: -10 }}
          id="customer"
          label=""
          variant="standard"
          color="secondary"
          type="string"
          size="medium"
          value={bonediseasesearch}
          onChange={handleSearchChange}
        />
        <div className={classes.margin}>
          <Button
            onClick={() => {
              SearchBonedisease();
            }}
            variant="contained"
            color="primary"
          >
            ค้นหา
          </Button>
        </div>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">อันดับที่</TableCell>
                <TableCell align="center">ผู้ป่วยที่มารับการตรวจ</TableCell>
                <TableCell align="center">ID Card</TableCell>
                <TableCell align="center">เลือกวิธีการรักษา</TableCell>
                <TableCell align="center">รายละเอียดเพิ่มเติม</TableCell>
                <TableCell align="center">ผู้บันทึกข้อมูล</TableCell>
                <TableCell align="center">วัน-เวลา</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bonedisease.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                  <TableCell align="center">{item.identificationCard}</TableCell>
                  <TableCell align="center">{item.edges?.remedy?.remedy}</TableCell>
                  <TableCell align="center">{item.advice}</TableCell>
                  <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                  <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>
      </Content>
    </Page>
  );
};
//issue