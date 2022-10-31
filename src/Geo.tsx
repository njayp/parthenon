import React from "react";
import { LoadingButton } from "@mui/lab";
import { Stack, TextField } from "@mui/material";

export function Geo() {
  const [loading, setLoading] = React.useState(false);
  const [lat, setLat] = React.useState("");
  const [long, setLong] = React.useState("");
  const [err, setErr] = React.useState("");

  function getLoc() {
    setLoading(true);
    setErr("");
    navigator.geolocation.getCurrentPosition(
      (position) => {
        setLat(String(position.coords.latitude));
        setLong(String(position.coords.longitude));
        setLoading(false);
      },
      (err) => {
        setLat("");
        setLong("");
        setErr(err.message);
        setLoading(false);
      }
    );
  }

  return (
    <Stack spacing={2} sx={{ width: "80%" }}>
      <LoadingButton loading={loading} onClick={getLoc}>
        Get Location
      </LoadingButton>
      <TextField disabled label="Latitude" value={lat} />
      <TextField disabled label="Longitude" value={long} />
      <TextField disabled label="Error" value={err} />
    </Stack>
  );
}
