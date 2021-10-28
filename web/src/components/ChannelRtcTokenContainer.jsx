import { Input } from "@material-ui/core";
import React from "react";

const ChannelRtcTokenContainer = ({ rtcToken = "", roomId = "" }) => {
  const styles = {
    container: {
      display: "flex",
      alignItems: "center",
    },
  };
  return (
    <div style={styles.container}>
      <Input
        style={{ marginRight: "1rem" }}
        variant="contained"
        color="primary"
        placeholder="channel name"
        value={roomId}
      />
      <Input
        variant="contained"
        color="primary"
        placeholder="rtc token"
        value={rtcToken}
      />
    </div>
  );
};

export default ChannelRtcTokenContainer;
