import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { BFCaller } from "./BFCaller";
import { Button, Stack } from "@mui/material";
import { Route, Routes, Link, Outlet, useLocation } from "react-router-dom";

type Page = {
  buttonLabel: string;
  path: string;
  element: JSX.Element;
  subPages?: Page[];
};

function App() {
  let location = useLocation();

  const rootPage: Page = {
    path: "/",
    buttonLabel: "Home",
    element: <Header />,
    subPages: [
      {
        path: "/logo",
        buttonLabel: "Logo",
        element: <img src={logo} className="App-logo" alt="logo" />,
      },
      { path: "/bf", buttonLabel: "Boyfriend", element: <BFCaller /> },
    ],
  };

  function Header() {
    function makeButtons(page: Page) {
      let buttons = [
        <Button
          component={Link}
          to={page.path}
          disabled={page.path === location.pathname}
          key={page.buttonLabel}
        >
          {page.buttonLabel}
        </Button>,
      ];

      page.subPages?.forEach((page) => {
        buttons = buttons.concat(makeButtons(page));
      });

      return buttons;
    }

    return (
      <Stack
        spacing={2}
        sx={{ backgroundColor: "#282c34", minHeight: "100vh" }}
      >
        <Stack spacing={2} direction="row">
          {makeButtons(rootPage)}
        </Stack>
        <Outlet />
      </Stack>
    );
  }

  function makeRoutes(page: Page) {
    return (
      <Route path={page.path} element={page.element}>
        {page.subPages?.map((page) => makeRoutes(page))}
      </Route>
    );
  }

  return <Routes>{makeRoutes(rootPage)}</Routes>;
}

export default App;
