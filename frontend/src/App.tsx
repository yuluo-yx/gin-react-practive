import {useRoutes} from "react-router-dom"
import router from  "./router/router"
import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

function App() {

  const route = useRoutes(router);

  return (
    <React.Fragment>
    <CssBaseline />
    <Container maxWidth="lg">
      <Typography component="div" style={
          { 
            backgroundColor: '#cfe8fc', 
            height: '100vh',
            paddingTop: '4vh',
            paddingLeft: '38vh'
          }
        }>
        {/* 引用路由组件 */}
        {route}
      </Typography>
    </Container>
  </React.Fragment>

  )
}

export default App
