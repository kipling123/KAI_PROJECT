import React, { useState } from 'react';
import LinearProgress from '@mui/material/LinearProgress';
import {
  Box, Typography, Card, CardContent, Grid, List, ListItem, ListItemText, Divider, Button, Avatar, Chip,
  Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Tabs, Tab
} from "@mui/material";
import { Build as BuildIcon, Add, Code, Engineering, Settings, BarChart, Assignment, Handyman } from "@mui/icons-material"; // Added icons
import { useTheme } from '@mui/material/styles'; // Import useTheme

const projects = [
  {
    id: 1,
    name: "Pengembangan Sistem Kontrol",
    status: "Dalam Pengerjaan",
    team: ["BS", "AW", "CD"],
    deadline: "2023-12-31",
    progress: 65
  },
  {
    id: 2,
    name: "Optimasi Produksi",
    status: "Selesai",
    team: ["DP", "ES"],
    deadline: "2023-10-15",
    progress: 100
  },
  {
    id: 3,
    name: "Desain Komponen Baru",
    status: "Perencanaan",
    team: ["BS", "ES"],
    deadline: "2024-02-28",
    progress: 15
  },
];

// Helper function for TabPanel
function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          {children}
        </Box>
      )}
    </div>
  );
}

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

export default function Rekayasa() {
  const [value, setValue] = useState(0);
  const theme = useTheme(); // Access the theme object

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ p: 3, backgroundColor: theme.palette.mode === 'dark' ? '#121212' : '#f5f5f5' }}> {/* Light grey background for the whole page */}
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <BuildIcon fontSize="large" sx={{ color: theme.palette.primary.main, mr: 2 }} />
        <Typography variant="h4" sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>
          Departemen Rekayasa
        </Typography>
      </Box>

      <Paper elevation={3} sx={{ borderRadius: 2, mb: 3, backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White background for tabs */}
        <Tabs
          value={value}
          onChange={handleChange}
          aria-label="engineering department tabs"
          centered
          indicatorColor="primary"
          textColor="inherit" // Use inherit and control color via sx
          sx={{
            '& .MuiTabs-indicator': {
              height: 4,
              borderRadius: '4px 4px 0 0',
              backgroundColor: theme.palette.primary.main, // Primary color for indicator
            },
            '& .MuiTab-root': {
              color: theme.palette.mode === 'dark' ? 'rgba(255,255,255,0.7)' : 'rgba(0,0,0,0.7)', // Unselected tab color
            },
            '& .Mui-selected': {
              color: theme.palette.mode === 'dark' ? theme.palette.primary.light : theme.palette.primary.main, // Selected tab color
            },
          }}
        >
          <Tab label="Proyek" icon={<Assignment />} iconPosition="start" {...a11yProps(0)} />
          <Tab label="Alat" icon={<Handyman />} iconPosition="start" {...a11yProps(1)} />
          <Tab label="Progres" icon={<LinearProgress sx={{ transform: 'scale(0.6)', color: 'inherit' }} />} iconPosition="start" {...a11yProps(2)} />
          <Tab label="Statistik" icon={<BarChart />} iconPosition="start" {...a11yProps(3)} />
        </Tabs>
      </Paper>

      <TabPanel value={value} index={0}>
        <Card sx={{ borderRadius: 2, boxShadow: '0 4px 20px rgba(0,0,0,0.05)', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White card background */}
          <CardContent>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
              <Typography variant="h6" sx={{ fontWeight: 'bold', color: theme.palette.text.secondary }}>
                Daftar Proyek Rekayasa
              </Typography>
              <Button
                variant="contained"
                startIcon={<Add />}
                sx={{
                  backgroundColor: theme.palette.primary.main, // Use theme primary color
                  '&:hover': { backgroundColor: theme.palette.primary.dark, transform: 'translateY(-2px)' },
                  transition: 'all 0.3s ease-in-out',
                  borderRadius: 2,
                  boxShadow: `0 4px 10px ${theme.palette.mode === 'dark' ? 'rgba(144, 202, 249, 0.3)' : 'rgba(63, 81, 181, 0.3)'}`
                }}
              >
                Tambah Proyek
              </Button>
            </Box>
            <TableContainer component={Paper} sx={{ boxShadow: 'none', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White table container background */}
              <Table aria-label="project table">
                <TableHead>
                  <TableRow sx={{ backgroundColor: theme.palette.mode === 'dark' ? '#333' : '#f0f0f0' }}> {/* Lighter grey for table head */}
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Nama Proyek</TableCell>
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Status</TableCell>
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Tim</TableCell>
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Deadline</TableCell>
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Progres</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {projects.map((project) => (
                    <TableRow
                      key={project.id}
                      sx={{
                        '&:nth-of-type(odd)': { backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#ffffff' }, // White odd rows
                        '&:hover': { backgroundColor: theme.palette.mode === 'dark' ? '#3A3A3A' : '#e0e7ed', transition: 'background-color 0.3s ease' },
                      }}
                    >
                      <TableCell sx={{ color: theme.palette.text.primary }}>{project.name}</TableCell>
                      <TableCell>
                        <Chip
                          label={project.status}
                          color={
                            project.status === "Selesai" ? "success" :
                              project.status === "Dalam Pengerjaan" ? "warning" : "info"
                          }
                          variant="filled"
                          size="small"
                        />
                      </TableCell>
                      <TableCell>
                        <Box sx={{ display: 'flex', alignItems: 'center' }}>
                          {project.team.map((member, i) => (
                            <Avatar
                              key={i}
                              sx={{
                                width: 28,
                                height: 28,
                                fontSize: 12,
                                bgcolor: theme.palette.primary.main, // Use theme primary color
                                ml: i > 0 ? -1 : 0,
                                border: '1px solid ' + (theme.palette.mode === 'dark' ? '#1E1E1E' : 'white')
                              }}
                            >
                              {member}
                            </Avatar>
                          ))}
                        </Box>
                      </TableCell>
                      <TableCell sx={{ color: theme.palette.text.primary }}>{project.deadline}</TableCell>
                      <TableCell>
                        <Box sx={{ display: 'flex', alignItems: 'center' }}>
                          <LinearProgress
                            variant="determinate"
                            value={project.progress}
                            sx={{
                              width: 100,
                              height: 8,
                              borderRadius: 4,
                              mr: 1,
                              backgroundColor: theme.palette.mode === 'dark' ? '#444' : '#e0e0e0',
                              '& .MuiLinearProgress-bar': {
                                borderRadius: 4,
                                backgroundColor: theme.palette.primary.main, // Use theme primary color
                              }
                            }}
                          />
                          <Typography variant="body2" color="text.secondary">{project.progress}%</Typography>
                        </Box>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </TabPanel>

      <TabPanel value={value} index={1}>
        <Card sx={{ borderRadius: 2, boxShadow: '0 4px 20px rgba(0,0,0,0.05)', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White card background */}
          <CardContent>
            <Typography variant="h6" sx={{ fontWeight: 'bold', mb: 2, color: theme.palette.text.secondary }}>
              Daftar Alat Rekayasa
            </Typography>
            <Grid container spacing={2}>
              <Grid item xs={12} sm={6} md={4}>
                <Card sx={{ p: 2, display: 'flex', alignItems: 'center', transition: 'transform 0.2s', '&:hover': { transform: 'translateY(-5px)', boxShadow: '0 8px 25px rgba(0,0,0,0.1)' }, backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#ffffff' }}> {/* White card background */}
                  <Avatar sx={{ bgcolor: theme.palette.primary.main, mr: 2, width: 50, height: 50 }}>
                    <Code sx={{ fontSize: 30 }} />
                  </Avatar>
                  <Box>
                    <Typography variant="subtitle1" sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Software CAD</Typography>
                    <Typography variant="body2" color="text.secondary">Versi 2023.2</Typography>
                  </Box>
                </Card>
              </Grid>
              <Grid item xs={12} sm={6} md={4}>
                <Card sx={{ p: 2, display: 'flex', alignItems: 'center', transition: 'transform 0.2s', '&:hover': { transform: 'translateY(-5px)', boxShadow: '0 8px 25px rgba(0,0,0,0.1)' }, backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#ffffff' }}> {/* White card background */}
                  <Avatar sx={{ bgcolor: theme.palette.primary.main, mr: 2, width: 50, height: 50 }}>
                    <Engineering sx={{ fontSize: 30 }} />
                  </Avatar>
                  <Box>
                    <Typography variant="subtitle1" sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Simulator Produksi</Typography>
                    <Typography variant="body2" color="text.secondary">Versi 2.1.5</Typography>
                  </Box>
                </Card>
              </Grid>
              <Grid item xs={12} sm={6} md={4}>
                <Card sx={{ p: 2, display: 'flex', alignItems: 'center', transition: 'transform 0.2s', '&:hover': { transform: 'translateY(-5px)', boxShadow: '0 8px 25px rgba(0,0,0,0.1)' }, backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#ffffff' }}> {/* White card background */}
                  <Avatar sx={{ bgcolor: theme.palette.primary.main, mr: 2, width: 50, height: 50 }}>
                    <Settings sx={{ fontSize: 30 }} />
                  </Avatar>
                  <Box>
                    <Typography variant="subtitle1" sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Analisis Data</Typography>
                    <Typography variant="body2" color="text.secondary">Versi 1.0.3</Typography>
                  </Box>
                </Card>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      </TabPanel>

      <TabPanel value={value} index={2}>
        <Card sx={{ borderRadius: 2, boxShadow: '0 4px 20px rgba(0,0,0,0.05)', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White card background */}
          <CardContent>
            <Typography variant="h6" sx={{ fontWeight: 'bold', mb: 2, color: theme.palette.text.secondary }}>
              Progres Proyek
            </Typography>
            <TableContainer component={Paper} sx={{ borderRadius: 2, boxShadow: 'none', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White table container background */}
              <Table size="medium">
                <TableHead>
                  <TableRow sx={{ backgroundColor: theme.palette.mode === 'dark' ? '#333' : '#f0f0f0' }}> {/* Lighter grey for table head */}
                    <TableCell sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Nama Proyek</TableCell>
                    <TableCell align="right" sx={{ fontWeight: 'bold', color: theme.palette.text.primary }}>Progres</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {projects.map((project) => (
                    <TableRow
                      key={project.id}
                      sx={{
                        '&:nth-of-type(odd)': { backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#ffffff' }, // White odd rows
                        '&:hover': { backgroundColor: theme.palette.mode === 'dark' ? '#3A3A3A' : '#e0e7ed', transition: 'background-color 0.3s ease' },
                      }}
                    >
                      <TableCell sx={{ color: theme.palette.text.primary }}>{project.name}</TableCell>
                      <TableCell align="right">
                        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-end' }}>
                          <Typography variant="body2" sx={{ mr: 1, minWidth: '35px', color: theme.palette.text.primary }}>{project.progress}%</Typography>
                          <LinearProgress
                            variant="determinate"
                            value={project.progress}
                            sx={{
                              width: 120,
                              height: 8,
                              borderRadius: 3,
                              backgroundColor: theme.palette.mode === 'dark' ? '#444' : '#e0e0e0',
                              '& .MuiLinearProgress-bar': {
                                borderRadius: 3,
                                backgroundColor: theme.palette.primary.main // Use theme primary color
                              }
                            }}
                          />
                        </Box>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </CardContent>
        </Card>
      </TabPanel>

      <TabPanel value={value} index={3}>
        <Card sx={{ borderRadius: 2, boxShadow: '0 4px 20px rgba(0,0,0,0.05)', backgroundColor: theme.palette.mode === 'dark' ? '#1E1E1E' : '#ffffff' }}> {/* White card background */}
          <CardContent>
            <Typography variant="h6" sx={{ fontWeight: 'bold', mb: 3, color: theme.palette.text.secondary }}>
              Statistik Rekayasa
            </Typography>
            <Grid container spacing={3}>
              <Grid item xs={12} sm={6}>
                <Card sx={{
                  p: 3,
                  textAlign: 'center',
                  backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#f5f7fa', // Very light blue tint
                  borderRadius: 2,
                  transition: 'transform 0.2s',
                  '&:hover': { transform: 'translateY(-5px)', boxShadow: `0 8px 25px ${theme.palette.mode === 'dark' ? 'rgba(144, 202, 249, 0.1)' : 'rgba(63, 81, 181, 0.1)'}` } // Lighter hover shadow
                }}>
                  <Typography variant="h3" sx={{ fontWeight: 'bold', color: theme.palette.primary.main, mb: 1 }}>
                    {projects.length}
                  </Typography>
                  <Typography variant="body1" color="text.secondary">Total Proyek</Typography>
                </Card>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Card sx={{
                  p: 3,
                  textAlign: 'center',
                  backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#f5fff5', // Very light green tint
                  borderRadius: 2,
                  transition: 'transform 0.2s',
                  '&:hover': { transform: 'translateY(-5px)', boxShadow: `0 8px 25px ${theme.palette.mode === 'dark' ? 'rgba(129, 199, 132, 0.1)' : 'rgba(76, 175, 80, 0.1)'}` } // Lighter hover shadow
                }}>
                  <Typography variant="h3" sx={{ fontWeight: 'bold', color: theme.palette.success.main, mb: 1 }}>
                    {projects.filter(p => p.status === "Selesai").length}
                  </Typography>
                  <Typography variant="body1" color="text.secondary">Proyek Selesai</Typography>
                </Card>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Card sx={{
                  p: 3,
                  textAlign: 'center',
                  backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#fffaf5', // Very light orange tint
                  borderRadius: 2,
                  transition: 'transform 0.2s',
                  '&:hover': { transform: 'translateY(-5px)', boxShadow: `0 8px 25px ${theme.palette.mode === 'dark' ? 'rgba(255, 171, 64, 0.1)' : 'rgba(255, 152, 0, 0.1)'}` } // Lighter hover shadow
                }}>
                  <Typography variant="h3" sx={{ fontWeight: 'bold', color: theme.palette.warning.main, mb: 1 }}>
                    {projects.filter(p => p.status === "Dalam Pengerjaan").length}
                  </Typography>
                  <Typography variant="body1" color="text.secondary">Proyek Dalam Pengerjaan</Typography>
                </Card>
              </Grid>
              <Grid item xs={12} sm={6}>
                <Card sx={{
                  p: 3,
                  textAlign: 'center',
                  backgroundColor: theme.palette.mode === 'dark' ? '#2A2A2A' : '#f5faff', // Very light blue tint
                  borderRadius: 2,
                  transition: 'transform 0.2s',
                  '&:hover': { transform: 'translateY(-5px)', boxShadow: `0 8px 25px ${theme.palette.mode === 'dark' ? 'rgba(100, 181, 246, 0.1)' : 'rgba(33, 150, 243, 0.1)'}` } // Lighter hover shadow
                }}>
                  <Typography variant="h3" sx={{ fontWeight: 'bold', color: theme.palette.info.main, mb: 1 }}>
                    {projects.filter(p => p.status === "Perencanaan").length}
                  </Typography>
                  <Typography variant="body1" color="text.secondary">Proyek Perencanaan</Typography>
                </Card>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      </TabPanel>
    </Box>
  );
}