import React, { useState, useEffect } from 'react';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Alert } from '@material-ui/lab';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,

  Button,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntStatus } from '../../api/models/EntStatus';
import { EntPhysicaltherapyroom } from '../../api/models/EntPhysicaltherapyroom';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(1),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  datetimepicker:{
    width: 300,
  },
}));
 
const HeaderCustom = {
  minHeight: '70px',
};

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  
  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [physicaltherapyrooms, setPhysicaltherapyrooms] = useState<EntPhysicaltherapyroom[]>([]);
  const [statuss, setStatuss] = useState<EntStatus[]>([]); 
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [physicaltherapyroomid, setPhysicaltherapyroomid] = useState(Number);
  const [statusid, setStatusid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
 

// 
useEffect(() => {

  const getPersonnels = async () => {

    const personnels = await api.listPersonnel({ limit: 10, offset: 0 });
    setLoading(false);
    setPersonnels(personnels);
  };
  getPersonnels();

  const getPhysicaltherapyrooms = async () => {
    const phr = await api.listPhysicaltherapyroom({ limit: 10, offset: 0 });
    setLoading(false);
    setPhysicaltherapyrooms(phr);
  };
  getPhysicaltherapyrooms();

  const getPatients = async () => {
   const pa = await api.listPatient({ limit: 10, offset: 0 });
   setLoading(false);
   setPatients(pa);
 };
 getPatients();

 const getStatuss = async () => {
   const st = await api.listStatus({ limit: 10, offset: 0 });
   setLoading(false);
   setStatuss(st);
 };
 getStatuss();

}, [loading]);
 
const createPhysicaltherapyrecord = async () =>{
  const physicaltherapyrecord = {
    time      : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
    personnel : personnelid,
    patient : patientid,
    physicaltherapyroom : physicaltherapyroomid,
    status  : statusid,
  }
  console.log(physicaltherapyrecord);

    const res:any = await api.createPhysicaltherapyrecord({ physicaltherapyrecord : physicaltherapyrecord});
    setStatus(true);
    if (res.id != ''){
      setAlert(true);
      window.location.reload(false);
    } else {
      setAlert(false);
    }
 
  };
 
  const Patiant_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientid(event.target.value as number);
    };

   const Personnel_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnelid(event.target.value as number);
    };

    const Physicaltherapyroom_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setPhysicaltherapyroomid(event.target.value as number);
     };

     const Status_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setStatusid(event.target.value as number);
     }; 

     const Added_Time_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setAddedTime(event.target.value as string);
    };

 
  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`ระบบบันทึกนัดกายภาพบำบัด`}>
      </Header>
      <Content>
      <ContentHeader title="ระบบบันทึกนัดกายภาพบำบัด">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
      </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>Personnel</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>TestPersonnel</InputLabel>
                <Select
                  name="personnel"
                  value={personnelid || ''}
                  onChange={Personnel_handleChange}
                >
                  {personnels.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>Patient</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Patient</InputLabel>
                <Select
                  name="patient"
                  value={patientid || ''}
                  onChange={Patiant_handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>Physicaltherapyroom</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Physicaltherapyroom</InputLabel>
                <Select
                  name="Physicaltherapyroom"
                  value={physicaltherapyroomid || ''}
                  onChange={Physicaltherapyroom_handleChange}
                >
                  {physicaltherapyrooms.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.physicaltherapyroomname}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>Status</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Status</InputLabel>
                <Select
                  name="Status"
                  value={statusid || ''}
                  onChange={Status_handleChange}
                >
                  {statuss.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.statusname}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            
            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
            <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="added"
                  type="datetime-local"
                  value={addedTime|| ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={Added_Time_handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  createPhysicaltherapyrecord();
                }}
              >
                บันทึก
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};
