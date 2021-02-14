import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader, 
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import moment from 'moment';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';

import { EntPatient,EntDentalappointment,EntPersonnel } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const check = {
    personnelcheck : true
  }
export default function SearchTableDentalappointment() {
  const classes = useStyles();
  const profile = { givenName: ''};
  const api = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [dentalappointment, setDentalappointment] = useState<EntDentalappointment[]>([]);

  const [patientName, setPatient] = useState(String);



  const SearchDentalappointment = async () => {
    const res = await api.listDentalappointment({ offset: 0 });
    const search = DentalappointmentSearch(res);
    setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setDentalappointment([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setDentalappointment(search);
                }
            })
        }

        setStatus(true);
  }

  const DentalappointmentSearch = (res: any) => {
    
    const data = res.filter((filter: EntDentalappointment) => filter?.edges?.patient?.name?.includes(patientName))
    if (data.length != 0 && patientName != "") {
        return data;
    }
    else{
      return data;
        }
    }


  const SearchhandleChange = (event: any) => {
    setPatient(event.target.value as string);
  };


  return (
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'Search Dentalappointment '}`}
      subtitle=""
     ></Header>
      <Content>
        <ContentHeader title="ค้นหาบันทึกการนัดทำทันตกรรม">
        {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity ={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
          <div>
            <Link component={RouterLink} to="/Dentalappointment">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              Back
            </Button>
          </Link>
          </div>
          
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>ค้นหาจากชื่อผู้ป่วย</div>
                <TextField id="searchnote" type='string' 
                  onChange={SearchhandleChange}
                  value={patientName}
                />
              </FormControl>
            </div>
          

                        
            <div className={classes.margin}>
            <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  SearchDentalappointment();
                }}
                variant="contained"
                color="primary"
              >
                Search
             </Button>
              </Typography>
            </div>
          </form>
        </div>

        <TableContainer component={Paper}>
          <Table className={classes.withoutLabel} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">อันดับที่</TableCell>
                <TableCell align="center">ผู้ป่วยที่มารับการตรวจ</TableCell>
                <TableCell align="center">ประเภทการรักษา</TableCell>
                <TableCell align="center">ผู้บันทึกข้อมูล</TableCell>
                <TableCell align="center">วัน-เวลา</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {dentalappointment.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                  <TableCell align="center">{item.edges?.dentalkind?.kindname}</TableCell>
                  <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                  <TableCell align="center">{moment(item.appointtime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>

      </Content>
    </Page>
  );
}