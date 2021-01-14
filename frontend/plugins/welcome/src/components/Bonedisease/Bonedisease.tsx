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
import { EntRemedy } from '../../api/models/EntRemedy';

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
  const [remedys, setRemedys] = useState<EntRemedy[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [remedyid, setRemedyid] = useState(Number);
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
    const getRemedys = async () => {
 
     const remedys = await api.listRemedy({ limit: 10, offset: 0 });
       setLoading(false);
       setRemedys(remedys);
     };
     getRemedys();
//
 
  }, [loading]);
 
  const CreateBonedisease = async () => {
      const Bonedisease = {
         addedTime      : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         patientID          : +patientid,
         personnelID        : +personnelid,
         remedyID    : +remedyid,
      }
      console.log(Bonedisease);

    const res:any = await api.createBonedisease({ bonedisease : Bonedisease});
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

    const Remedy_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setRemedyid(event.target.value as number);
     };

     const Added_Time_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setAddedTime(event.target.value as string);
    };

 
  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`ระบบบันทึกการนัดตรวจโรคกระดูก`}>
      </Header>
      <Content>
      <ContentHeader title="ระบบบันทึกการนัดตรวจโรคกระดูก">
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
              <div className={classes.paper}>Remedy</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Remedy</InputLabel>
                <Select
                  name="Remedy"
                  value={remedyid || ''}
                  onChange={Remedy_handleChange}
                >
                  {remedys.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.remedy}
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
                  CreateBonedisease();
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
