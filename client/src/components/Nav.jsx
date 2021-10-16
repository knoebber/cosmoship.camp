import React from 'react';
import {
  BrowserRouter as Router,
  Link,
  Route,
  Switch,
} from 'react-router-dom';
import Home from './home';
import Login from './auth/Login';
import Members from './members';
import Password from './auth/Password';

export default function Nav() {
  return (
    <Router>
      <nav>
        <h1>Cosmo&apos;s Camp</h1>
        <Link to="/">Home</Link>
        <Link to="/login">Login</Link>
        <Link to="/members">Members</Link>
      </nav>
      <main>
        <Switch>
          <Route
            component={Members}
            path="/members"
          />
          <Route
            component={Login}
            path="/login"
          />
          <Route
            component={Password}
            path="/password"
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
