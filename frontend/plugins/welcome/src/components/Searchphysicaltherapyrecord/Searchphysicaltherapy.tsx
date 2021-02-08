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

import { EntPhysicaltherapyrecord } from '../../api/models/EntPhysicaltherapyrecord';



import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import SearchPhysicaltherapyrecord from '.';

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

export default function PhysicaltherapyrecordSearch () {
    const classes = useStyles();
    const api = new DefaultApi();
    const profile = { givenName: 'ยินดีต้อนรับ' };
    const [physicaltherapyrecords, setphysicaltherapyrecords] = useState<EntPhysicaltherapyrecord []>([]);
    const [loading, setLoading] = useState(true);
    const [status, setStatus] = useState(false);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [personnelsearch, setPersonnelSearch] = useState(String);
    const [patientsearch, setPatientSearch] = useState(String);


    const SearchPhysicaltherapyrecord = async () => {
        const res = await api.listPhysicaltherapyrecord({ limit: 100, offset: 0});
        console.log(res);
        const personnelsearch = PersonnelSearch(res);
        const patientsearch = PatientSearch(personnelsearch);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setphysicaltherapyrecords([]);
        if(patientsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setphysicaltherapyrecords(patientsearch);
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
        const data = res.filter((filter: EntPhysicaltherapyrecord) => filter.edges?.personnel?.name?.includes(personnelsearch))
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
        const data = res.filter((filter: EntPhysicaltherapyrecord) => filter.edges?.patient?.name?.includes(patientsearch))
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
                            SearchPhysicaltherapyrecord();
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
                            <TableCell align="center">อันดับที่</TableCell>
                <TableCell align="center">เลขเลขบัตรประชาชน</TableCell>
                <TableCell align="center">อายุ</TableCell>
                <TableCell align="center">วัน-เวลา</TableCell>
                <TableCell align="center">เบอร์โทรผู้ป่วย</TableCell>
                <TableCell align="center">ผู้ป่วยที่มารับการตรวจ</TableCell>
                <TableCell align="center">ผู้บันทึกข้อมูล</TableCell>
                <TableCell align="center">ห้องกายภาพบำบัด</TableCell>
                <TableCell align="center">อาการ</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {physicaltherapyrecords.map((item: EntPhysicaltherapyrecord) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.idnumber}</TableCell>
                  <TableCell align="center">{item.age}</TableCell>
                  <TableCell align="center">{moment(item.appointtime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                  <TableCell align="center">{item.telephone}</TableCell>  
                  <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                  <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                  <TableCell align="center">{item.edges?.physicaltherapyroom?.physicaltherapyroomname}</TableCell>
                  <TableCell align="center">{item.edges?.status?.statusname}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
}