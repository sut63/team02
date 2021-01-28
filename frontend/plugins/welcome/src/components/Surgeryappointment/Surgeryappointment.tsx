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
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import FormControl from '@material-ui/core/FormControl';

import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntPatient } from '../../api/models/EntPatient';
import { EntSurgerytype } from '../../api/models/EntSurgerytype';
import { EntSurgeryappointment } from '../../api/models/EntSurgeryappointment';
import TextField from '@material-ui/core/TextField';

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
  const profile = { givenName: 'ระบบบันทึกการนัดผ่าตัด' };
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  const [personnels, setPersonnel] = useState<EntPersonnel[]>([]);
  const [patients, setPatient] = useState<EntPatient[]>([]);
  const [surgerytypes, setSurgerytype] = useState<EntSurgerytype[]>([]);

  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [personnelid, setPersonnelID] = useState(Number);
  const [patientid, setPatientID] = useState(Number);
  const [surgerytypeid, setSurgerytypeID] = useState(Number);
  const [addtime, setAddtime] = useState(String);
  const [phone, setPhone] = useState(String);
  const [age, setAge] = useState(Number);
  const [note, setNote] = useState(String);

  useEffect(() => {
    const getPersonnel = async () => {
      const res = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnel(res);
    };
    getPersonnel();

    const getPatient = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatient(res);
    };
    getPatient();

    const getSurgerytype = async () => {
      const res = await api.listSurgerytype({ limit: 10, offset: 0 });
      setLoading(false);
      setSurgerytype(res);
    };
    getSurgerytype();

  }, [loading]);

  
   const AddtimehandleChange = (event: any) => {
    setAddtime(event.target.value as string);
   };
   const PhonehandleChange = (event: any) => {
    setPhone(event.target.value as string);
   };

   const AgehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setAge(event.target.value as number);
   };

   const NotehandleChange = (event: any) => {
    setNote(event.target.value as string);
   };

  const PersonnelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnelID(event.target.value as number);
  };

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
  };

  const SurgerytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSurgerytypeID(event.target.value as number);
  };


  const CreateSurgeryappointment = async () => {

    const surgeryappointment = {
      personelid: personnelid,
      patientid: patientid,
      surgerytypeid: surgerytypeid,
      appointtime: addtime + ":00+07:00" ,
      phone: phone,
      note: note,
      age: Number(age)
    };
    console.log(surgeryappointment);
    console.log(ErrorCaseCheck);
      const apiUrl = 'http://localhost:8080/api/v1/surgeryappointments';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(surgeryappointment),
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            setErrorMessege("บันทึกข้อมูลสำเร็จ");
            setAlertType("success");

          }
          else {
            ErrorCaseCheck(data.error.Name);
            setAlertType("error");
          }
          
        });
      setStatus(true);
  };

  const ErrorCaseCheck = (casename: string) => {
    if (casename == "age") { setErrorMessege("กรุณาใส่อายุให้ถูกต้อง"); }
    else if (casename == "phone") { setErrorMessege("กรุณาใส่เบอร์โทรศัพท์ให้ถูกต้อง"); }
    else if (casename == "note") { setErrorMessege("กรุณณาใส่หมายเหตุให้ถูกต้อง"); }
    else { setErrorMessege("บันทึกไม่สำเร็จ"); }
  }
  

  return (
    <Page theme={pageTheme.website}>
      <Header
        title={` ${profile.givenName || 'to Backstage'}`}
        subtitle=""
      ></Header>
      <Content>
        <ContentHeader title="นัดหมายผ่าตัด">
        <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
             </Button>
        </ContentHeader>
        <div>
            <p><font size='2'><b>แพทย์</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="personnel-label"></InputLabel>
              <Select
                labelId="personnel-label"
                id="personnel"
                value={personnelid}
                onChange={PersonnelhandleChange}
                style={{ width: 250 }}
              >
                {personnels.map((item: EntPersonnel) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>
          <div>
            <p><font size='2'><b>คนไข้</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="patient-label"></InputLabel>
              <Select
                labelId="patient-label"
                id="patient"
                value={patientid}
                onChange={PatienthandleChange}
                style={{ width: 250 }}
              >
                {patients.map((item: EntPatient) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>
          <div>
            <p><font size='2'><b>ประเภทการผ่าตัด</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="surgerytype-label"></InputLabel>
              <Select
                labelId="surgerytype-label"
                id="surgerytype"
                value={surgerytypeid}
                onChange={SurgerytypehandleChange}
                style={{ width: 250 }}
              >
                {surgerytypes.map((item: EntSurgerytype ) => (
                  <MenuItem value={item.id}>{item.typename}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>
          <p><font size='2'><b>อายุผู้เข้ารับการผ่าตัด</b></font></p>  
          <TextField className={classes.textField}
                style={{ width:  200 ,marginLeft:20,marginRight:-10}}
              id="age"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={age}
              onChange={AgehandleChange}
            /> 
          <div>
          <p><font size='2'><b>เวลา</b></font></p> 
          <TextField
        id="datetime-local"
        label="Next appointment"
        type="datetime-local"
        //defaultValue="2020-10-24T11:30"
        value={addtime}
        className={classes.textField}
        onChange={AddtimehandleChange}
        InputLabelProps={{
          shrink: true,
        }}
      />
       <p><font size='2'><b>เบอร์โทรศัพท์</b></font></p>  
          <TextField className={classes.textField}
                style={{ width:  200 ,marginLeft:20,marginRight:-10}}
              id="phone"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={phone}
              onChange={PhonehandleChange}
            />
          <p><font size='2'><b> Note </b></font></p>  
          <TextField className={classes.textField}
                style={{ width:  200 ,marginLeft:20,marginRight:-10}}
              id="note"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={note}
              onChange={NotehandleChange}
            />
          </div>
          {status ? (
            <div>
              {alerttype != "" ? (
                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                  {errormessege}
                </Alert>
              ) : null}
            </div>
          ) : null}
          <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateSurgeryappointment();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
            </div>
          

        
      </Content>
    </Page>
  );
}