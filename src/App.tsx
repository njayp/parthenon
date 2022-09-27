import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { BFCaller } from "./BFCaller";
import { ThemeProvider } from "@mui/material";
import { darkTheme } from "./theme";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Header } from "./Header";

function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Header />}>
            <Route path="bf" element={<BFCaller />} />
            <Route
              path="logo"
              element={<img src={logo} className="App-logo" alt="logo" />}
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
