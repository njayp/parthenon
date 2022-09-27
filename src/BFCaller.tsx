import React from "react";
import { BFFClient } from "./api/BffServiceClientPb";
import { BoyfriendRequest } from "./api/bff_pb";
import { Button, TextField, Stack } from "@mui/material";

const address = "http://localhost:8080";
const client = new BFFClient(address, null, null);

export function BFCaller() {
  const [text, setText] = React.useState("");
  const [err, setErr] = React.useState("");

  function Meow() {
    client.boyfriendBot(new BoyfriendRequest(), null, (err, response) => {
      setText(response?.getEmoji());
      setErr(err?.message);
    });
  }

  return (
    <Stack spacing={2} sx={{ width: "80%" }}>
      <Button onClick={Meow}>Meow at Boyfriend</Button>
      <TextField disabled label="Response" value={text} />
      <TextField disabled label="Error" value={err} />
    </Stack>
  );
}
