import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import { EntPersonnel } from '../../api/models/EntPersonnel';

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
    media: {
      height: 140,
    },
  }),
);


export default function Login(props: any) {
  const classes = useStyles();
  const api = new DefaultApi();

  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const [user, setUser] = useState(String);
  const [password, setPassword] = useState(String);
  const [loading, setLoading] = useState(false);

  useEffect(() => {

    const getPersonnels = async () => {
 
      const personnels = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnels(personnels);
    };
    getPersonnels();

    const resetsetPersonnelData = async () => {
      setLoading(false);
      localStorage.setItem("Personneldata", JSON.stringify(null));
      localStorage.setItem("jobpositiondata", JSON.stringify(null));
    }
    resetsetPersonnelData();

  }, [loading]);

  const UserhandleChange = (event: any) => {
    setUser(event.target.value as string);
  };

  const PasswordhandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const LoginChecker = async () => {
    personnels.map((item: any) => {
      console.log(item.user);
      if ((item.user == user) && (item.password == password)) {
        setAlert(true);
        if (item.edges.Personnel == item.edges.Personnel) {
          localStorage.setItem("personnel", JSON.stringify(item.id));
          history.pushState("", "", "/WelcomePage");
        }
        window.location.reload(false);

      }
    })
    setStatus(true);
  };



  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับสู่ ระบบบันทึกนัดหมาย`}
      ></Header>
      <Content>
        <ContentHeader title="โปรดทำการ Login ก่อนใช้งาน">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success" onClose={() => { }}>
                  เข้าสู่ระบบสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="error" onClose={() => { setStatus(false) }}>
                    เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">  
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="email"
                  label="Email"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={user}
                  onChange={UserhandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="password"
                  label="Password"
                  variant="outlined"
                  type="password"
                  size="medium"
                  value={password}
                  onChange={PasswordhandleChange}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>
                

                
            <div className={classes.margin}>
              <Button
                style={{ width: 400, backgroundColor: "#0E9AFF" }}
                onClick={() => {
                  LoginChecker();
                }}
                variant="contained"
                color="primary"
              >
                Enter
             </Button>
            </div>

            
          </form>
        </div>
      </Content>
    </Page>
  );
}

