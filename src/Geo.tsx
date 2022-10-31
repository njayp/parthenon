import React from "react";
import { Button, Stack, TextField } from "@mui/material";

export function Geo() {
  const [lat, setLat] = React.useState("");
  const [long, setLong] = React.useState("");
  const [err, setErr] = React.useState("");

  function getLoc() {
    setLat("Loading...");
    setLong("Loading...");
    setErr("");
    navigator.geolocation.getCurrentPosition(
      (position) => {
        setLat(String(position.coords.latitude));
        setLong(String(position.coords.longitude));
      },
      (err) => {
        setLat("");
        setLong("");
        setErr(err.message);
      }
    );
  }

  return (
    <Stack spacing={2} sx={{ width: "80%" }}>
      <Button onClick={getLoc}>Get Location</Button>
      <TextField disabled label="Latitude" value={lat} />
      <TextField disabled label="Longitude" value={long} />
      <TextField disabled label="Error" value={err} />
    </Stack>
  );
}
