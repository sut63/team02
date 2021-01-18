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
import { EntDisease } from './../../api/models/EntDisease';
import { EntDoctorordersheet } from './../../api/models/EntDoctorordersheet';


// css style  
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },

  margin: {
    margin: theme.spacing(3),
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
  const [diseases, setDiseases] = useState<EntDisease[]>([]);
  const [doctorordersheets, setDoctorordersheets] = useState<EntDoctorordersheet[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [diseaseid, setDiseaseid] = useState(Number);
  const [doctorordersheetid, setDoctorordersheetid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
  const [times, setTimes] = useState(String);
  const [note, setNote] = useState(String);
 
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
    const getDiseases = async () => {
 
     const diseases = await api.listDisease({ limit: 10, offset: 0 });
       setLoading(false);
       setDiseases(diseases);
     };
     getDiseases();
//
const getDoctorordersheets = async () => {
 
  const doctorordersheets = await api.listDoctorordersheet({ limit: 10, offset: 0 });
    setLoading(false);
    setDoctorordersheets(doctorordersheets);
  };
  getDoctorordersheets();
//
 
  }, [loading]);
 
  const CreateChecksymptom = async () => {
    if((addedTime != "" )&& (patientid != 0) &&(personnelid != 0)&&(doctorordersheetid != 0)&&
    (diseaseid != 0)&&(times != "")&&(note != "")){
      const checksymptom = {
         date        : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         patientID          : patientid,
         personnelID        : personnelid,
         doctorordersheetID           : doctorordersheetid,
         diseaseID      : diseaseid,
         times : times,
         note : note,
      }
      console.log(checksymptom);

    const res:any = await api.createChecksymptom({ checksymptom : checksymptom});
    setStatus(true);
    if (res.id != ''){
      setAlert(true);
    }
    } else {
      setAlert(false);
      setStatus(true);
    }
 
  };
 
  const Patiant_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientid(event.target.value as number);
    };

   const Personnel_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnelid(event.target.value as number);
    };

    const Disease_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setDiseaseid(event.target.value as number);
     };

     const Doctorordersheet_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setDoctorordersheetid(event.target.value as number);
     };

     const Added_Time_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setAddedTime(event.target.value as string);
    };

    const TimeshandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setTimes(event.target.value as string);
     };

     const NotehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setNote(event.target.value as string);
     };

 
  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`ระบบบันทึกการนัดตรวจอาการ`}>
      </Header>
      <Content>
      <ContentHeader title="ระบบบันทึกการนัดตรวจอาการ">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกสำเร็จ
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 30 }}>
                 บันทึกไม่สำเร็จ ไปทำมาใหม่
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
              <div className={classes.paper}>Disease</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกโรค</InputLabel>
                <Select
                  name="Disease"
                  value={diseaseid || ''}
                  onChange={Disease_handleChange}
                >
                  {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.disease}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={4}>
              <div className={classes.paper}>Doctorordersheet</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>นายแพทย์</InputLabel>
                <Select
                  name="doctorordersheet"
                  value={doctorordersheetid || ''}
                  onChange={Doctorordersheet_handleChange}
                >
                  {doctorordersheets.map(item => {
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
              <div className={classes.paper}>หมายเหตุ</div>
            </Grid>
            <Grid item xs={8}>
              <TextField
                id="note"
                label="Note"
                variant="outlined"
                type="string"
                size="medium"
                value={note}
                onChange={NotehandleChange}
                style = {{width: 300}}
              />
            </Grid>




            <Grid item xs={4}>
              <div className={classes.paper}>วันที่</div>
            </Grid>
            <Grid item xs={8}>
            <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันและเวลา"
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








            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
              <TextField
                id="times"
                label="Times"
                variant="outlined"
                type="string"
                size="medium"
                value={times}
                onChange={TimeshandleChange}
                style = {{width: 300}}
              />
            </Grid>


            <div className={classes.margin}>
                    <Button
                    onClick={() => {
                      CreateChecksymptom();}}
                      variant="contained"
                      color="primary"
                    >
                      ยืนยัน
              </Button>
                                           
        </div>






        
        

          </Grid>
        </Container>
      </Content>
    </Page>
  );
};
