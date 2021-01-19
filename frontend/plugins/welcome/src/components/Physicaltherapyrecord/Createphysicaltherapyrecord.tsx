import React, { useState, useEffect } from 'react';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
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
  Link,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntPhysicaltherapyroom } from '../../api/models/EntPhysicaltherapyroom';
import { EntStatus } from '../../api/models/EntStatus';
import Physicaltherapyrecord from '.';

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
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
 
 
  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [physicaltherapyroomid, setPhysicaltherapyroomid] = useState(Number);
  const [statusid, setStatusid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
 
 
  useEffect(() => {
// 
    const getPersonnels = async () => {
 
      const personnels = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnels(personnels);
    };
    getPersonnels();
// 
    const getPatients = async () => {
 
    const patients = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(patients);
    };
    getPatients();
// 
    const getPhysicaltherapyrooms = async () => {
 
     const physicaltherapyrooms = await api.listPhysicaltherapyroom({ limit: 10, offset: 0 });
       setLoading(false);
       setPhysicaltherapyrooms (physicaltherapyrooms);
     };
     getPhysicaltherapyrooms();
//
const getStatuss = async () => {
 
     const statuss = await api.listStatus({ limit: 10, offset: 0 });
       setLoading(false);
       setStatuss(statuss);
     };
     getStatuss();
 
  }, [loading]);
 
  const CreatePhysicaltherapyrecord = async () => {
    if((addedTime != null)&& (addedTime != "") && (patientid != null) && (personnelid != null) && (physicaltherapyroomid != null) && (statusid != null)){
      const physicaltherapyrecord = {
         time          : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         patient          : patientid,
         personnel        : personnelid,
         physicaltherapyroom : physicaltherapyroomid,
        status  : statusid,

      }
      console.log(physicaltherapyrecord);

      const res:any = await api.createPhysicaltherapyrecord({ physicaltherapyrecord : physicaltherapyrecord});
             setStatus(true);
            if (res.id != '') {
                setAlert(true);
                setTimeout(() => {
                  setStatus(false);
                }, 10000);
   
            }
        }
        else {
            setStatus(true);
            setAlert(false);
            setTimeout(() => {
              setStatus(false);
            }, 10000);
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
      <Header style={HeaderCustom} title={`ระบบบันทึกการนัดกายภาพบำบัด`}>
      </Header>
      <Content>
      <ContentHeader title="ระบบบันทึกการนัดกายภาพบำบัด">
        
      {status ? (
         <div>
           {alert ? (
             <Alert variant="filled" severity="success">
               บันทึกเรียบร้อย!
             </Alert>
           ) : (<Alert variant="filled" severity="error">
               บันทึกไม่สำเร็จ!
           </Alert>)}
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
                <InputLabel>เลือกผู้ใช้งาน</InputLabel>
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
                <InputLabel>เลือกผู้ป่วย</InputLabel>
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
                <InputLabel>เลือกห้อง</InputLabel>
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
                <InputLabel>เลือกสถานะ</InputLabel>
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
                  CreatePhysicaltherapyrecord();
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
