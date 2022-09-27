import React from "react";
import { Button, Stack } from "@mui/material";
import { Link, Outlet, useLocation } from "react-router-dom";

type HeaderButton = {
  label: string;
  path: string;
};

export function Header() {
  let location = useLocation();

  function makeButton(button: HeaderButton) {
    return (
      <Button
        component={Link}
        to={button.path}
        disabled={button.path === location.pathname}
        key={button.label}
      >
        {button.label}
      </Button>
    );
  }

  function makeHeader() {
    const headerButtons: HeaderButton[] = [
      { path: "/", label: "Home" },
      { path: "/logo", label: "Logo" },
      { path: "/bf", label: "Boyfriend" },
    ];

    return (
      <Stack spacing={2} direction="row">
        {headerButtons.map((button) => makeButton(button))}
      </Stack>
    );
  }

  return (
    <Stack spacing={2} sx={{ backgroundColor: "#282c34", minHeight: "100vh" }}>
      {makeHeader()}
      <Outlet />
    </Stack>
  );
}
