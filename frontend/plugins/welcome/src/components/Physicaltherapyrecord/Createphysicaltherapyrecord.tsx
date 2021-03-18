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
  const [priceerror, setPriceerror] = React.useState('');
  const [phoneerror, setPhoneerror] = React.useState('');
  const [noteerror, setNoteerror] = React.useState('');
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);


  const [loading, setLoading] = useState(true);

  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [physicaltherapyroomid, setPhysicaltherapyroomid] = useState(Number);
  const [statusid, setStatusid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
  const [priceid, setPriceid] = useState(Number);
  const [phoneid, setPhoneid] = useState(String);
  const [noteid, setNoteid] = useState(String);





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

  const PricehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPriceid(event.target.value as number);
  };
  const PhonehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPhoneid(event.target.value as string);
  };
  const NotehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNoteid(event.target.value as string);
  };




  const listphre = () => {
    setStatus(false);
    if(errormessege == "บันทึกข้อมูลสำเร็จ"){
    window.location.href ="http://localhost:3000/physicaltherapyrecords";
    }
  };

  
 
  const checkCaseSaveError = (field: string) => {
    
    if (field == "price") { setErrorMessege("กรอกจำนวนเงินให้ถูกต้อง"); }
        else if (field == "phone") { setErrorMessege("กรอกเบอร์โทรให้ถูกต้อง"); }
        else if (field == "note") { setErrorMessege("กรอกหมายเหตุให้ถูกต้อง"); }
        else { setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ"); }
  }
  const CreatePhysicaltherapyrecord = async ()=>{
  
    
    const physicaltherapyrecord ={
      time          : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      patient          : patientid,
      personnel        : personnelid,
      physicaltherapyroom : physicaltherapyroomid,
      status  : statusid,
      phone : phoneid,
      note : noteid,
      price  : Number(priceid) ,
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
            





          <div className={classes.paper}><strong>จำนวนเงิน</strong></div>

            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
      
              id="priceid"
              error = {priceerror ? true : false}
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {priceerror}
              value={priceid}
              onChange={PricehandleChange}
            />
            
            <div className={classes.paper}><strong>เบอร์โทร</strong></div>

<TextField className={classes.textField}
style={{ width: 400 ,marginLeft:20,marginRight:-10}}

  id="phoneid"
  error = {phoneerror ? true : false}
  label=""
  variant="standard"
  color="secondary"
  type="string"
  size="medium"
  helperText= {phoneerror}
  value={phoneid}
  onChange={PhonehandleChange}
/>



            <div className={classes.paper}><strong>หมายเหตุ</strong></div>

<TextField className={classes.textField}
style={{ width: 400 ,marginLeft:20,marginRight:-10}}

  id="noteid"
  error = {noteerror ? true : false}
  label=""
  variant="standard"
  color="secondary"
  type="string"
  size="medium"
  helperText= {noteerror}
  value={noteid}
  onChange={NotehandleChange}
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