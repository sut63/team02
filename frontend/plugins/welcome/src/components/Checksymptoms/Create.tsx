import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import SaveIcon from '@material-ui/icons/Save';


import InputAdornment from '@material-ui/core/InputAdornment';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import clsx from 'clsx';

import { EntDoctorordersheet } from '../../api/models/EntDoctorordersheet';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntDisease } from './../../api/models/EntDisease';




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


export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบนัดหมาย' };
 const api = new DefaultApi();

 //const [user, setUser] = useState(initialUserState);
 const [diseases, setDiseases] = useState<EntDisease[]>([]);
 const [doctorordersheets, setDoctorOrderSheets] = useState<EntDoctorordersheet[]>([]);
 const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
 const [patients, setPatients] = useState<EntPatient[]>([]);
 
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 const [loading, setLoading] = useState(true);

 //const [user, setUser] = useState(String);
 

 
 const [doctorordersheetids, setDoctorOrderSheetid] = useState(Number);
 const [patientids, setPatientid] = useState(Number);
 const [personnelids, setPersonnelid] = useState(Number);
 const [diseaseids, setDiseaseid] = useState(Number);
 const [date, setDate] = useState(String);
 const [times, setTimes] = useState(String);
 const [note, setNote] = useState(String);
 
 useEffect(() => {
   

   const getDiseases = async () => {
    const r = await api.listDisease({ limit: 100, offset: 0 });
    setLoading(false);
    setDiseases(r);
  };
  getDiseases();

   const getDoctorOrderSheets = async () => {
    const doc = await api.listDoctorordersheet({ limit: 7, offset: 0 });
    setLoading(false);
    setDoctorOrderSheets(doc);
  };
  getDoctorOrderSheets();  

  const getPatients = async () => {
    const ut = await api.listPatient({ limit: 2, offset: 0 });
    setLoading(false);
    setPatients(ut);
  };
  getPatients();

  const getPersonnels = async () => {
    const t = await api.listPersonnel({ limit: 2, offset: 0 });
    setLoading(false);
    setPersonnels(t);
  };
  getPersonnels();


 }, [loading]);
 
 
 const CreateChecksymptoms = async () =>{
  const checksymptoms = {

    doctorordersheet : doctorordersheetids,
    disease    : diseaseids,
    patient : patientids,
    personnel : personnelids,
    date : date + ":T00:00:00Z",
    times : times,
    note : note,
  }
  debugger
  console.log(checksymptoms)
   const res:any = await api.createChecksymptom({ checksymptom : checksymptoms });
   setStatus(true);
   if (res.id != ''){
     setAlert(true);

   } 
   
  
   else {
     setAlert(false);
   }
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };


 const disease_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setDiseaseid(event.target.value as number);
 };


const doctorordersheet_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setDoctorOrderSheetid(event.target.value as number);
 };

 const patient_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPatientid(event.target.value as number);
 };

 const personnel_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPersonnelid(event.target.value as number);
 };

 const handleDateChange = (event: any) => {
  setDate(event.target.value as string);
 };

 const handleTimesChange = (event: any) => {
  setTimes(event.target.value as string);
 };

 const handleNoteChange = (event: any) => {
  setNote(event.target.value as string);
 };



 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || 'ระบบประว้ติ'}`}
     ></Header>
     <Content>
       <ContentHeader title="ระบบสมัคร">
      
       


          {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                ✔️ ADD DATA COMPLETE : 😁 😆 🤩!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ❌ CAN'T ADD DATA : 🥺 😵 😱!
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
            <InputLabel id="personnel-label">personnel</InputLabel>
            <Select
              labelId="personnel-label"
              label="personnel"
              id="personnel_id"
              value={personnelids}
              onChange={personnel_id_handleChange}
              style = {{width: 400}}
            >
              {personnels.map((item:EntPersonnel)=>
              <MenuItem value={item.id}>{item.name}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>


 <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="disease-label">disease</InputLabel>
            <Select
              labelId="disease-label"
              label="disease"
              id="disease_id"
              value={diseaseids}
              onChange={disease_id_handleChange}
              style = {{width: 400}}
            >
              {diseases.map((item:EntDisease)=>
              <MenuItem value={item.id}>{item.disease}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>




         <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="doctorordersheet-label">doctorordersheet</InputLabel>
            <Select
              labelId="doctorordersheet-label"
              label="doctorordersheet"
              id="doctorordersheet_id"
              value={doctorordersheetids}
              onChange={doctorordersheet_id_handleChange}
              style = {{width: 400}}
            >
              {doctorordersheets.map((item:EntDoctorordersheet)=>
              <MenuItem value={item.id}>{item.name}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>




 <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="patient-label">patient</InputLabel>
            <Select
              labelId="patient-label"
              label="Patient"
              id="patientID"
              value={patientids}
              onChange={patient_id_handleChange}
              style = {{width: 400}}
            >
              {patients.map((item:EntPatient)=>
              <MenuItem value={item.id}>{item.name}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>



         

 <form className={classes.margin} >
        <TextField
    id="date"
    label="วันที่"
    type="date"
    value={date}
    defaultValue="24-05-2777"
    onChange={handleDateChange}
    className={classes.textField}
    style = {{width: 400}}
    InputLabelProps={{
      shrink: true,
    }}
  />
</form>



<div>
           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="times"
               label="times"
               variant="outlined"
               type="string"
               size="medium"
               value={times}
               onChange={handleTimesChange}
               style = {{width: 200}}
             />
           </FormControl>
           </div>



           <div>
           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="note"
               label="note"
               variant="outlined"
               type="string"
               size="medium"
               value={note}
               onChange={handleNoteChange}
               style = {{width: 1000}}
             />
           </FormControl>
           </div>




           
           <div>
           
           <FormControl
           
            // fullWidth
            
             
             
           >
       
             
           </FormControl>


           </div>

           
           <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateChecksymptoms();
                }}
              >
                บันทึก
              </Button>


           
         </form>
       </div>

       <script>
       
       </script>
     </Content>
   </Page>
 );
}