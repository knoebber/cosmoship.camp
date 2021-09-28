import React, { useContext } from 'react';
import { NotificationContext } from '../providers/NotificationProvider';

export default function Notification() {
  const { notification } = useContext(NotificationContext);

  return notification && <div id="notification">{notification}</div>;
}
