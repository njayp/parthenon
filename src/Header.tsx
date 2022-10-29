import React from "react";
import { Button, Stack } from "@mui/material";
import { Link, Outlet, useLocation } from "react-router-dom";
import * as Pages from "./pages";

export function Header() {
  let location = useLocation();

  function makeButton(page: Pages.Page) {
    return (
      <Button
        component={Link}
        to={page.routeProps.path || "/"} // TODO replace with error page
        disabled={location.pathname.endsWith(page.routeProps.path || "")}
        key={page.buttonLabel}
      >
        {page.buttonLabel}
      </Button>
    );
  }

  return (
    <Stack spacing={2} sx={{ backgroundColor: "#282c34", minHeight: "100vh" }}>
      <Stack spacing={2} direction="row">
        {[Pages.homePage, Pages.bfPage].map((page) => makeButton(page))}
      </Stack>
      <Outlet />
    </Stack>
  );
}
