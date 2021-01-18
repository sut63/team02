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
//import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
//import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';

import { EntPatient,EntDentalkind,EntPersonnel } from '../../api';

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
export default function createDentalappointment() {
  const classes = useStyles();
  const profile = { givenName: ''};
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [kindnames, setKindnames] = useState<EntDentalkind[]>([]);
  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);

  const [loading, setLoading] = useState(true);

  const [patientName, setPatient] = useState(Number);
  const [kindName, setKindname] = useState(Number);
  const [personnelName, setPersonnel] = useState(Number);
  const [datetime, setDateTime] = useState(String);

  useEffect(() => {

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
      console.log(res);
    };
    getPatients();

    const getDentalkinds = async () => {
      const res = await api.listDentalkind({ limit: 10, offset: 0 });
      setLoading(false);
      setKindnames(res);
    };
    getDentalkinds();

    const getPersonnels = async () => {
      const res = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnels(res);
    };
    getPersonnels();
    setPersonnel(Number(localStorage.getItem("personaldata")))
  }, [loading]);

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const KindnamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setKindname(event.target.value as number);
  };

  const PersonnelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnel(event.target.value as number);
  };
  const DateTimehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDateTime(event.target.value as string);
  };

  const listDapm = () => {
    window.location.href ="http://localhost:3000/Dentalappointment";
  };


  const createDentalappointment = async () => {
    if((patientName!=0)&&(kindName!=0)&&(personnelName!=0)&&(datetime!="")){
    const dentalappointment = {
      patientID: patientName,
      kindName: kindName,
      personnelID: personnelName,
      appointTime: datetime + ":00+07:00",

    };
    console.log(dentalappointment);
    const res: any =  await api.createDentalappointment({ dentalappointment : dentalappointment});
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      }
    } 
    else {
      setAlert(false);
      setStatus(true);
    }
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'Dentalappointment '}`}
      subtitle=""
     ></Header>
      <Content>
        <ContentHeader title="บันทึกนัดการทำทันตกรรม">
          <div>
            <Link component={RouterLink} to="/WelcomePage">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              Back
            </Button>
          </Link>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => {window.location.reload(false)}} style={{ marginTop: 20 }}>
                  success!
                </Alert>
              ) : (
                  <Alert severity="warning" onClose={() => {window.location.reload(false)}} style={{ marginTop: 20 }}>
                    This is a warning alert — Please enter all information!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div><strong>ชื่อทันตแพทย์</strong></div>
                <TextField
                                    id="user"
                                    type="string"
                                    size="medium"
                                    value={personnels.filter((filter:EntPersonnel) => filter.id == personnelName).map((item:EntPersonnel) => `${item.name}`)}
                                    style={{ width: 400 }}
                                />
              </FormControl>
            </div>

            
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography variant="h6" gutterBottom  align="center">
                Patient Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                  labelId="patient"
                  id="patient"
                  value={patientName}
                  onChange={PatienthandleChange}
                  style={{ width: 400 }}
                >
               {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography variant="h6" gutterBottom  align="center">
                Dental Kind Name : 
                <Typography variant="body1" gutterBottom> 
              <Select
                labelId="dentalkind"
                id="dentalkind"
                value={kindName}
                onChange={KindnamehandleChange}
                style={{ width: 400 }}
              > 
                {kindnames.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.kindname}
                      </MenuItem>
                    );
                  })}
              </Select>
                </Typography>
                </Typography>
            </FormControl>

            
            <tr>
            <td>
            <FormControl
                  fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <form className={classes.root} noValidate>
                      <TextField
                        id="datetime-local"                                                            
                        label="Date"
                        type="datetime-local"
                        //defaultValue="2017-05-24T10:30"
                        className={classes.textField}
                        onChange={DateTimehandleChange}
                        InputLabelProps={{
                          shrink: true,
                        }}
                      />
                </form>
               </FormControl>
            </td>
          </tr>

                        
            <div className={classes.margin}>
            <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  createDentalappointment();
                }}
                variant="contained"
                color="primary"
              >
                Create
             </Button>
              </Typography>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}