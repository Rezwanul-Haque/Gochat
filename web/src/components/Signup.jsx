import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useHistory } from "react-router-dom";
import { useFormFields } from "../lib/hooksLib";
import "./Signup.css";

export default function Signup({ userHasAuthenticated }) {
  const [fields, handleFieldChange] = useFormFields({
    email: "",
    password: "",
    confirmPassword: "",
  });
  const history = useHistory();
  const [newUser, setNewUser] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  function validateForm() {
    return (
      fields.email.length > 0 &&
      fields.password.length > 0 &&
      fields.password === fields.confirmPassword
    );
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
      await fetch("http://localhost:8080/api/v1/users/signup", {
        method: "POST",
        body: JSON.stringify(user),
        headers: { "Content-type": "application/json; charset=UTF-8" },
      })
        .then((response) => {
          localStorage.setItem(response);
        })
        .catch((err) => console.log(err));

      userHasAuthenticated(true);
      // Redirect to the homepage
      history.push("/room");
    } catch (e) {
      alert(e);
      setIsLoading(false);
    }
  }

  function renderForm() {
    return (
      <Form onSubmit={handleSubmit}>
        <Form.Group controlId="email" size="lg">
          <Form.Label>Email</Form.Label>
          <Form.Control
            autoFocus
            type="email"
            value={fields.email}
            onChange={handleFieldChange}
          />
        </Form.Group>
        <Form.Group controlId="password" size="lg">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            value={fields.password}
            onChange={handleFieldChange}
          />
        </Form.Group>
        <Form.Group controlId="confirmPassword" size="lg">
          <Form.Label>Confirm Password</Form.Label>
          <Form.Control
            type="password"
            onChange={handleFieldChange}
            value={fields.confirmPassword}
          />
        </Form.Group>
        <Button
          block
          size="lg"
          type="submit"
          variant="success"
          disabled={isLoading || !validateForm()}
        >
          Signup
        </Button>
      </Form>
    );
  }

  return <div className="Signup">{renderForm()}</div>;
}
