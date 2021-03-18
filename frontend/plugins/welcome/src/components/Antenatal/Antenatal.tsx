import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
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
import 'date-fns';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntPatient } from '../../api/models/EntPatient';
import { EntGender } from '../../api/models/EntGender';
import { EntPregnancystatus } from '../../api/models/EntPregnancystatus';
import { EntRisks } from '../../api/models/EntRisks';
import { Alert } from '@material-ui/lab';
import Swal from 'sweetalert2'; // alert
import Typography from '@material-ui/core/Typography';
import SearchIcon from '@material-ui/icons/Search';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    formControl: {
      margin: theme.spacing(3),
      minWidth: 300,
    },
    selectEmpty: {
      marginTop: theme.spacing(1),
    },
    root: {
        '& > *': {
          margin: theme.spacing(3),
          width: '39ch',
        },
    },
    paper: {
        marginTop: theme.spacing(3),
        marginBottom: theme.spacing(3),
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        marginLeft: theme.spacing(3),
        marginRight: theme.spacing(3),
        width: 400,
    },
    margin: {
        margin: theme.spacing(3),
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

export default function createAntenatalinformation() {
    const classes = useStyles();
    const profile = { givenName: 'ระบบบันทึกนัดฝากครรภ์' };
    const api = new DefaultApi();

    const [personnels, setPersonnel] = useState<EntPersonnel[]>([]);
    const [patients, setPatient] = useState<EntPatient[]>([]);
    const [genders, setGender] = useState<EntGender[]>([]);
    const [pregnancystatuss, setPregnancystatus] = useState<EntPregnancystatus[]>([]);
    const [riskss, setRisks] = useState<EntRisks[]>([]);

    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);

    const [personnelid, setPersonnelid] = useState(Number);
    const [patientid, setPatientid] = useState(Number);
    const [pregnancystatusid, setPregnancystatusid] = useState(Number);
    const [risksid, setRisksid] = useState(Number);
    const [gestationalage, setGestationalage] = useState(Number);
    const [examinationresult, setExaminationresult] = useState(String);
    const [advice, setAdvice] = useState(String);
    const [time, setTime] = useState(String);


    useEffect(() => {
        const getPersonnel = async () => {
            setPersonnelid(Number(localStorage.getItem("personnel")))
            const per = await api.listPersonnel({ limit: 10, offset: 0 });
            setLoading(false);
            setPersonnel(per);
          };
        getPersonnel();
    
        const getPatient = async () => {
            const pat = await api.listPatient({ limit: 10, offset: 0 });
            setLoading(false);
            setPatient(pat);
            console.log(pat);
          };
        getPatient();
    
        const getPregnancystatus = async () => {
            const pre = await api.listPregnancystatus({ limit: 10, offset: 0 });
            setLoading(false);
            setPregnancystatus(pre);
          };
        getPregnancystatus();
    
        const getRisks = async () => {
            const r = await api.listRisks({ limit: 10, offset: 0 });
            setLoading(false);
            setRisks(r);
          };
        getRisks();
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
            case 'gestationalage' :
                alertMessage("error","รูปแบบอายุครรภ์ไม่ถูกต้อง");
                return;
            case 'examinationresult' :
                alertMessage("error","รูปแบบผลการตรวจไม่ถูกต้อง");
                return;
            case 'advice' :
                alertMessage("error","รูปแบบคำแนะนำไม่ถูกต้อง");
                return;
            default:
                alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
                return;
        }
    }

    function save(){
        const apiUrl = 'http://localhost:8080/api/v1/antenatalinformations';
        const antenatalinformation = {
            personnel: personnelid,
            patient: patientid,
            pregnancystatus: pregnancystatusid,
            risks: risksid,
            gestationalage: Number(gestationalage),
            examinationresult: examinationresult,
            advice: advice,
            time: time + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
        };
    
        console.log(antenatalinformation);
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(antenatalinformation),
        };

        console.log(antenatalinformation); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            if (data.status === true) {
            //clear();
            Toast.fire({
                icon: 'success',
                title: 'บันทึกข้อมูลสำเร็จ',
            });window.setTimeout(function(){location.reload()},5000);
            } else {
            checkCaseSaveError(data.error.Name)
            }
        });
    }
    
    const PersonnelhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPersonnelid(event.target.value as number);
    };
    const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPatientid(event.target.value as number);
    };
    const PregnancystatushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPregnancystatusid(event.target.value as number);
    };
    const RiskshandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setRisksid(event.target.value as number);
    };

    const GestationalagehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setGestationalage(event.target.value as number);
    };

    const ExaminationresulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setExaminationresult(event.target.value as string);
    };

    const AdvicehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAdvice(event.target.value as string);
    };

    const TimehandleDateChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setTime(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.service}>
            <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "./SearchAntenatalinformation");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 300, backgroundColor: "#54c571", color: "#000000", fontSize: 18 }}
                        startIcon={<SearchIcon />}
                    >
                        <b>ค้นหาผลการตรวจครรภ์</b>
                    </Button>
                </div>
            <Header title={`${profile.givenName || 'ระบบันทึกนัดฝากครรภ์'}`}></Header>
            <Content>
                <ContentHeader title="กรุณากรอกข้อมูล"></ContentHeader>
                    <Container maxWidth="sm">
                        <Grid container spacing={3}>
                            <Grid item xs={4}>
                                <div className={classes.root}>
                                    <form noValidate autoComplete="off">


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

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl variant="filled" className={classes.formControl}>
                                                    <InputLabel>ผู้รับการตรวจ</InputLabel>
                                                    <Select
                                                        labelId="ผู้รับการตรวจ"
                                                        id="patient"
                                                        value={patientid}
                                                        onChange={PatienthandleChange}
                                                    >
                                                        {patients.map((item: EntPatient) =>
                                                        <MenuItem value ={(item.id)}>{item.name}</MenuItem>)
                                                        }
                                                    </Select>
                                                </FormControl>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl variant="filled" className={classes.formControl}>
                                                    <InputLabel>สถานะครรภ์</InputLabel>
                                                    <Select
                                                        labelId="สถานะครรภ์"
                                                        id="pregnancystatus"
                                                        value={pregnancystatusid}
                                                        onChange={PregnancystatushandleChange}
                                                    >
                                                        {pregnancystatuss.map((item: EntPregnancystatus) =>
                                                        <MenuItem value={item.id}>{item.pregnancystatus}</MenuItem>)
                                                        }
                                                    </Select>
                                                </FormControl>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl className={classes.root} variant="outlined">
                                                    <InputLabel>สัปดาห์</InputLabel>
                                                    <TextField
                                                        id="gestationalage"
                                                        label="อายุครรภ์"
                                                        type="int"
                                                        value={gestationalage}
                                                        onChange={GestationalagehandleChange}
                                                    />
                                                </FormControl>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl className={classes.root} variant="outlined">
                                                    <TextField
                                                
                                                        id="examinationresult"
                                                        label="ผลการตรวจ"
                                                        type="string"
                                                       
                                                        value={examinationresult}
                                                        onChange={ExaminationresulthandleChange}
                                                    />
                                                </FormControl>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl variant="filled" className={classes.formControl}>
                                                    <InputLabel>ความเสี่ยงที่พบ</InputLabel>
                                                    <Select
                                                        labelId="ความเสี่ยงที่พบ"
                                                        id="risks"
                                                        value={risksid}
                                                        onChange={RiskshandleChange}
                                                    >
                                                        {riskss.map((item: EntRisks) =>
                                                        <MenuItem value={item.id}>{item.risks}</MenuItem>)
                                                        }
                                                    </Select>
                                                </FormControl>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={12}>
                                                <form className={classes.container} noValidate>
                                                    <TextField
                                                    label="เลือกเวลา"
                                                    name="time"
                                                    type="datetime-local"
                                                    value={time|| ''} // (undefined || '') = ''
                                                    className={classes.textField}
                                                    InputLabelProps={{
                                                        shrink: true,
                                                    }}
                                                    onChange={TimehandleDateChange}
                                                    />
                                                </form>
                                            </Grid>
                                        </div>

                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl className={classes.root} variant="outlined">
                                                    
                                                    <TextField
                                                        
                                                        id="advice"
                                                        label="คำแนะนำ"
                                                        type="string"
                                                       
                                                        value={advice}
                                                        onChange={AdvicehandleChange}
                                                    />
                                                </FormControl>
                                            </Grid>
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
                                                บันทึก
                                            </Button>

                                            &nbsp;&nbsp;&nbsp;
                                            <Link component={RouterLink} to="/WelcomePage">
                                            <Button variant="contained" color="primary" style={{backgroundColor: "#e55451"}}>
                                            กลับ
                                            </Button>
                                            </Link>
                                        
                                            </Typography>
                                            </div>
                                    </form>
                                </div>
                            </Grid>
                        </Grid>
                    </Container>
            </Content>
        </Page>
    );
};