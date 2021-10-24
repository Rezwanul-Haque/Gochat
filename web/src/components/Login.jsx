import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useHistory } from "react-router-dom";

import { useFormFields } from "../lib/hooksLib";
import "./Login.css";

const Login = ({ userHasAuthenticated }) => {
  const [fields, handleFieldChange] = useFormFields({
    email: "",
    password: "",
  });

  const history = useHistory();
  const [isLoading, setIsLoading] = useState(false);

  function validateForm() {
    return fields.email.length > 0 && fields.password.length > 0;
  }

  async function handleSubmit(event) {
    event.preventDefault();

    setIsLoading(true);

    try {
      const user = {
        email: fields.email,
        password: fields.password,
      };

      // Sign the user in
      await fetch("http://localhost:8080/api/v1/login", {
        method: "POST",
        body: JSON.stringify(user),
        headers: { "Content-type": "application/json; charset=UTF-8" },
      })
        .then((resp) => {
          // resolving promise
          if (resp.status == 401) {
            throw "email/password is not correct."
          } else {
            return resp.json();
          }
        })
        .then((data) => {
          console.log(data)
          window.localStorage.setItem("current_user", JSON.stringify(data));
          userHasAuthenticated(true);
          // Redirect to the homepage
          history.push("/room");
        })
        .catch((err) => alert(err));
    } catch (e) {
      alert(e);
      setIsLoading(false);
    }
  }

  return (
    <div className="Login">
      <Form onSubmit={handleSubmit}>
        <Form.Group size="lg" controlId="email">
          <Form.Label>Email</Form.Label>
          <Form.Control
            autoFocus
            type="email"
            value={fields.email}
            onChange={handleFieldChange}
          />
        </Form.Group>
        <Form.Group size="lg" controlId="password">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            value={fields.password}
            onChange={handleFieldChange}
          />
        </Form.Group>
        <Button block size="lg" type="submit" disabled={!validateForm()}>
          Login
        </Button>
      </Form>
    </div>
  );
};

export default Login;
