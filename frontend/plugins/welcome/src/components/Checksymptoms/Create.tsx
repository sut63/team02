import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader, 
} from '@backstage/core';

import {
  Grid,
  InputLabel,
  Container,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';

import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
//import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Swal from 'sweetalert2';

import { EntPatient,EntDoctorordersheet,EntPersonnel,EntDisease } from '../../api';



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

//updateforlab9
const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: (toast: { addEventListener: (arg0: string, arg1: any) => void; }) => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

export default function CreateChecksymptom() {
  const classes = useStyles();
  const profile = { givenName: ''};
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);

  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [diseases, setDiseases] = useState<EntDisease[]>([]);
  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
  const [doctorordersheets, setDoctorordersheets] = useState<EntDoctorordersheet[]>([]);

  const [loading, setLoading] = useState(true);

  const [patientName, setPatient] = useState(Number);
  const [diseaseName, setDisease] = useState(Number);
  const [personnelName, setPersonnel] = useState(Number);
  const [doctorordersheetName, setDoctorordersheet] = useState(Number);
  const [datetime, setDateTime] = useState(String);
  const [phone, setPhone] = useState(String);
  const [identitycard, setIdentitycard] = useState(String);
  const [note, setNote] = useState(String);

  useEffect(() => {

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
      console.log(res);
    };
    getPatients();

    const getDiseases = async () => {
      const res = await api.listDisease({ limit: 10, offset: 0 });
      setLoading(false);
      setDiseases(res);
    };
    getDiseases();

    const getDoctorordersheets = async () => {
      const res = await api.listDoctorordersheet({ limit: 10, offset: 0 });
      setLoading(false);
      setDoctorordersheets(res);
    };
    getDoctorordersheets();


    const getPersonnels = async () => {
      const res = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnels(res);
    };
    getPersonnels();
  }, [loading]);

  //updateforlab9
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      
      case 'phone':
        alertMessage("warning","กรุณากรอกเบอร์โทรศัพท์ให้ครบ 10 ตัว");
        return;
      case 'Identitycard':
        alertMessage("warning","กรุณากรอกเลขบัตรประชาชนให้ครบ 13 ตัว");
        return;
      case 'note':
        alertMessage("warning","กรุณากรอกข้อความไม่เกิน 40 ตัวอักษร ถ้าไม่มีกรอก '-' ");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  


  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/checksymptoms';
    const checksymptom = {
      patientID: patientName,
      diseaseID: diseaseName,
      personnelID: personnelName,
      doctorordersheetID: doctorordersheetName,
      date: datetime + ":00+07:00",
      phone: phone,
      identitycard: identitycard,
      note: note,
    };
    console.log(checksymptom);

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(checksymptom),
    };

    console.log(checksymptom); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });window.setTimeout(function(){location.reload()},15000);
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const DiseasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDisease(event.target.value as number);
  };


  const PhonehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPhone(event.target.value as string);
  };

  const PersonnelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnel(event.target.value as number);
  };
  const DateTimehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDateTime(event.target.value as string);
  };

  const IdentitycardhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setIdentitycard(event.target.value as string);
  };

  const DoctorordersheethandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDoctorordersheet(event.target.value as number);
  };

  const NotehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNote(event.target.value as string);
  };




  return (
    <Page theme={pageTheme.website}>
      <Header title={`ยินดีต้อนรับสู่ระบบบันทึกการนัดตรวจอาการ`}>
      <Link component={RouterLink} to="/WelcomePage">
 <Button variant="contained" color="secondary" >
   หน้า Welcome
 </Button>
 </Link>
      </Header>
      <Content>
      <ContentHeader title="-ระบบบันทึกการนัดตรวจอาการ-">
         <Link component={RouterLink} to="/Checksymptomsearch">
 <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
   ค้นหา
 </Button>
 </Link>
 
 <Link component={RouterLink} to="/Checksymptommain">
 <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
   ตาราง
 </Button>
 </Link>
        
  

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
                  value={personnelName}
                  onChange={PersonnelhandleChange}
                  style={{ width: 300 }}
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
    <FormControl variant="outlined" >
      <InputLabel>เลือกผู้ป่วย</InputLabel>
      <Select
        name="patient"
        value={patientName || ''}
        onChange={PatienthandleChange}
        style={{ width: 300 }}
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
    <div >Disease</div>
  </Grid>
  <Grid item xs={8}>
    <FormControl variant="outlined" >
      <InputLabel>เลือกโรค</InputLabel>
      <Select
        name="disease"
        value={diseaseName || ''}
        onChange={DiseasehandleChange}
        style={{ width: 300 }}
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
    <div >Doctorordersheet</div>
  </Grid>
  <Grid item xs={8}>
    <FormControl variant="outlined" >
      
      <InputLabel>เลือกแพทย์ที่ทำการรักษา</InputLabel>
      <Select
        name="Doctorordersheet"
        value={doctorordersheetName || ''}
        onChange={DoctorordersheethandleChange}
        style={{ width: 300 }}
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
    <div >วันและเวลา</div>
  </Grid>
  <Grid item xs={8}>
  <form  noValidate>
      <TextField
        label="เลือกเวลา"
        name="added"
        type="datetime-local"
        value={datetime|| ''} // (undefined || '') = ''
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
        onChange={DateTimehandleChange}
      />
    </form>
  </Grid>



  <Grid item xs={4}>
    <div >เลขบัตรประชาชน</div>
  </Grid>
  <Grid item xs={8}>
    <TextField
      id="identitycard"
      label="Identitycard"
      variant="outlined"
      type="string"
      size="medium"
      value={identitycard}
      onChange={IdentitycardhandleChange}
      style = {{width: 300}}
    />
    </Grid>

    <Grid item xs={4}>
    <div >เบอร์โทรศัพท์</div>
  </Grid>
  <Grid item xs={8}>
    <TextField
      id="phone"
      label="Phone"
      variant="outlined"
      type="string"
      size="medium"
      value={phone}
      onChange={PhonehandleChange}
      style = {{width: 300}}
    />
  </Grid>


  <Grid item xs={4}>
    <div >หมายเหตุ</div>
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

            

 
  <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  save();
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