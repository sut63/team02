import React, { useState, useEffect } from 'react';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles, formatMs } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import SaveIcon from '@material-ui/icons/Save';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntPhysicaltherapyroom } from '../../api/models/EntPhysicaltherapyroom';
import { EntStatus } from '../../api/models/EntStatus';
import { EntPhysicaltherapyrecord } from '../../api/models/EntPhysicaltherapyrecord';
import { Link as RouterLink } from 'react-router-dom';


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
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 400,
    },
  }),
);


export default function AmbulanceCreate() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ระบบบันทึกนัดกายกายภาพบำบัด' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);

  //เก็บข้อมูลที่จะดึงมา
  const [physicaltherapyrecord, setPhysicaltherapyrecords] = useState<EntPhysicaltherapyrecord[]>([]);
  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [physicaltherapyrooms, setPhysicaltherapyrooms] = useState<EntPhysicaltherapyroom[]>([]);
  const [statuss, setStatuss] = useState<EntStatus[]>([]); 
  const [idnumbererror, setIdnumbererror] = React.useState('');
  const [ageerror, setAgeerror] = React.useState('');
  const [telophoneerror, setTelophoneerror] = React.useState('');
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);


  const [loading, setLoading] = useState(true);

  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [physicaltherapyroomid, setPhysicaltherapyroomid] = useState(Number);
  const [statusid, setStatusid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
  const [idnumberid, setIDnumberid] = useState(String);
  const [ageid, setAgeid] = useState(Number);
  const [telephoneid, setTelephoneid] = useState(String);




  useEffect(() => {
    // 
        const getPersonnels = async () => {
          setPersonnelid(Number(localStorage.getItem("personnel")))

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
  const idnumber = (val: string) => {
    return val.match("[0-9]\\d{12}") && val.length == 13 ? true:false;
  
  }

  const age = (val: number) => {
    return val <= 1 && val >=99 ? true:false
  }

  const telephone = (val: string) => {
    return val.match("[0]\\d{9}") && val.length == 10 ? true:false;
  }
  const checkPattern  = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'idnumber':
        idnumber(value) ? setIdnumbererror('') : setIdnumbererror('เลขเลขบัตรประชาชน 13 ตัว');
        return;
      case 'age':
        age(Number(value)) ? setAgeerror('') : setAgeerror('อายุ');
      return;
      case 'telephone':
        telephone(value) ? setTelophoneerror('') : setTelophoneerror('เบอร์โทรศัพท์');
      return;
        default:
          return;
    }
  }
 
  console.log(personnelid)

  const PersonnelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonnelid(event.target.value as number);
  };
  const Patiant_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientid(event.target.value as number);
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

  const IDnumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setIDnumberid(event.target.value as string);
  };

  const AgehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAgeid(event.target.value as number);
  };

  const TelephonehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTelephoneid(event.target.value as string);
  };



  const listphre = () => {
    setStatus(false);
    if(errormessege == "บันทึกข้อมูลสำเร็จ"){
    window.location.href ="http://localhost:3000/physicaltherapyrecords";
    }
  };

 
  const IDnumberhandlehange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof idnumberid;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setIDnumberid(event.target.value as string);
    
  };
  
  
  
 
  const checkCaseSaveError = (field: string) => {
    if (field == "idnumber") { setErrorMessege("ข้อมูลfield เลขบัตรประชาชนผิด"); }
        else if (field == "age") { setErrorMessege("ข้อมูลfield อายุผิด"); }
        else if (field == "telephone") { setErrorMessege("ข้อมูลfield โทรศัพท์ผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ"); }
  }
  const CreatePhysicaltherapyrecord = async ()=>{
  
    
    const physicaltherapyrecord ={
      time          : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      patient          : patientid,
      personnel        : personnelid,
      physicaltherapyroom : physicaltherapyroomid,
      status  : statusid,
      age : Number(ageid),
      telephone :telephoneid,
      idnumber  : idnumberid,
    };

    console.log(physicaltherapyrecord)
    const apiUrl = 'http://localhost:8080/api/v1/physicaltherapyrecords';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(physicaltherapyrecord),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        setStatus(true);
        if (data.status === true) {
          setErrorMessege("บันทึกข้อมูลสำเร็จ");
          setAlertType("success");

      }
      else {
        checkCaseSaveError(data.error.Name);
          setAlertType("error");
      }
  });

     };
     
        


  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      >
       <table>
       <tr>
        <th>
  
        <Link component={RouterLink} to="/WelcomePage">
                        <Button variant="contained" color="secondary" >
                            กลับหน้าหลัก
                        </Button>
                    </Link>
          </th>
       </tr>
       </table>

      </Header>
      <Content>
        
      <ContentHeader title="ระบบบันทึกนัดกายกายภาพบำบัด" >
        <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/Searchphysicaltherapyrecord"
                variant="contained"
                size="large"
              >
                ค้นหา
             </Button>
             <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/WelcomePage"
                variant="contained"
                size="large"
              >
                กลับ
             </Button>    
      {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { listphre() }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">

          
              <div className={classes.paper}>Personnel</div>
            
              <FormControl variant="outlined" className={classes.formControl}>
              <TextField className={classes.textField}
                style={{ width:  200 ,marginLeft:20,marginRight:-10}}
              id="personnel"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={personnels.filter((filter: EntPersonnel) => filter.id == personnelid).map((item: EntPersonnel) => `${item.name}`)}
            />
              </FormControl>
            
            
              <div className={classes.paper}>Patient</div>
            
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
           

            
              <div className={classes.paper}>Physicaltherapyroom</div>
            
            
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
            

            
              <div className={classes.paper}>Status</div>
            
            
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
            





          <div className={classes.paper}><strong>เลขเลขบัตรประชาชน</strong></div>

            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
      
              id="idnumberid"
              error = {idnumbererror ? true : false}
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {idnumbererror}
              value={idnumberid}
              onChange={IDnumberhandlehange}
            />
            
            <div className={classes.paper}><strong>อายุ</strong></div>

<TextField className={classes.textField}
style={{ width: 400 ,marginLeft:20,marginRight:-10}}

  id="ageid"
  error = {ageerror ? true : false}
  label=""
  variant="standard"
  color="secondary"
  type="string"
  size="medium"
  helperText= {ageerror}
  value={ageid}
  onChange={AgehandleChange}
/>
<div className={classes.paper}><strong>เบอร์โทรศัพท์</strong></div>

<TextField className={classes.textField}
style={{ width: 400 ,marginLeft:20,marginRight:-10}}

  id="telephoneid"
  error = {telophoneerror ? true : false}
  label=""
  variant="standard"
  color="secondary"
  type="string"
  size="medium"
  helperText= {telophoneerror}
  value={telephoneid}
  onChange={TelephonehandleChange}
/>


            
            <div className={classes.paper}><strong>วันที่เวลาบันทึกข้อมูล</strong></div>
            
                <TextField
                    className={classes.formControl}
                    id="datetime"
                    type="datetime-local"
                    value={addedTime}
                    onChange={Added_Time_handleChange}
                    InputLabelProps={{
                     shrink: true,
                     }}
                 />
                <Typography align ="right"></Typography>
                
            <div className={classes.margin}>
              
              <Button
                onClick={() => {
                  
                  CreatePhysicaltherapyrecord();
                }}
                variant="outlined"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
              >
                บันทึกข้อมูล
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/WelcomePage"
                variant="contained"
                size="large"
              >
                กลับ
             </Button>
             
            </div>
          </form>
        </div>
        


        
        
      </Content>
    </Page>
  );
}