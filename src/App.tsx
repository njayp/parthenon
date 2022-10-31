import React from "react";
import { Page, homePage } from "./pages";
import { Route, Routes } from "react-router-dom";

function App() {
  function makeRoutes(page: Page) {
    return (
      <Route {...page.routeProps} key={page.buttonLabel}>
        {page.subPages?.map((page) => makeRoutes(page))}
      </Route>
    );
  }

  return <Routes>{makeRoutes(homePage)}</Routes>;
}

export default App;
