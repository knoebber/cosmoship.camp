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
import Users from './Users';

export default function Nav() {
  return (
    <Router>
      <nav>
        <h1>Cosmo&apos;s Camp</h1>
        <Link to="/">Home</Link>
        <Link to="/login">Login</Link>
        <Link to="/signup">Signup</Link>
        <Link to="/users">Users</Link>
      </nav>
      <main>
        <Switch>
          <Route
            component={Users}
            path="/users"
          />
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
