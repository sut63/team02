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
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
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
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'Checksymptom '}`}
      subtitle=""
     ></Header>
      <Content>
        <ContentHeader title="บันทึกนัดตรวจอาการ">
          <div>
            <Link component={RouterLink} to="/WelcomePage">
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
                <Typography variant="h6" gutterBottom  align="center">
                Personnel Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                  labelId="personnel"
                  id="personnel"
                  value={personnelName}
                  onChange={PersonnelhandleChange}
                  style={{ width: 400 }}
                >
               {personnels.map(item => {
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
<div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography variant="h6" gutterBottom  align="center">
                Disease Name : 
                <Typography variant="body1" gutterBottom> 
              <Select
                labelId="disease"
                id="disease"
                value={diseaseName}
                onChange={DiseasehandleChange}
                style={{ width: 400 }}
              > 
                {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.disease}
                      </MenuItem>
                    );
                  })}
              </Select>
                </Typography>
                </Typography>
            </FormControl>
            </div>

            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography variant="h6" gutterBottom  align="center">
                Doctorordersheet Name : 
                <Typography variant="body1" gutterBottom> 
              <Select
                labelId="doctorordersheet"
                id="doctorordersheet"
                value={doctorordersheetName}
                onChange={DoctorordersheethandleChange}
                style={{ width: 400 }}
              > 
                {doctorordersheets.map(item => {
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




          <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>เลขบัตรประชาชน</div>
                <TextField id="identitycard"  InputLabelProps={{
                  shrink: true,
                }} label="identitycard" variant="outlined"
                  onChange={IdentitycardhandleChange}
                  value={identitycard || ''}
                />
              </FormControl>
            </div>


            

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>เบอร์โทรศัพท์</div>
                <TextField id="phone"  InputLabelProps={{
                  shrink: true,
                }} label="phone" variant="outlined"
                  onChange={PhonehandleChange}
                  value={phone || ''}
                />
              </FormControl>
            </div>


            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>หมายเหตุ * (ไม่เกิน 40 ตัวอักษร ถ้าไม่มีกรอก -)</div>
                <TextField id="note" type='string' InputLabelProps={{
                  shrink: true,
                }} label="note" variant="outlined"
                  onChange={NotehandleChange}
                  value={note || ''}
                />
              </FormControl>
            </div>

                        
            <div className={classes.margin}>
            <Typography variant="h6" gutterBottom  align="center">
              <Button
                onClick={() => {
                  save();
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