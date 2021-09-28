import React, {
  createContext,
  useCallback,
  useRef,
  useState,
} from 'react';
import PropTypes from 'prop-types';

const notificationDurationMS = 5000;

export const NotificationContext = createContext({
  setNotification: () => null,
  notification: '',
});

export default function NotificationProvider({ children }) {
  const timeout = useRef(null);
  const [state, setState] = useState('');
  const setNotification = useCallback((n) => {
    clearTimeout(timeout.current);
    setState(n);
    timeout.current = setTimeout(() => setNotification(''), notificationDurationMS);
  }, []);

  return (
    <NotificationContext.Provider value={{
      setNotification,
      notification: state,
    }}
    >
      {children}
    </NotificationContext.Provider>
  );
}

NotificationProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
