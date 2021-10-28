import React, { useState } from "react";
import { Button } from "@material-ui/core";
import VideoCall from "./VideoCall";
import ChannelRtcTokenContainer from "./ChannelRtcTokenContainer";

const Room = (props) => {
  const [token, setToken] = useState("");
  const [inCall, setInCall] = useState(false);
  const [rtcTokenExists, setRtcTokenExists] = useState(false);

  const joinCall = async (e) => {
    e.preventDefault();

    const exists = window.localStorage.getItem("rtc_token") ? true : false;
    setRtcTokenExists(exists);

    if (!rtcTokenExists) {
      const resp = await fetch("http://localhost:8080/api/v1/rtc/token", {
        method: "POST",
        headers: {
          "Content-type": "application/json; charset=UTF-8",
          Authorization: window.localStorage.getItem("current_user")
            ? "Bearer " +
              JSON.parse(window.localStorage.getItem("current_user")).idToken
            : "",
        },
        body: JSON.stringify({
          channel_name: window.localStorage.getItem("channel_name")
            ? window.localStorage.getItem("channel_name")
            : null,
          role: "publisher",
          token_type: "uid",
          uid: "0",
        }),
      });

      const { rtc_token } = await resp.json();
      setToken(rtc_token);
      console.log(rtc_token);
      window.localStorage.setItem("rtc_token", rtc_token);
    } else {
      const rtc_token = window.localStorage.getItem("rtc_token");
      console.log("rtc_token", rtc_token);
      setToken(rtc_token);
    }
    
    setInCall(true);
  };

  return (
    <div style={{ height: "100%", padding: "1rem"}}>
      {inCall ? (
        <VideoCall setInCall={setInCall} channelName={props.match.params.roomID} token={token} />
      ) : (
        <div style={{ display: 'flex', justifyContent: "space-between", marginBottom: "1.25rem"}}>
        <Button variant="contained" color="primary" onClick={joinCall}>
          Join Call
        </Button>
        <ChannelRtcTokenContainer roomId={props.match.params.roomID} rtcToken={token} />
        </div>
      )}
    </div>
  );
};

export default Room;
