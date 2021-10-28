import { AgoraVideoPlayer } from "agora-rtc-react";
import { Grid } from "@material-ui/core";
import React, { useState, useEffect } from "react";

export default function Video(props) {
  const { users, tracks } = props;
  const [gridSpacing, setGridSpacing] = useState(12);
  console.log(users);
  useEffect(() => {
    setGridSpacing(Math.max(Math.floor(12 / (users.length + 1)), 6));
  }, [users, tracks]);

  return (
    <Grid container style={{ height: "100%" }}>
      <Grid item xs={gridSpacing}>
        <AgoraVideoPlayer
          videoTrack={tracks[1]}
          style={{
            height: "100%",
            width: "100%",
            border: "1px solid",
            position: "relative",
          }}
        />
      </Grid>
      {users.length > 0 &&
        users.map((user) => {
          if (user.videoTrack) {
            return (
              <Grid item xs={gridSpacing}>
                <AgoraVideoPlayer
                  videoTrack={user.videoTrack}
                  key={user.uid}
                  style={{
                    height: "100%",
                    width: "100%",
                    border: "1px solid",
                    position: "relative",
                  }}
                />
              </Grid>
            );
          } else return null;
        })}
    </Grid>
  );
}
