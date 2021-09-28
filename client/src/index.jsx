import React from 'react';
import ReactDOM from 'react-dom';
import './style.css';
import Nav from './components/Nav';
import Notification from './components/Notification';
import NotificationProvider from './providers/NotificationProvider';

ReactDOM.render(
  <React.StrictMode>
    <NotificationProvider>
      <Notification />
      <Nav />
    </NotificationProvider>
  </React.StrictMode>,
  document.getElementById('root'),
);
