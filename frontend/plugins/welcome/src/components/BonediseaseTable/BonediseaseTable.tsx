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
import { DefaultApi } from '../../api/apis';
import { EntBonedisease } from '../../api/models/EntBonedisease';
import moment from 'moment';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 
 const [bonedisease, setBonediseases] = useState<EntBonedisease[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
  const getBonedisease = async () => {
    const res = await api.listBonedisease({ limit: 10, offset: 0 });
    setLoading(false);
    setBonediseases(res);
    console.log(res);
  };
  getBonedisease();

}, [loading]);
 
 const deleteBonedisease = async (id: number) => {
   const res = await api.deleteBonedisease({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
            <TableCell align="center">ลำดับ</TableCell>
            <TableCell align="center">เลือกผู้ใช้งาน</TableCell>
            <TableCell align="center">เลือกผู้ป่วย</TableCell>
            <TableCell align="center">เลือกวิธีการรักษา</TableCell>
            <TableCell align="center">แนะนำ</TableCell>
            <TableCell align="center">เลือกเวลา</TableCell>
            <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {bonedisease === undefined
            ? null
            : bonedisease.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.personnel?.name}</TableCell>
             <TableCell align="center">{item.edges?.patient?.name}</TableCell>
             <TableCell align="center">{item.edges?.remedy?.remedy}</TableCell>
             <TableCell align="center">{item.edges?.advice?.advice}</TableCell>
             <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   if (item.id === undefined){
                     return;
                   }
                   deleteBonedisease(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
