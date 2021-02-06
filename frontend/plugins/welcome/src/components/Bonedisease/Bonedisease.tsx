import React, { useEffect,useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link, Link as RouterLink } from 'react-router-dom'; //link
import SearchIcon from '@material-ui/icons/Search';


import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core'

import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntRemedy } from '../../api/models/EntRemedy';


import { ContentHeader } from '@backstage/core';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 2,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
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
  media: {
    height: 0,
    marginLeft: 25,
    maxWidth: 300,
}
}));

export default function Insurance() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [alerttype, setAlertType] = useState(String);

  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [remedys, setRemedys] = useState<EntRemedy[]>([]);
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
 
 
  const [personnelid, setPersonnelid] = useState(Number);
  const [patientid, setPatientid] = useState(Number);
  const [remedyid, setRemedyid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
  const [adviceid, setAdvice] = useState(String);
  const [identificationCardid, setIdentificationCard] = useState(String);
  const [telid, setTel] = useState(String);

  const [errorAdvice, setErrorAdvice] = React.useState('');
  const [errorTel, setErrorTel] = React.useState('');
  const [errorIdentificationCard, setErrorIdentificationCard] = React.useState('');
  

  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    //timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  
  const setErrorMessege = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
 
  useEffect(() => {
 
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

  }, [loading]);

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

  const Advice_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('advice', validateValue)
    setAdvice(event.target.value as string);
    };

    const Tel_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('tel', validateValue)
    setTel(event.target.value as string);
    };
    
    const  IdentificationCard_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('identificationCard', validateValue)
    setIdentificationCard(event.target.value as string);
    };

    const validateIdentificationCard = (val: string) => {
      return val.match("^[0-9]{13}$");
    }
    const validateTel = (val: string) => {
      return val.match("^[0]{1}[8926]{1}[0-9]{8}$");
    }
    const validateAdvice = (val: string) => {
      return val.match("^[ก-๙]");
    }
  
    

     const checkpattern = (id: string, value:string) => {
      console.log(value);
      switch(id) {
        case 'identificationCard':
          validateIdentificationCard(value) ? setErrorIdentificationCard('') : setErrorIdentificationCard('ใส่หมายบัตรประชาชนให้ถูกต้อง');
          return;
        case 'tel':
            validateTel(value) ? setErrorTel('') : setErrorTel('ใส่หมายเลขโทรศัพท์ให้ถูกต้อง');
            return;
        case 'advice' :
          validateAdvice(value) ? setErrorAdvice('') : setErrorAdvice('ในการกรอกข้อมูลต้องมีภาษาไทยด้วย');
          return;
        default:
          return;
      }
    }


  const CreateBonedisease = async () => {
    if((addedTime != null)&& (addedTime != "") && (patientid != null) && (personnelid != null) && (remedyid != null) ){
      const apiUrl = 'http://localhost:8080/api/v1/bonediseases';
      const bonedisease = {
        addedTime          : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
        patientID          : patientid,
        personnelID        : personnelid,
        remedyID           : remedyid,
        advice             : adviceid,
        tel                : telid,
        identificationCard : identificationCardid,
      };
      console.log(bonedisease);
      
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(bonedisease),
      };
        fetch(apiUrl, requestOptions)
          .then(response => response.json())
          .then(data => {
            console.log(data);
            if (data.status === true) {
               Toast.fire({
                  icon: 'success',
                  title: 'บันทึกข้อมูลสำเร็จ',
                });
              }
              else {
                ErrorCaseCheck(data.error.Name);
                setAlertType("error");
              }  
            });
        }
        else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
          setAlertType("error");
        }
          setStatus(true);
        
      };

      const ErrorCaseCheck = (field: string) => {
        switch(field) {
          case 'identificationCard':
            setErrorMessege("error","ใส่หมายเลขบัตรประชาชนให้ถูกต้อง");
            return;
          case 'tel':
            setErrorMessege("error","ใส่หมายเลขโทรศัพท์ให้ถูกต้อง");
            return;
          case 'advice':
            setErrorMessege("error","ในการกรอกข้อมูลต้องมีภาษาไทยด้วย");
            return;
          default:
            setErrorMessege("error","บันทึกข้อมูลไม่สำเร็จ");
            return;
        }
      }

    return (
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`ระบบบันทึกนัดตรวจโคกระดูก`}>
        </Header>
        <Content>
        <ContentHeader title="บันทึกนัดตรวจโคกระดูก">

     </ContentHeader>
          <Container maxWidth="sm">
            <Grid container spacing={3}>
              <Grid item xs={12}></Grid>
              <Grid item xs={4}>
                <div className={classes.paper}>Personnel</div>
              </Grid>
              <Grid item xs={8}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกผลิตภัณฑ์</InputLabel>
                  <Select
                    name="personnel"
                    value={personnelid || ''} // (undefined || '') = ''
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
              <div className={classes.paper}>Remedy</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกวิธีการรักษา</InputLabel>
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
              <div className={classes.paper}>เลขบัตรประชาชน</div>
            </Grid>
            <Grid item xs={8}>
              <TextField
                error = {errorIdentificationCard ? true : false}
                id="identificationCard"
                name="identificationCard"
                type="string"
                size="medium"
                value={identificationCardid}
                onChange={IdentificationCard_handleChange}
                InputLabelProps = {{
                  shrink: true,
                }}
                variant="outlined"
                style = {{width: 300}}
                helperText = {errorIdentificationCard}
              />
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>Tel</div>
            </Grid>
            <Grid item xs={8}>
              <TextField
                error = {errorTel ? true : false}
                id="tel"
                name="Tel"
                type="string"
                size="medium"
                value={telid}
                onChange={Tel_handleChange}
                InputLabelProps = {{
                  shrink: true,
                }}
                variant="outlined"
                style = {{width: 300}}
                helperText = {errorTel}
              />
            </Grid>
            
            <Grid item xs={4}>
              <div className={classes.paper}>เพิ่มเติมรายละเอียด</div>
            </Grid>
            <Grid item xs={8}>
              <TextField
                error = {errorAdvice ? true : false}
                id="advice"
                name="advice"
                type="string"
                size="medium"
                value={adviceid}
                onChange={Advice_handleChange}
                InputLabelProps={{
                  shrink: true,
                }}
                variant="outlined"
                style = {{width: 300}}
                helperText = {errorAdvice}
              />
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
            <Grid item xs={3}>
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

            <Grid item xs={3}>
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  component = {RouterLink}
                  to = "/BonediseaseSearch"
                  startIcon={<SearchIcon />}
                  
                >
                  Search
                </Button>
            </Grid>
            </Grid>
          </Container>
        </Content>
      </Page>
    );
  } 