import React, { useState, useEffect } from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import axios from 'axios';

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary" align="center" sx={{ m: 3 }}>
      {'Copyright © '}
      <Link color="inherit" href="https://mui.com/">
        Databases Unlimited, LLC
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

async function createCluster(name, engine_version, node_count) {
  const result = await axios.post('http://localhost:1323/databases', {
    name: name,
    engine_version: engine_version,
    node_count: node_count,
  });
  console.log("new cluster:", result);
  return { name, engine_version, node_count };
}

function DatabaseTable() {
  const [rows, setRows] = useState([]);

  useEffect(() => {
    async function fetchData() {
      const result = await axios('http://localhost:1323/databases');
      setRows(result.data);
    }
    fetchData();
  }, []);

  return (
    <>
      <Grid container spacing={2}>
        <Grid item xs={9}>
        </Grid>
        <Grid item xs={3}>
          <Button onClick={() => {
            const c = createCluster('New Cluster', '15.1', 1);
            setRows([...rows, c]);
          }} variant="outlined">New Cluster</Button>
        </Grid>
      </Grid>
      <TableContainer component={Paper} sx={{ marginTop: 5 }}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Cluster</TableCell>
              <TableCell align="right">Engine Version</TableCell>
              <TableCell align="right">Node Count</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.name}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.engine_version}</TableCell>
                <TableCell align="right">{row.node_count}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}

export default function App() {
  return (
    <Container maxWidth="md">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Database Clusters
        </Typography>
        <DatabaseTable />
        <Copyright />
      </Box>
    </Container>
  );
}
