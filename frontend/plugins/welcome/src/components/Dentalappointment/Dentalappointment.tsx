import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader, 
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Swal from 'sweetalert2';

import { EntPatient,EntDentalkind,EntPersonnel } from '../../api';

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

export default function createDentalappointment() {
  const classes = useStyles();
  const profile = { givenName: ''};
  const api = new DefaultApi();

  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [kindnames, setKindnames] = useState<EntDentalkind[]>([]);
  const [personnels, setPersonnels] = useState<EntPersonnel[]>([]);

  const [loading, setLoading] = useState(true);

  const [patientName, setPatient] = useState(Number);
  const [kindName, setKindname] = useState(Number);
  const [personnelName, setPersonnel] = useState(Number);
  const [datetime, setDateTime] = useState(String);
  const [amount, setAmount] = useState(Number);
  const [price, setPrice] = useState(Number);
  const [note, setNote] = useState(String);

  useEffect(() => {

    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(false);
      setPatients(res);
      console.log(res);
    };
    getPatients();

    const getDentalkinds = async () => {
      const res = await api.listDentalkind({ limit: 10, offset: 0 });
      setLoading(false);
      setKindnames(res);
    };
    getDentalkinds();

    const getPersonnels = async () => {
      setPersonnel(Number(localStorage.getItem("personnel")))
      const res = await api.listPersonnel({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonnels(res);
    };
    getPersonnels();
    
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
      case 'amount':
        alertMessage("warning","กรุณาจำนวนฟันที่ต้องรักษา");
        return;
      case 'price':
        alertMessage("warning","กรุณากรอกราคาค่ารักษา");
        return;
      case 'note':
        alertMessage("warning","กรุณากรอกข้อความไม่เกิน 25 ตัวอักษร ถ้าไม่มีกรอก '-' ");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  


  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/dentalappointments';
    const dentalappointment = {
      patientID: patientName,
      kindName: kindName,
      personnelID: personnelName,
      appointTime: datetime + ":00+07:00",
      amount: Number(amount),
      price: Number(price),
      note: note,
    };
    console.log(dentalappointment);

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dentalappointment),
    };

    console.log(dentalappointment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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

  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatient(event.target.value as number);
  };

  const KindnamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setKindname(event.target.value as number);
  };

  const DateTimehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDateTime(event.target.value as string);
  };

  const AmounthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAmount(event.target.value as number);
  };

  const PricehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPrice(event.target.value as number);
  };

  const NotehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNote(event.target.value as string);
  };




  return (
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'Dentalappointment '}`}
      subtitle=""
     ></Header>
      <Content>
        <ContentHeader title="บันทึกการนัดทำทันตกรรม">
        <div>
            <Link component={RouterLink} to="/SearchDentalappointment">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              ค้นหา
            </Button>
          </Link>
          </div>

          <div>
            <Link component={RouterLink} to="/WelcomePage">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              Back
            </Button>
          </Link>
          </div>
          
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div><strong>ชื่อทันตแพทย์</strong></div>
                <TextField
                                    id="personnel"
                                    type="string"
                                    size="medium"
                                    value={personnels.filter((filter:EntPersonnel) => filter.id == personnelName).map((item:EntPersonnel) => `${item.name}`)}
                                    style={{ width: 200 }}
                                />
              </FormControl>
            </div>

            
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography variant="h6" gutterBottom  align="center">
                Patient Name : 
                <Typography variant="body1" gutterBottom> 
                <Select
                  labelId="patient"
                  id="patient"
                  value={patientName}
                  onChange={PatienthandleChange}
                  style={{ width: 400 }}
                >
               {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
            </div>

            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography variant="h6" gutterBottom  align="center">
                Dental Kind Name : 
                <Typography variant="body1" gutterBottom> 
              <Select
                labelId="dentalkind"
                id="dentalkind"
                value={kindName}
                onChange={KindnamehandleChange}
                style={{ width: 400 }}
              > 
                {kindnames.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.kindname}
                      </MenuItem>
                    );
                  })}
              </Select>
                </Typography>
                </Typography>
            </FormControl>

            
            <tr>
            <td>
            <FormControl
                  fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <form className={classes.root} noValidate>
                      <TextField
                        id="datetime-local"                                                            
                        label="Date"
                        type="datetime-local"
                        //defaultValue="2017-05-24T10:30"
                        className={classes.textField}
                        onChange={DateTimehandleChange}
                        InputLabelProps={{
                          shrink: true,
                        }}
                      />
                </form>
               </FormControl>
            </td>
          </tr>

          <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>จำนวนฟันที่ต้องทำการรักษา</div>
                <TextField id="amount"  InputLabelProps={{
                  shrink: true,
                }} label="amount" variant="outlined"
                  onChange={AmounthandleChange}
                  value={amount || ''}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>กรอกค่ารักษาที่ต้องชำระ</div>
                <TextField id="price"  InputLabelProps={{
                  shrink: true,
                }} label="price" variant="outlined"
                  onChange={PricehandleChange}
                  value={price || ''}
                />
              </FormControl>
            </div>


            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div>หมายเหตุ * (ไม่เกิน 25 ตัวอักษร ถ้าไม่มีกรอก -)</div>
                <TextField id="note"  InputLabelProps={{
                  shrink: true,
                }} label="note" variant="outlined"
                  onChange={NotehandleChange}
                  value={note || ''}
                />
              </FormControl>
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
                Create
             </Button>
              </Typography>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}