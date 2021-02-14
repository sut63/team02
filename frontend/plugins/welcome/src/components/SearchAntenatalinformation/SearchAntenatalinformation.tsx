import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';

import moment from 'moment';

import { EntAntenatalinformation } from '../../api/models/EntAntenatalinformation';



import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import SearchAntenatalinformation from '.';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        table: {
            minWidth: 650,
        },
        margin: {
            margin: theme.spacing(3),
        },
    }),
);

const searchcheck = {
    patientcheck: true,
    pregnancystatuscheck: true,
}

export default function AntenatalinformationSearchPage() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบบันทึกการนัดฝากครรภ์' };
    const [Antenatalinformations, setAntenatalinformations] = useState<EntAntenatalinformation[]>([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [patientsearch, setPatientSearch] = useState(String);
    const [pregnancystatussearch, setPregnancystatusSearch] = useState(String);


    const SearchAntenatalinformation = async () => {
        const res = await api.listAntenatalinformation({ limit: 100, offset: 0});
        console.log(res);
        const patientsearch = PatientSearch(res);
        const pregnancystatussearch = PregnancystatusSearch(patientsearch);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setAntenatalinformations([]);
        if(pregnancystatussearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setAntenatalinformations(pregnancystatussearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.patientcheck = true;
        searchcheck.pregnancystatuscheck = true;
    }

    const PatientSearch = (res: any) => {
        const data = res.filter((filter: EntAntenatalinformation) => filter.edges?.patient?.name?.includes(patientsearch))
        console.log(data);
        if (data.length != 0 && patientsearch != "") {
            return data;
        }
        else {
            searchcheck.patientcheck = false;
            if(patientsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const PregnancystatusSearch = (res: any) => {
        const data = res.filter((filter: EntAntenatalinformation) => filter.edges?.pregnancystatus?.pregnancystatus?.includes(pregnancystatussearch))
        console.log(data);
        if (data.length != 0  && pregnancystatussearch != "") {
            return data;
        }
        else {
            searchcheck.pregnancystatuscheck = false;
            if(pregnancystatussearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }




    const PatientSearchhandleChange = (event: any) => {
        setPatientSearch(event.target.value as string);
    };

    const PregnancystatusSearchhandleChange = (event: any) => {
        setPregnancystatusSearch(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.website}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="ค้นหาข้อมูลการนัดฝากครรภ์">
                {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>
                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาผู้รับการตรวจ"
                        type="string"
                        size="medium"
                        value={patientsearch}
                        onChange={PatientSearchhandleChange}
                        style={{ width: 200 }}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาสถานะครรภ์"
                        type="string"
                        size="medium"
                        value={pregnancystatussearch}
                        onChange={PregnancystatusSearchhandleChange}
                        style={{ width: 300 }}
                    />
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                          SearchAntenatalinformation();
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, backgroundColor: "#365391" }}
                        startIcon={<SearchIcon />}
                    >
                        ค้นหา
                    </Button>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "./Antenatal");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 180, marginLeft: 40, backgroundColor: "#a31f1f" }}
                        startIcon={<ArrowBackIcon />}
                    >
                        กลับ
                    </Button>
                </div>

                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">แพทย์ผู้ตรวจ</TableCell>
                                <TableCell align="center">ผู้รับการตรวจ</TableCell>
                                <TableCell align="center">สถานะครรภ์</TableCell>
                                <TableCell align="center">อายุครรภ์</TableCell>
                                <TableCell align="center">ผลการตรวจ</TableCell>
                                <TableCell align="center">ความเสี่ยงที่พบ</TableCell>
                                <TableCell align="center">เวลา</TableCell>
                                <TableCell align="center">คำแนะนำ</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {Antenatalinformations.map((item: EntAntenatalinformation) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.pregnancystatus?.pregnancystatus}</TableCell>
                                    <TableCell align="center">{item.gestationalage}</TableCell>
                                    <TableCell align="center">{item.examinationresult}</TableCell>
                                    <TableCell align="center">{item.edges?.risks?.risks}</TableCell>
                                    <TableCell align="center">{moment(item.time).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                    <TableCell align="center">{item.advice}</TableCell>
                              
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}