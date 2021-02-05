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

import { EntSurgeryappointment } from '../../api/models/EntSurgeryappointment';



import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import SearchSurgeryappointment from '.';

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
    personnelcheck: true,
    patientcheck: true,
}

export default function SurgeryappointmentSearchPage() {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบบันทึกการนัดผ่าตัด' };
    const [surgeryappointments, setSurgeryappointments] = useState<EntSurgeryappointment[]>([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [personnelsearch, setPersonnelSearch] = useState(String);
    const [patientsearch, setPatientSearch] = useState(String);


    const SearchSurgeryappointment = async () => {
        const res = await api.listSurgeryappointment({ limit: 100, offset: 0});
        console.log(res);
        const personnelsearch = PersonnelSearch(res);
        const patientsearch = PatientSearch(personnelsearch);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setSurgeryappointments([]);
        if(patientsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setSurgeryappointments(patientsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.personnelcheck = true;
        searchcheck.patientcheck = true;
    }

    const PersonnelSearch = (res: any) => {
        const data = res.filter((filter: EntSurgeryappointment) => filter.edges?.personnel?.name?.includes(personnelsearch))
        console.log(data);
        if (data.length != 0 && personnelsearch != "") {
            return data;
        }
        else {
            searchcheck.personnelcheck = false;
            if(personnelsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const PatientSearch = (res: any) => {
        const data = res.filter((filter: EntSurgeryappointment) => filter.edges?.patient?.name?.includes(patientsearch))
        console.log(data);
        if (data.length != 0  && patientsearch != "") {
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




    const PersonnelSearchhandleChange = (event: any) => {
        setPersonnelSearch(event.target.value as string);
    };

    const PatientSearchhandleChange = (event: any) => {
        setPatientSearch(event.target.value as string);
    };

    return (
        <Page theme={pageTheme.website}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            ></Header>
            <Content>
                <ContentHeader title="ค้นหาข้อมูลการนัดผ่าตัด">
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
                        label="ค้นหาแพทย์ keyword"
                        type="string"
                        size="medium"
                        value={personnelsearch}
                        onChange={PersonnelSearchhandleChange}
                        style={{ width: 200 }}
                    />
                </FormControl>

                <FormControl
                    className={classes.margin}
                    variant="outlined"
                >
                    <TextField
                        id="search"
                        label="ค้นหาคนไข้ keyword"
                        type="string"
                        size="medium"
                        value={patientsearch}
                        onChange={PatientSearchhandleChange}
                        style={{ width: 300 }}
                    />
                </FormControl>

                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            SearchSurgeryappointment();
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
                            history.pushState("", "", "./Surgeryappointment");
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
                                <TableCell align="center">แพทย์</TableCell>
                                <TableCell align="center">คนไข้</TableCell>
                                <TableCell align="center">ประเภทการผ่าตัด</TableCell>
                                <TableCell align="center">อายุผู้เข้ารับการผ่าตัด</TableCell>
                                <TableCell align="center">เวลานัดหมาย</TableCell>
                                <TableCell align="center">เบอร์โทรศัพท์</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {surgeryappointments.map((item: EntSurgeryappointment) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.surgerytype?.typename}</TableCell>
                                    <TableCell align="center">{item.age}</TableCell>
                                    <TableCell align="center">{moment(item.appointTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                    <TableCell align="center">{item.phone}</TableCell>
                                    <TableCell align="center">{item.note}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}