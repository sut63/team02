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

import { EntUser } from '../../api/models/EntUser';
import { EntSemester } from '../../api/models/EntSemester';
import { EntApprovedresult } from '../../api/models/EntApprovedresult';
import { EntScholarshippaidrecord } from '../../api/models/EntScholarshippaidrecord';
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
  const profile = { givenName: 'บันทึกการจ่ายทุนการศึกษา' };
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  const [users, setUser] = useState<EntUser[]>([]);
  const [semesters, setSemester] = useState<EntSemester[]>([]);
  const [approvedresults, setApprovedresult] = useState<EntApprovedresult[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [userid, setUserID] = useState(Number);
  const [semesterid, setSemesterID] = useState(Number);
  const [approvedresultid, setApprovedresultID] = useState(Number);
  const [addtime, setAddtime] = useState(String);

  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUser(res);
    };
    getUser();

    const getSemester = async () => {
      const res = await api.listSemester({ limit: 10, offset: 0 });
      setLoading(false);
      setSemester(res);
    };
    getSemester();

    const getApprovedresult = async () => {
      const res = await api.listApprovedresult({ limit: 10, offset: 0 });
      setLoading(false);
      setApprovedresult(res);
    };
    getApprovedresult();

  }, [loading]);

  
   const handletimeChange = (event: any) => {
    setAddtime(event.target.value as string);
   };

  const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

  const SemesterhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSemesterID(event.target.value as number);
  };

  const ApprovedresulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setApprovedresultID(event.target.value as number);
  };


  const CreateScholarshippaidrecord = async () => {
    const schordshippaidrecord = {
      userId: userid,
      semesterId: semesterid,
      approvedresultId: approvedresultid,
      datetime: addtime + ":00+07:00" 
    };

    console.log(schordshippaidrecord);
    const res: any = await api.createSchordshippaidrecord({ schordshippaidrecord: schordshippaidrecord });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    };
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  
  };
  

  return (
    <Page theme={pageTheme.website}>
      <Header
        title={` ${profile.givenName || 'to Backstage'}`}
        subtitle=""
      ></Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการจ่ายทุนการศึกษา">
        <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
             </Button>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning">
                    ข้อมูลไม่ถูก
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div>
            <p><font size='2'><b>ป้อนชื่อผู้จ่ายทุน</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="user-label"></InputLabel>
              <Select
                labelId="user-label"
                id="user"
                value={userid}
                onChange={UserhandleChange}
                style={{ width: 250 }}
              >
                {users.map((item: EntUser) => (
                  <MenuItem value={item.id}>{item.firstname}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>

          <div>
            <p><font size='2'><b>ป้อนเลขที่การอนุมัติทุน</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="approvedresult-label"></InputLabel>
              <Select
                labelId="approvedresult-label"
                id="approvedresult"
                value={approvedresultid}
                onChange={ApprovedresulthandleChange}
                style={{ width: 250 }}
              >
                {approvedresults.map((item: EntApprovedresult) => (
                  <MenuItem value={item.id}>{item.id}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>
          <div>
            <p><font size='2'><b>ป้อนภาคการศึกษา</b></font></p>  
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <InputLabel id="semester-label"></InputLabel>
              <Select
                labelId="semester-label"
                id="semester"
                value={semesterid}
                onChange={SemesterhandleChange}
                style={{ width: 250 }}
              >
                {semesters.map((item: EntSemester) => (
                  <MenuItem value={item.id}>{item.term}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </div>
          <div>
          <p><font size='2'><b>เวลา</b></font></p> 
          <TextField
        id="datetime-local"
        label="Next appointment"
        type="datetime-local"
        //defaultValue="2020-10-24T11:30"
        value={addtime}
        className={classes.textField}
        onChange={handletimeChange}
        InputLabelProps={{
          shrink: true,
        }}
      />
          </div>
          
          <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateScholarshippaidrecord();
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