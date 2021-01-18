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
import { EntPregnancystatus } from '../../api/models/EntPregnancystatus';
import { EntRisks } from '../../api/models/EntRisks';
import { Alert } from '@material-ui/lab';

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

export default function Antenatal() {
    const classes = useStyles();
    const profile = { givenName: 'ระบบบันทึกนัดฝากครรภ์' };
    const api = new DefaultApi();

    const [personnels, setPersonnel] = useState<EntPersonnel[]>([]);
    const [patients, setPatient] = useState<EntPatient[]>([]);
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
    const [time, setTime] = useState(String);

    useEffect(() => {
        const getPersonnel = async () => {
            const per = await api.listPersonnel({ limit: 10, offset: 0 });
            setLoading(false);
            setPersonnel(per);
          };
        getPersonnel();
    
        const getPatient = async () => {
            const pat = await api.listPatient({ limit: 10, offset: 0 });
            setLoading(false);
            setPatient(pat);
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

    const createAntenatalinformation = async () => {
        if( ( personnelid != 0 ) && ( patientid != 0) && ( pregnancystatusid != 0) && ( risksid != 0) && ( gestationalage != 0) && ( time != "" )){
            const antenatalinformation = {
                    personnel: personnelid,
                    patient: patientid,
                    pregnancystatus: pregnancystatusid,
                    risks: risksid,
                    gestationalage: +gestationalage,
                    time: time + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
            };
            
            console.log(antenatalinformation);
            const res: any = await api.createAntenatalinformation({ antenatalinformation: antenatalinformation });
            setStatus(true);
            if (res.id != '') {
                setAlert(true);
            } 
        }
        else {
            setAlert(false);setStatus(true);
        };
        const timer = setTimeout(() => {
            setStatus(false);
        }, 1000);
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

    const TimehandleDateChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setTime(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.service}>
            <Header title={`${profile.givenName || 'ระบบันทึกนัดฝากครรภ์'}`}></Header>
            <Content>
                <ContentHeader title="กรุณากรอกข้อมูล"></ContentHeader>
                    <Container maxWidth="sm">
                        <Grid container spacing={3}>
                            <Grid item xs={4}>
                                <div className={classes.root}>
                                    <form noValidate autoComplete="off">
                            
                                        <div>
                                            <Grid item xs={8}>
                                                <FormControl variant="filled" className={classes.formControl}>
                                                    <InputLabel>แพทย์ผู้ตรวจ</InputLabel>
                                                    <Select
                                                        labelId="แพทย์ผู้ตรวจ"
                                                        id="personnel"
                                                        value={personnelid}
                                                        onChange={PersonnelhandleChange}
                                                    >
                                                        {personnels.map((item: EntPersonnel) =>
                                                        <MenuItem value={item.id}>{item.name}</MenuItem>)
                                                        }
                                                    </Select>
                                                </FormControl>
                                            </Grid>
                                        </div>

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
                                                        <MenuItem value={item.id}>{item.name}</MenuItem>)
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

                                        <div className={classes.margin}>
                                            <Button
                                                onClick={() => {createAntenatalinformation();}}
                                                variant="contained"
                                                color="primary"
                                            >
                                            ยืนยัน
                                            </Button>
                                            {status ? (
                                                <div>
                                                    {alert ? (<Alert severity="success">บันทึกสำเร็จ</Alert>) 
                                                    : (<Alert severity="warning">ข้อมูลไม่ถูกต้อง</Alert>)}
                                                </div>
                                            ) : null}
                                            <Button style={{marginLeft: 20,backgroundColor: 'red'}}
                                                component={RouterLink}
                                                to="/WelcomePage"
                                                variant="contained"
                                            >
                                            ยกเลิก
                                            </Button>
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