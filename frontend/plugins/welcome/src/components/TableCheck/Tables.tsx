import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

import moment from 'moment';
import { DefaultApi } from '../../api/apis';

import { EntPersonnel } from '../../api/models/EntPersonnel';
import { EntChecksymptom } from '../../api/models/EntChecksymptom';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [carservices, setCarservices] = useState<EntChecksymptom[]>([]);
 const [loading, setLoading] = useState(true);
 const [users, setUsers] = useState<EntPersonnel[]>([]);
 const [userid, setUser] = useState(Number);
 
 useEffect(() => {
   const getCarservices = async () => {
     const res = await api.listChecksymptom({ limit: 10, offset: 0 });
     setLoading(false);
     setCarservices(res);
   };
   getCarservices();

   const checkJobPosition = async () => {
    const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
    setLoading(false);
    if (jobdata != "เจ้าหน้าที่โอเปอร์เรเตอร์" ) {
      localStorage.setItem("userdata",JSON.stringify(null));
      localStorage.setItem("jobpositiondata",JSON.stringify(null));
      history.pushState("","","./");
           
    }
    else{
        setUser(Number(localStorage.getItem("userdata")))
    }
  }
checkJobPosition();
 }, [loading]);
 
 const deleteCarservices = async (id: number) => {
  const res = await api.deleteChecksymptom({ id: id });
  setLoading(true);
};

 return (
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
         {carservices.map((item:any) => (
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
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteCarservices(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 ลบบันทึก
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}