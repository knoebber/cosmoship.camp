import React from 'react';
import {
  BrowserRouter as Router,
  Link,
  Route,
  Switch,
} from 'react-router-dom';
import Signup from './Signup';
import Home from './Home';
import Login from './Login';

export default function Nav() {
  return (
    <Router>
      <nav>
        <Link to="/">Home</Link>
        <Link to="/signup">Signup</Link>
        <Link to="/login">Login</Link>
      </nav>
      <main>
        <Switch>
          <Route
            component={Signup}
            path="/signup"
          />
          <Route
            component={Login}
            path="/login"
          />
          <Route
            component={Home}
            exact
            path="/"
          />
        </Switch>
      </main>
    </Router>
  );
}
