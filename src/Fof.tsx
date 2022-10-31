import React from "react";
import { Button } from "@mui/material";
import { Link } from "react-router-dom";

export function Fof() {
  return (
    <Button component={Link} to="/">
      <img src={"/fof.png"} alt="404 dragon" />
    </Button>
  );
}
