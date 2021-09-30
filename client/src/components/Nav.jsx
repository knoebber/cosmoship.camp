import React from 'react';
import {
  BrowserRouter as Router,
  Link,
  Route,
  Switch,
} from 'react-router-dom';
import Home from './Home';
import Login from './Login';
import Members from './Members';

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
            component={Home}
            exact
            path="/"
          />
        </Switch>
      </main>
    </Router>
  );
}
