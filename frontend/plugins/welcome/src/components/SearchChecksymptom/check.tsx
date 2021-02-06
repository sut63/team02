import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import { Link, Link as RouterLink } from 'react-router-dom';
//import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
/* import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select'; */
import moment from 'moment';
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
/* import { EntUser } from '../../api/models/EntUser';
import { EntDistance } from '../../api/models/EntDistance';
import { EntUrgent } from '../../api/models/EntUrgent'; */
//import { EntUser } from '../../api/models/EntUser'; 
import { EntChecksymptom } from './../../api/models/EntChecksymptom';
import { EntPersonnel } from './../../api/models/EntPersonnel';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },

    withoutLabel: {
      marginTop: theme.spacing(1),
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
      width: 100,
    },
  }),
);

/* const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'ระบบบันทึกการใช้รถพยาบาล' };
 */
const check = {
  customercheck : true
}

export default function Searchtable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [checksymptoms, setChecksymptoms] = useState<EntChecksymptom[]>([]);
  const [loading, setLoading] = useState(true);
  const profile = { givenName: 'ระบบค้นหาการนัดตรวจอาการ' };
  const [status, setStatus] = useState(false);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  
  const [personnelid, setPersonnel] = useState(Number);
  const [patientName, setPatient] = useState(String);
  const [diseaseName, setDisease] = useState(String);
  const [doctorordersheetName, setDoctorordersheet] = useState(String);
  const [personnelName, setPersonnels] = useState(String);

  
  const [carservicesearch, setCarserviceSearch] = useState(String);
  const [edgesearch, setEdgeSearch] = useState(Number);
  //const [customersearch, setCustomerSearch] = useState(String);

  useEffect(() => {
   
    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
      setLoading(false);
      if (jobdata != "เจ้าหน้าที่โอเปอร์เรเตอร์" ) {
        localStorage.setItem("userdata",JSON.stringify(null));
        localStorage.setItem("jobpositiondata",JSON.stringify(null));
        history.pushState("","","./");
           
      }
      else{
          setPersonnel(Number(localStorage.getItem("userdata")))
      }
    }
  checkJobPosition();
  }, [loading]);



  const SearchChecksymptom = async () => {
    const res = await api.listChecksymptom({ offset: 0 });
    const search = Checksymptomsearch(res);
    setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setChecksymptoms([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setChecksymptoms(search);
                }
            })
        }

        setStatus(true);
  }

  const Checksymptomsearch = (res: any) => {
    const data = res.filter((filter: EntChecksymptom) => filter?.identitycard?.includes(carservicesearch) 
    || filter?.note?.includes(carservicesearch)|| filter?.phone?.includes(carservicesearch) 
    || filter?.date?.includes(carservicesearch) || filter?.edges?.patient?.name?.includes(patientName) 
    || filter?.edges?.disease?.name?.includes(diseaseName) || filter?.edges?.doctorordersheet?.name?.includes(doctorordersheetName)
    || filter?.edges?.personnel?.name?.includes(personnelName))
    
    if (data.length != 0 && carservicesearch != "") {
        return data;
    }
    else{
      return data;
        }
    }




    

    
    

  /* const UrgentSearch = (res: any) => {
    const data = res.filter((filter: EntCarservice) => filter?.edges?.urgentid?.id == urgentid)
    console.log(data.length)
    if (data.length != 0) {
      return data;
    }
    else {
      return res;
    }
  }
  const DistanceSearch = (res: any) => {
    const data = res.filter((filter: EntCarservice) => filter?.edges?.disid?.id == disid)
    console.log(data.length)
    if (data.length != 0) {
      return data;
    }
    else {
      return res;
    }
  } */

  const handleSearchChange = (event: any) => {
    setCarserviceSearch(event.target.value as string);
    setPersonnels(event.target.value as string);
    setDisease(event.target.value as string);
    setDoctorordersheet(event.target.value as string);
    setPatient(event.target.value as string);

  };

  
  return (
    <Page theme={pageTheme.library}>
     <Header
       title={`ท่านกำลังใช้งาน ${profile.givenName || ':)'}`}
       subtitle="สวัสดีครับท่านสมาชิกชมรมคนชอบการนัดหมายตรวจอาการ"
     >
   <Link component={RouterLink} to="/WelcomePage">
     <Button variant="contained" color="secondary" >
       หน้า Welcome
     </Button>
   </Link>
   </Header>
   
     <Content>
        <ContentHeader title="ค้นหาการนัดหมายตรวจอาการ">
        <Link component={RouterLink} to="/Checksymptom">
     <Button variant="contained" color="primary" >
       เพิ่มบันทึก
     </Button>
   </Link>

   <Link component={RouterLink} to="/Checksymptommain">
     <Button variant="contained" color="primary" >
       ตาราง
     </Button>
   </Link>
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
  
        <div className={classes.paper}><strong>ค้นหาข้อมูลผู้บริการ</strong></div>
        <TextField 
          style={{ width: 400, marginLeft: 20, marginRight: -10 }}
          id="customer"
          label=""
          variant="standard"
          color="secondary"
          type="string"
          size="medium"
          value={carservicesearch}
          onChange={handleSearchChange}
        />
        <div className={classes.margin}>
          <Button
            onClick={() => {
              SearchChecksymptom();
            }}
            variant="contained"
            color="primary"
          >
            ค้นหา
          </Button>
        </div>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
              <TableCell align="center">No.</TableCell>
                <TableCell align="center">ชื่อแพทย์</TableCell>
                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                <TableCell align="center">ชื่อโรค</TableCell>
                <TableCell align="center">ชื่อแพทย์ที่จะทำการตรวจ</TableCell>
                <TableCell align="center">วันที่และเวลา</TableCell>
                <TableCell align="center">เลขบัตรประชาชน</TableCell>
                <TableCell align="center">เบอร์โทรศัพท์</TableCell>
                <TableCell align="center">หมายเหตุ</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {checksymptoms.map((item: any) => (
                <TableRow key={item.id}>
                <TableCell align="center">{item.id}</TableCell>
                     <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
                     <TableCell align="center">{item.edges?.patient?.name}</TableCell>
                     <TableCell align="center">{item.edges?.disease?.disease}</TableCell>
                     <TableCell align="center">{item.edges?.doctorordersheet?.name}</TableCell>
                     <TableCell align="center">{moment(item.date).format('DD/MM/YYYY HH:mm')}</TableCell>
                     <TableCell align="center">{item.identitycard}</TableCell>
                     <TableCell align="center">{item.phone}</TableCell>
                     <TableCell align="center">{item.note}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>
      </Content>
    </Page>
  );
};