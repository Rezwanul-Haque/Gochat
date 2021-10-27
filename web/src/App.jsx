import React, { useState, useEffect } from "react";
import { Route, Switch, useHistory } from "react-router-dom";

import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import { LinkContainer } from "react-router-bootstrap";

import CreateRoom from "./components/CreateRoom";
import Room from "./components/Room";
import Signup from "./components/Signup";
import Login from "./components/Login";
import "./App.css";

function App() {
  const history = useHistory();
  // Track if authentication is in progress
  const [isAuthenticating, setIsAuthenticating] = useState(true);
  // Track is the user has authenticated
  const [isAuthenticated, userHasAuthenticated] = useState(false);

  // Props that'll be passed to all the routes
  const routeProps = { isAuthenticated, userHasAuthenticated };

  useEffect(() => {
    async function onLoad() {
      try {
        // Check if the user is authenticated
        // await Auth.currentSession();
        if (checkLogin) {
          setIsAuthenticating(false);
          return userHasAuthenticated(true);
        }
        userHasAuthenticated(false);
      } catch (e) {
        if (e !== "No current user") {
          alert(e);
        }
      }

      setIsAuthenticating(false);
    }

    onLoad();
  }, []);

  async function handleLogout() {
    // Log the user out
    window.localStorage.removeItem("current_user");
    window.localStorage.removeItem("channel_name");
    window.localStorage.removeItem("rtc_token");

    userHasAuthenticated(false);

    history.push("/login");
  }

  return (
    !isAuthenticating && (
      <div className="App">
        <Navbar bg="light">
          <LinkContainer to="/">
            <Navbar.Brand>Go Chat</Navbar.Brand>
          </LinkContainer>
          <Navbar.Toggle />
          <Navbar.Collapse className="justify-content-end">
            <Nav activeKey={window.location.pathname}>
              {isAuthenticated ? (
                <Nav.Link onClick={handleLogout}>Logout</Nav.Link>
              ) : (
                <>
                  <LinkContainer to="/">
                    <Nav.Link>Signup</Nav.Link>
                  </LinkContainer>
                  <LinkContainer to="/login">
                    <Nav.Link>Login</Nav.Link>
                  </LinkContainer>
                </>
              )}
            </Nav>
          </Navbar.Collapse>
        </Navbar>
        <Switch>
          <Route exact path="/">
            <Signup {...routeProps} />
          </Route>
          <Route exact path="/login">
            <Login {...routeProps} />
          </Route>
          <Route exact path="/room" component={CreateRoom}>
            {/* <CreateRoom {...routeProps} /> */}
          </Route>          
          <Route path="/room/:roomID" component={Room}>
            {/* <Room {...routeProps} /> */}
          </Route>
        </Switch>
      </div>
    )
  );
}

const checkLogin = () => {
  const currentUser = window.localStorage.getItem('current_user') ? JSON.parse(window.localStorage.getItem('current_user')) : null;
  const hasIdToken = currentUser && currentUser.idToken;
  const hasRefreshToken = currentUser && currentUser.refreshToken;
  if (hasIdToken || hasRefreshToken) return true;
  else return false;
}

export default App;
