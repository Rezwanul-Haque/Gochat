import React from "react";
import { Button } from "@material-ui/core";

const CreateRoom = (props) => {
  const create = async (e) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8080/api/v1/room", {
      method: "GET",
      headers: {
        "Content-type": "application/json; charset=UTF-8",
        Authorization: window.localStorage.getItem("current_user")
          ? "Bearer " +
            JSON.parse(window.localStorage.getItem("current_user")).idToken
          : "",
      },
    });
    const { room_id } = await resp.json();

    console.log(room_id);

    window.localStorage.setItem("channel_name", room_id);

    props.history.push(`/room/${room_id}`);
  };

  return (
    <div style={{ height: "100%", padding: "1rem" }}>
      <Button variant="contained" color="primary" onClick={create}>
        Create Channel
      </Button>
    </div>
  );
};

export default CreateRoom;
